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
* 事务。参考mysql/transaction/目录内容
* 主从。参考mysql/master-slave/目录内容。
* 锁。参考mysql/lock/目录内容。

# Redis

* 持久化。参考Redis/Redis持久化.md
* 有效期。参考Redis/key有效期.md
* LRU及淘汰策略。参考Redis/LRU.md
* 主从。
* 集群。

