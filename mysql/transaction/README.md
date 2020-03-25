# ACID 特性

* A，Atomicity，原子性。简单理解即事务中的操作要么都成功，要么都失败。
* C，Consistency，一致性。无论事务并发量多大，事务执行前后，系统仍保持一致状态。如账户间互相转账，转账事务结束之后，账户总余额不变。
* I，Isolation，隔离性。对同一数据处理的并发事务之间互不影响，执行效果如同串行执行。
* D，Durability，持久性。事务完成之后，更新会被持久的保存。

# 事务实现原理

InnoDB引擎事务实现原理参考 https://www.cnblogs.com/davygeek/p/7995072.html

## 原子性

### undo log

事务修改操作先记入undo log再执行操作。可以简单理解为，每个修改操作在undo log里都会记录对应的逆向修改。事务回滚时执行undo log。

### 事务状态

* Active：事务的初始状态，表示事务正在执行；
* Partially Commited：在最后一条语句执行之后；
* Failed：发现事务无法正常执行之后；
* Aborted：事务被回滚并且数据库恢复到了事务进行之前的状态之后；
* Commited：成功执行整个事务；

## 持久性

### redo log

* 数据修改时，会生成一条重做日志记录写入缓冲区。提交事务时，缓冲区内容写入日志文件，再将修改数据持久化。
* redo log以块形式存储，块大小与磁盘扇区大小相同，从而保证了重做日志写入的原子性。
* undo log的持久化也使用了redo log机制。
* 数据持久化前发生故障时，重启后可通过redo log重新写入数据。

## 隔离性

为提高性能，数据库允许事务并行处理，因此需要机制保证并行事务间的隔离性。

### 事务并发的问题

* 脏读：事务A读取了事务B更新的数据，然后B回滚操作，那么A读取到的数据是脏数据
* 不可重复读：事务 A 多次读取同一数据，事务 B 在事务A读取的过程中，对数据作了更新并提交，导致事务A多次读取同一数据时，结果不一致。
* 幻读：系统管理员A将数据库中所有学生的成绩从具体分数改为ABCDE等级，但是系统管理员B就在这个时候插入了一条具体分数的记录，当系统管理员A改结束后发现还有一条记录没有改过来，就好像发生了幻觉一样，这就叫幻读。

### 事务隔离级别

* read-uncommitted。
* read-committed。可解决脏读的问题。写数据会锁住行。
* repeatable-read。默认的隔离级别，可解决脏读和不可重复读的问题。
* serializable。读写会锁表，可解决脏读、不可重复读和幻读的问题。

隔离级别越高，越能保证数据的完整性和一致性，但是对并发性能的影响也越大。

RR级别通过next-key锁解决幻读问题。

#### REPEATABLE READ

事务中的第一次读操作会建立快照，后续读操作结果会基于快照来保证事务中读操作的一致性。

SELECT ... LOCK IN SHARE MODE、SELECT ... FOR UPDATE、UPDATE、DELETE语句使用以下锁：

* 在唯一索引上使用结果唯一的搜索条件，InnoDB引擎只锁住查找到的index，不会锁住index前的间隙。
* 使用间隙锁或next-key lock锁住被扫描的index范围，阻塞其他事务在其中插入数据。

#### READ COMMITTED

事务中每次读都设置及刷新各自快照。SELECT ... LOCK IN SHARE MODE、SELECT ... FOR UPDATE、UPDATE、DELETE语句只会锁住index记录而不是间隙。

UPDATE语句在该级别下使用“semi-consistent”读：如果行已被锁住，从而检查最新提交的版本数据是否满足WHERE条件，满足则再次读取并对这些行加锁或阻塞。

UPDATE、DELETE语句会在WHERE语句执行后立即释放不满足条件行的锁，而不是等到事务结束。

### 隔离级别的实现

MySQL通过 锁 + MVCC 实现。

#### MVCC

多版本数据可以让事务不等待写锁释放而读取数据，可显著提升读操作性能。

每行记录有两个隐藏字段：创建时间和删除时间，保存对应操作的事务id。

* insert时，创建时间保存当前事务id。
* delete时，删除时间保存当前事务id。
* update时，删除时间保存当前事务id，并添加新记录，其创建时间保存当前事务id。
* 查询时，返回满足以下条件的row：
  * 创建时间小于等于当前事务id
  * 删除时间未定义或大于当前事务id

# 参考资料

* https://www.jianshu.com/p/4e3edbedb9a8
* https://www.cnblogs.com/huanongying/p/7021555.html
* https://www.cnblogs.com/protected/p/6526857.html
* https://www.cnblogs.com/aipiaoborensheng/p/5767459.html
* https://www.cnblogs.com/dongqingswt/p/3460440.html
* https://dev.mysql.com/doc/refman/8.0/en/innodb-locking.html
