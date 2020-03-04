[简介](https://www.cnblogs.com/qingyunzong/p/9004509.html)
[架构](https://www.cnblogs.com/qingyunzong/p/9004593.html)

# 概念

* topic。消息类别。物理上topic数据是存在不同broker上的，但生产者和消费者只需要关心topic即可。
* partition。
  * 相当于分片的概念。一个topic可配置多个partition，消息数据在不同partition中不会重复，同一partition可保证消息的顺序。
  * 在Kafka中实现消费的方式是将日志中的分区划分到每一个消费者实例上，以便在任何时间，每个实例都是分区唯一的消费者。
* replica。一个消息会保存多个副本，实际是partition的副本。
* broker。kafka集群中的一个服务器节点。partition副本会物理存储在broker上。
* consumer group。一个消息只会被group中的一个consumer消费。
* leader。一个partition的多个replica中只会有一个leader，读写操作都在该节点上进行。
* follower。被动的同步leader上的数据。当leader宕机了，followers 中的一台服务器会自动成为新的 leader
* offset：
  * 分区中的每一个记录都会分配一个id号来表示顺序，我们称之为offset，offset用来唯一的标识分区中每一条记录。
  * consumer消费的位置信息，由consumer维护。

# 数据清除策略

可按时间和partition文件大小进行清除。

# 消息路由

发消息时根据partition机制确定消息存储到哪个partition。

1、 指定了 patition，则直接使用；
2、 未指定 patition 但指定 key，通过对 key 的 value 进行hash 选出一个 patition
3、 patition 和 key 都未指定，使用轮询选出一个 patition。

# delivery guarantee

## Exactly once

指生产者只生产一次消息。生产者可指定key发送消息，当出现故障，生产者可进行幂等性重试。

## At most once/at least one

consumer读取消息后需进行commit，该操作会更新consumer在当前partition上的offset。如果未commit，则下次读取依然会从上一次的offset读取。即实现了at least one。

如果消息处理的同时进行异步commit，或处理前commit，则实现了at most once。

# 高可用

同一个Partition可能会有多个Replica，之间选出一个Leader，Producer和Consumer只与这个Leader交互，其它Replica作为Follower从Leader中复制数据。

## replica分布策略

1.将所有Broker（假设共n个Broker）和待分配的Partition排序
2.将第i个Partition分配到第（i mod n）个Broker上
3.将第i个Partition的第j个Replica分配到第（(i + j) mode n）个Broker上

## 消息同步策略

* 消息发送到该Partition的Leader
* Leader会将该消息写入其本地Log。
* 每个Follower都从Leader pull数据，Follower在收到该消息并写入其Log后，向Leader发送ACK。
* Leader收到了ISR中的所有Replica的ACK，该消息就被认为已经commit了，Leader将增加HW（high watermark，最后 commit 的 offset）并且向Producer发送ACK。
* 为了提高性能，每个Follower在接收到数据后就立马向Leader发送ACK，而非等到数据写入Log中。Follower可以批量的从Leader复制数据，这样极大的提高复制性能（批量写磁盘），极大减少了Follower与Leader的差距。

## ISR

broker存活的判断条件：
* 一是它必须维护与ZooKeeper的session（这个通过ZooKeeper的Heartbeat机制来实现）。
* 二是Follower必须能够及时将Leader的消息复制过来，不能“落后太多”。

落后太多的判断条件：
* Follower复制的消息落后于Leader后的条数超过预定值
* 或Follower超过一定时间未向Leader发送fetch请求。

Leader会通过以上两点跟踪与其保持同步的Replica列表，该列表称为ISR（即in-sync Replica）。

## Leader Election算法

kafka没有使用分布式锁的方式选举，而是通过一个controller进行。所有broker都尝试创建 controller path，只有一个竞选成功并当选为 controller。

1、 controller 在 zookeeper 的 /brokers/ids/ 节点注册 Watcher，当 broker 宕机时 zookeeper 会 fire watch
2、 controller 从 /brokers/ids 节点读取可用broker 
3、 controller决定set_p，该集合包含宕机 broker 上的所有 partition 
4、 对 set_p 中的每一个 partition 
    4.1、 从/brokers/topics/\[topic\]/partitions/\[partition\]/state 节点读取 ISR 
    4.2、 决定新 leader 
    4.3、 将新 leader、ISR、controller_epoch 和 leader_epoch 等信息写入 state 节点
5、 通过 RPC 向相关 broker 发送 leaderAndISRRequest 命令