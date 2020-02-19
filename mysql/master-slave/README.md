可参考 https://baijiahao.baidu.com/s?id=1617888740370098866&wfr=spider&for=pc

# 同步机制简述

* master 维护bin-log。
* slave读取bin-log内容记入relay-log。再根据relay-log执行SQL完成更新。
* 三个thread：
  * log dump thread。master启动的线程，slave读取时启动，将bin-log内容发送给slave。
  * IO thread。slave启动的线程，连接master读取bin-log，并写入relay-log。
  * SQL thread。slave启动的线程，检测relay-log变更，并执行SQL完成更新。

# 同步模式

* 异步同步。默认模式。
* 半同步。提交事务后，master会等待一个从节点确认relay-log写入后再响应。如果确认超时，则切换为异步模式。

# bin-log格式

* Statement-base Replication (SBR)。记录SQL语句。优点是日志量较小，缺点时一些语句会使主从不一致（比如now()等）。
* Row-based Relication(RBR)。将SQL解析为对Row更改的语句。会解决数据不一致问题，但日志体量较大。
* Mixed-format Replication(MBR)。对两者的结合，根据情况确定保存哪种日志。

# GTID

简单说是全局唯一的事务ID，由server_id和transaction_id组成。transaction_id在事务提交时分配的不重复的序列号。

5.6之前slave需要记录上一次读取的bin-log position。引入GTID之后，bin-log中会记录GTID，slave在开启bin-log条件下不再需要记录position，而是通过GTID判断SQL是否执行。

# 主从延迟

参考 https://blog.csdn.net/hao_yunfeng/article/details/82392261

## 原因

* 单线程处理。master dump线程和slave io线程对文件顺序写，效率较高。slave SQL线程随机写，效率较低。
* 延迟场景：
  * master并发高，负载高时，slave SQL单线程处理不及时。
  * slave负载高时，slave SQL线程计算资源不足导致处理速度降低。
  * slave慢查询产生锁。
  * 网络延迟或服务器性能差等硬件问题。

## 优化

* 优化技术架构，降低数据库单点压力。
* 消除慢查询。
* 提升硬件性能。
* 优化MySQL配置选项。
* 并行复制。5.7版本后支持，参考https://www.cnblogs.com/xiaotengyi/p/5532191.html