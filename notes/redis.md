# cluster



## 显示集群中的节点

```shell
cluster nodes
```



## 添加节点

```shell
cluster meet <ip> <port>
```

### gossip协议传播



## 客户端连接集群

```shell
redis-cli -c -p <port>
```





## slot

Redis 通过分片的形式保存键值对：集群的数据库被分为16384个slot





## 集群中执行命令

节点收到与key相关的命令是会检查key是否指派到自己：

如果指派到自己，则直接执行

否则，放回MOVED错误



## MOVED error
格式： MOVED <slot> <ip>:<port>
集群客户端通过moved错误转发请求





