# 程序

## 死锁

可参考 https://www.cnblogs.com/dingpeng9055/p/11705870.html

# 算法

* LRU。可参考algorithm/lru/lru.go示例，使用map+双向链表实现
* 环形打印二维数组。可参考algorithm/two_dimensional_array/print_in_circle.go示例
* 接雨水。可参考algorithm/trapping_rain_water/trap.go示例。
* 找到第k小的值。可参考algorithm/kth_min.go示例
* 最长不重复子串。滑动窗口法，可参考algorithm/sliding_window/max_non_repetitive_substring.go示例
* 找到二叉树节点值之和的最大值。可参考algorithm/binary_tree/max_sum.go示例。
* 判断字符单向链表是否为回文。TODO：从中间位置拆分成两个链表，倒置其中一个后与另一个比较是否相同。
* 数组中某一元素出现概率超过50%，找出该元素。TODO：

# Golang

## 进程、线程、goroutine区别

[参考资料](https://www.cnblogs.com/ghj1976/p/3642513.html)。

* 进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。
* 线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是的)。
* 协程和线程一样共享堆，不共享栈，协程由程序员在协程的代码里显示调度。

## goroutine和channel解析

参考golang/channel_and_goroutine内容

channel可参考golang/Understanding channels.pdf

## GC

参考golang/gc/目录内容

## context

参考golang/context/目录内容

## 内存逃逸

参考golang/逃逸分析/目录内容

## 内存模型

参考golang/内存模型/目录内容

# 设计

## 红包系统设计

参考：
* https://www.infoq.cn/article/2017hongbao-weixin/ 
* https://blog.csdn.net/gb4215287/article/details/90112274

## 微博系统设计

参考：
* https://www.zhihu.com/question/19715683
* feed流系统设计：https://yq.aliyun.com/articles/706808

# MySQL

* 索引。参考mysql/index/目录内容。B+树延伸知识：https://zhuanlan.zhihu.com/p/27700617
* 事务。参考mysql/transaction/目录内容。
* 主从。参考mysql/master-slave/目录内容。
* 锁。参考mysql/lock/目录内容。

# Redis

* Redis高效的原因。参考Redis/为什么高效.md
* 哈希表实现原理。参考Redis/哈希表实现原理.md
* 持久化。参考Redis/Redis持久化.md
* 有效期。参考Redis/key有效期.md
* LRU及淘汰策略。参考Redis/LRU.md
* 主从。
* 集群方案。参考https://www.zhihu.com/question/21419897。

## 集群方案

通常是三种做法：
* Redis 3.0之后提供了集群方案，属于服务端sharding策略。
* 客户端sharding。即在客户端实现分片算法。
* 中间件。即客户端与服务端之间通过proxy代理。

### 扩容方案

* Redis 3.0集群方案扩容时只需要迁移部分slot到新增节点。
* 客户端sharding通常无法平滑扩容，可提前规划好节点数量，然后针对单个node垂直扩展，即提升单点容量。

延伸知识点。一致性哈希：https://www.cnblogs.com/lpfuture/p/5796398.html
* 哈希值空间虚拟成一个环。
* 服务节点经过哈希后会定位到环上具体位置。
* 数据的key经过哈希后也会定位到环上的某一点，其所属的节点即为顺时针查找的第一个节点。
* 节点过少时，可能造成数据分布不均的情况，此时可通过虚拟节点方式实现。即一个物理节点虚拟出若干逻辑节点（如节点名称+编号）分布到环上。

### Redis Cluster

* 3.0版本开始支持。
* 分片使用哈希槽技术：
  * 整个集群共16384个槽，集群搭建时分配到每个节点。
  * key经过哈希计算后会对应到一个槽，由拥有该槽的节点服务该key的相关请求。
  * 扩容时，只需要分配哈希槽到新增节点，数据转移较小。
* 集群总线：
  * 集群中每个节点会启动一个总线端口，集群总线用于故障检测、配置更新、故障转移授权等等。
* 高可用相关：
  * 需为每个master分配slave，master故障后，集群自动升级slave为master。
  * 一对master和slave均不可用时，整个集群也无法服务请求。
  * 副本迁移。集群自动发现slave的可用性，如某个master成为“孤儿”，且集群中存在master有多个slave时，集群会将分配一个slave给变成“孤儿”的master。
* 重定向：
  * MOVED。由于数据进行了分片，一个节点只能访问拥有的哈希槽中的数据。当server收到一个请求访问其他哈希槽的数据时，server会返回客户端MOVED错误，告知客户端应该访问哪个节点访问数据。
  * ASK：
    * ASK重定向用于哈希槽迁移期间。迁移过程中，新旧节点各拥有一部分数据，旧节点用ASK重定向告知client本次访问的key已转移到新节点。
    * client需以ASKING指令开始访问新节点，节点会校验以ASKING指令标记的请求可以访问迁移中的哈希槽。

# HTTP

## Cookie

参考https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Cookies

# 分布式

https://blog.csdn.net/fct2001140269/article/details/84503176

CAP原则：
* C 一致性，A 可用性，P 网络分区容错性。
* CAP原则简单理解为：在分布式系统中，三个特性无法同时满足。
* 多数情况下，网络分区时常发生，在系统设计中需要在C和A中做权衡。
* 一些场景要求强一致性（如银行），则需要在A和P中做权衡。

BASE：
* BA 基本可用，S 软状态，E 最终一直性。
* BA。出现异常故障时，允许损失部分可用性。如上游服务异常时，响应时间允许由0.5s变为1s；促销活动时，为保障购物系统稳定性，可将部分用户引导至降级页面。
* S。允许数据有中间状态，数据在节点间同步允许一定的延迟。
* E。不需要保证强一致性，而是经过一段延迟，数据在节点间最终一致。

* CAP和BASE原则。常见中间件对CAP的应用。
* ETCD & ZK 协议

# MQ

## RabbitMQ

简介：https://blog.csdn.net/fct2001140269/article/details/84503176

概念：
* exchange。对queue的路由规则
* binding。exchange与queue的绑定
* queue。消息队列
* 发布者：
  * 通常是面向exchange发送，exchange负责根据binding设置将消息路由到满足条件的queue。
  * 可以直接向指定的queue发送消息：不指定exchange，且routing key与queue name一致。
* queue会将消息均分给消费者。

exchange类型：
* fanout。用于广播。发消息时routing key会被忽略。消费者通过binding关联queue和exchange。只要建立binding，消息就会转发到queue。
* direct。routing key和binding key时，消息才会被转发到绑定的queue。
* topic。binding key可指定通配符，满足规则的消息会被转发到绑定的queue。

## Kafka

