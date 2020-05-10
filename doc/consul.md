
## 1. consul

参考:
- [https://github.com/hashicorp/consul](https://github.com/hashicorp/consul)
- [使用consul实现分布式服务注册和发现](https://studygolang.com/articles/4476)
- [golang使用服务发现系统consul](https://studygolang.com/articles/9980)

## 2. 服务注册与发现
采用静态配置的方式进行注册, interface server 实现服务发现;

## 3. 统一配置中心


- [https://github.com/hashicorp/consul](https://github.com/hashicorp/consul)
- [https://github.com/hashicorp/consul/tree/master/api](https://github.com/hashicorp/consul/tree/master/api)
- [基于consul构建golang系统分布式服务发现机制](https://www.cnblogs.com/williamjie/p/9369774.html)
- [consul总结](https://www.cnblogs.com/duanxz/tag/consul/)
- [golang使用服务发现系统consul](https://studygolang.com/articles/9980)

## 1. start
```
./consul agent -dev -config-dir=./consul.d -advertise=10.199.105.9 -bind=0.0.0.0 -client=0.0.0.0
```

## 2. discovery
```
[root@node19062652 ~]# curl http://10.199.105.9:8500/v1/agent/services
{
    "hgw": {
        "ID": "hgw",
        "Service": "hgw",
        "Tags": [
            "gateway"
        ],
        "Meta": {},
        "Port": 8000,
        "Address": "",
        "Weights": {
            "Passing": 1,
            "Warning": 1
        },
        "EnableTagOverride": false
    }
}
[root@node19062652 ~]#
[root@node19062652 ~]#
[root@node19062652 ~]# curl http://10.199.105.9:8500/v1/agent/service/hgw
{
    "ID": "hgw",
    "Service": "hgw",
    "Tags": [
        "gateway"
    ],
    "Meta": null,
    "Port": 8000,
    "Address": "",
    "Weights": {
        "Passing": 1,
        "Warning": 1
    },
    "EnableTagOverride": false,
    "ContentHash": "6b6a4e55ea1fcc1e"
}
[root@node19062652 ~]#

```


## health
```
[root@node19062652 ~]# curl http://10.199.105.9:8500/v1/agent/health/service/name/hgw
[
    {
        "AggregatedStatus": "passing",
        "Service": {
            "ID": "hgw",
            "Service": "hgw",
            "Tags": [
                "gateway"
            ],
            "Meta": {},
            "Port": 8000,
            "Address": "",
            "Weights": {
                "Passing": 1,
                "Warning": 1
            },
            "EnableTagOverride": false
        },
        "Checks": []
    }
]
[root@node19062652 ~]# curl http://10.199.105.9:8500/v1/agent/health/service/name/hgw?format=text
passing
[root@node19062652 ~]#
```

## register
```
[root@node19062652 yvh]# cat payload.json
{
  "ID": "hgw",
  "Name": "hgw",
  "Tags": [
    "primary",
    "v1"
  ],
  "Address": "10.192.70.52:80",
  "Port": 8000,
  "Meta": {
    "version": "1.2.2"
  },
  "EnableTagOverride": false,
  "Check": {
    "DeregisterCriticalServiceAfter": "90m",
    "HTTP": "http://10.192.70.52:80/",
    "Interval": "10s"
  },
  "Weights": {
    "Passing": 10,
    "Warning": 1
  }
}
```

```
curl     --request PUT     --data @payload.json     http://10.199.105.9:8500/v1/agent/service/register
```

## deregister
```
[root@node19062652 yvh]# curl     --request PUT     http://10.199.105.9:8500/v1/agent/service/deregister/hgw
[root@node19062652 yvh]# curl http://10.199.105.9:8500/v1/agent/services
{}
```


## 集群
1. 启动三个节点
```
./consul agent -server -bootstrap -ui -data-dir=/home/cons/data -bind=10.192.70.52 -client=0.0.0.0 &

./consul agent -server -ui -data-dir=/home/cons/data -bind=10.192.70.203 -client=0.0.0.0 &

./consul agent -server -ui -data-dir=/home/cons/data -bind=10.192.70.244 -client=0.0.0.0 &
```

2. 加入集群
```
./consul join 10.192.70.52
./consul join 10.192.70.244
```

3. 查看members
```
[root@vm-node-70-203 cons]# ./consul members
Node            Address             Status  Type    Build  Protocol  DC   Segment
vm-node-70-203  10.192.70.203:8301  alive   server  1.5.0  2         dc1  <all>
vm-node-70-244  10.192.70.244:8301  alive   server  1.5.0  2         dc1  <all>
vm-node-70-52   10.192.70.52:8301   alive   server  1.5.0  2         dc1  <all>

```

```
http://10.192.70.52:8500/ui/dc1/services/consul
```

4. consul client 模式
```
[root@localhost consul_1.5.0_linux_amd64]# cat client.json
{
  "data_dir": "./data",
  "enable_script_checks": true,
  "bind_addr": "10.199.105.9",
  "retry_join": ["10.192.70.52","10.192.70.203", "10.192.70.244"],
  "retry_interval": "30s",
  "rejoin_after_leave": true,
  "start_join": ["10.192.70.52"] ,
  "node_name": "node1"
}

./consul agent -config-dir=client.json &


[root@localhost consul_1.5.0_linux_amd64]# ./consul members
Node            Address             Status  Type    Build  Protocol  DC   Segment
vm-node-70-203  10.192.70.203:8301  alive   server  1.5.0  2         dc1  <all>
vm-node-70-244  10.192.70.244:8301  alive   server  1.5.0  2         dc1  <all>
vm-node-70-52   10.192.70.52:8301   alive   server  1.5.0  2         dc1  <all>
node1           10.199.105.9:8301   alive   client  1.5.0  2         dc1  <default>

```

## 远端访问
```
[root@localhost consul_1.5.0_linux_amd64]# ./consul members -http-addr='10.192.70.52:8500'
Node            Address             Status  Type    Build  Protocol  DC   Segment
vm-node-70-203  10.192.70.203:8301  alive   server  1.5.0  2         dc1  <all>
vm-node-70-244  10.192.70.244:8301  alive   server  1.5.0  2         dc1  <all>
vm-node-70-52   10.192.70.52:8301   alive   server  1.5.0  2         dc1  <all>

```


## 开发模式
```
./consul agent -dev -config-dir=./consul.d -advertise=10.199.105.9 -bind=0.0.0.0 -client=0.0.0.0 -ui
```
