## admin 管理模块

### 1. 用户管理
用户信息保存在mysql数据库中,每个节点上运行一个admin服务和一个redis缓存.
admin与redis之间使用连接池.