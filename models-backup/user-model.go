package models

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// go-sql-driver 参考 https://studygolang.com/articles/13655

type UserInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	AK       string `json:"ak"`
	SK       string `json:"sk"`
}

type userManager struct {
	redisEndpoint string
	redisPool     *redis.Pool

	mysqlDsn string
	mysqlDb  *sql.DB

	initFlag bool
}

/*
@param  redisEndpoint  "127.0.0.1:6379"
@param  mysqlDsn       "root:123456@tcp(127.0.0.1:3306)/<db>"
*/
func NewUserManager(redisEndpoint string, mysqlDsn string) *userManager {
	return &userManager{
		redisEndpoint,
		nil,
		mysqlDsn,
		nil,
		false,
	}
}

func (m *userManager) Initialize() error {
	if m.initFlag {
		return nil
	}

	pool := &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", m.redisEndpoint)
		},
	}
	m.redisPool = pool

	db, err := sql.Open("mysql", m.mysqlDsn)
	if err != nil {
		log.Errorln(err)
		return err
	}

	db.SetMaxOpenConns(16)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(60 * time.Second)

	SQL := `CREATE TABLE IF NOT EXISTS 
			user(id varchar(64) PRIMARY KEY NOT NULL, 
				name varchar(64), 
				password varchar(64), 
				ak varchar(64), 
				sk varchar(64)
			  )`

	_, err = db.Exec(SQL)
	if err != nil {
		log.Errorln(err)
		db.Close()
		return err
	}
	m.mysqlDb = db

	m.initFlag = true

	return nil
}

func (m *userManager) Close() error {
	if m.initFlag == true {
		m.redisPool.Close()
		m.mysqlDb.Close()
		m.initFlag = false
	}
	return nil
}

// get
func (m *userManager) Get(id string) (*UserInfo, error) {
	userInfo := new(UserInfo)
	// 1. 从缓存获取
	conn := m.redisPool.Get()
	defer conn.Close()
	data, err := redis.String(conn.Do("HGET", "user:", id))
	if err == nil {
		// 1.1 命中， 返回
		err = json.Unmarshal([]byte(data), userInfo)
		if err != nil {
			log.Errorln(err)
			return nil, err
		}

		return userInfo, nil
	}

	// 1.2 未命中，从数据库获取
	SQL := `SELECT * FROM user WHERE id=?`
	row := m.mysqlDb.QueryRow(SQL, id)
	err = row.Scan(&userInfo.ID, &userInfo.Name, &userInfo.Password, &userInfo.AK, &userInfo.SK)
	if err != nil {
		// 1.2.1 数据库结果为空， 返回
		log.Errorln(err)
		return nil, err
	}

	// 1.2.2 数据库结果不为空，插入缓存后返回
	_, err = conn.Do("HSET", "users:", userInfo.AK, userInfo.ID)
	if err != nil {
		log.Warnln(err)
		// 插入缓存失败，不返回错误
	} else {
		data, err := json.Marshal(userInfo)
		if err != nil {
			log.Warnln(err)
		} else {
			_, err = conn.Do("HSET", "user:", userInfo.ID, data)
			if err != nil {
				log.Warnln(err)
			}
		}
	}

	return userInfo, nil
}

// update // 2. 延时双删
func (m *userManager) Put(userinfo *UserInfo) error {

	return nil
}

// insert CreateUser
func (m *userManager) CreateUser(userInfo *UserInfo) error {
	var err error

	// 1. 分布式锁
	lock_id := strings.Join([]string{"user", userInfo.ID}, ":")
	log.Printf("lock %s\n", lock_id)

	// 2. 插入数据库
	SQL := `INSERT INTO user(id, name, password, ak , sk) VALUES (?,?,?,?,?)`
	_, err = m.mysqlDb.Exec(SQL, userInfo.ID, userInfo.Name, userInfo.Password, userInfo.AK, userInfo.SK)
	if err != nil {
		log.Errorln(err)
		return err
	}

	// 3. 插入缓存
	// users:
	conn := m.redisPool.Get()
	defer conn.Close()

	_, err = conn.Do("HSET", "users:", userInfo.AK, userInfo.ID)
	if err != nil {
		log.Errorln(err)
		return err
	}

	// user:id
	//conn.Do("hset", "user:"+userinfo.ID, "id", userinfo.ID)
	//conn.Do("hset", "user:"+userinfo.ID, "name", userinfo.Name)
	//conn.Do("hset", "user:"+userinfo.ID, "password", userinfo.Password)
	//conn.Do("hset", "user:"+userinfo.ID, "ak", userinfo.AK)
	//conn.Do("hset", "user:"+userinfo.ID, "sk", userinfo.SK)

	data, err := json.Marshal(userInfo)
	if err != nil {
		log.Errorln(err)
		return err
	}
	_, err = conn.Do("HSET", "user:", userInfo.ID, data)
	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

// delete
func (m *userManager) Delete(id string) error {
	// 1. 从缓存删除
	conn := m.redisPool.Get()
	defer conn.Close()

	_, err := conn.Do("HDEL", "user:", id)
	if err != nil {
		log.Errorln(err)
		return err
	}

	// 2. 从数据库删除
	SQL := `DELETE FROM user WHERE id=?`
	_, err = m.mysqlDb.Exec(SQL, id)
	if err != nil {
		log.Errorln(err)
		return err
	}

	// 3. 延时
	time.Sleep(1 * time.Second)

	// 4. 再次从缓存删除
	_, err = conn.Do("HDEL", "user:", id)
	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}
