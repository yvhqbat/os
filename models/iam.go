package models

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	etcd "go.etcd.io/etcd/clientv3"
)

const(
	iamUsersPrefix = "iam/users/"
	iamGroupPrefix = "iam/groups/"

	defaultTimeOut = 30*time.Second
)

//type UserInfo struct {
//	ID       string `json:"id"`
//	Name     string `json:"name"`
//	Password string `json:"password"`
//	AK       string `json:"ak"`
//	SK       string `json:"sk"`
//}

type GroupInfo struct{

}

type IAMSys struct {
	sync.RWMutex
	iamUsersMap map[string]UserInfo
	iamGroupsMap map[string]GroupInfo

	etcdClient *etcd.Client
}

func NewIAMSys(c *etcd.Client)*IAMSys{
	return &IAMSys{
		iamUsersMap: make(map[string]UserInfo),
		iamGroupsMap: make(map[string]GroupInfo),
		etcdClient:c,
	}
}

func (sys *IAMSys)Init()error{
	if sys.etcdClient == nil{
		return errors.New("etcd client is nil")
	}

	defer sys.watchIAMEtcd()

	return sys.refreshEtcd()
}

// etcd watch 
func (sys *IAMSys)watchIAMEtcd(){
	watchEtcd := func() {
		for  {
			// watchChan := etcd.wa``
			time.Sleep(1*time.Second)
		}
	}

	go watchEtcd()
}

func (sys *IAMSys)LoadEtcdUser(userID string)(u UserInfo, err error){
	log.Debugf("load user, id is %s", userID)

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	resp, err := sys.etcdClient.Get(ctx, userID)
	if err!=nil{
		log.Errorf("get user info failed, user id is %s", userID)
		return
	}
	if resp.Count == 0{
		err = errors.New("not found error")
		return
	}

	for _, ev := range resp.Kvs{
		if string(ev.Key) == userID{
			err = json.Unmarshal(ev.Value, u)
			if err!=nil{
				log.Errorf("unmashal failed.")
				return
			}
		}
	}

	err = errors.New("not found error")

	return
}

func (sys *IAMSys)LoadEtcdUsers(m map[string]UserInfo)error{
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeOut)
	defer cancel()

	r, err := sys.etcdClient.Get(ctx, iamUsersPrefix, etcd.WithPrefix(), etcd.WithKeysOnly())
	if err!=nil{
		return err
	}
	for _, u := range r.Kvs{
		userID := u.Key
		user, err :=  sys.LoadEtcdUser(string(userID))
		if err!=nil{
			return err
		}
		m[string(userID)] = user
	}

	return nil
}

func (sys *IAMSys)LoadEtcdGroups(m map[string]GroupInfo)error{

	return nil
}

func (sys *IAMSys)refreshEtcd()error{
	iamUsersMap := make(map[string]UserInfo)
	iamGroupsMap := make(map[string]GroupInfo)

	err := sys.LoadEtcdUsers(iamUsersMap)
	if err!=nil{
		return err
	}

	err = sys.LoadEtcdGroups(iamGroupsMap)
	if err!=nil{
		return err
	}

	sys.Lock()
	defer sys.Unlock()
	sys.iamUsersMap = iamUsersMap
	sys.iamGroupsMap = iamGroupsMap

	return nil
}

