参考http://www.redis.cn/commands/expire.html

expire 指令设置超时时间

# 过期数据淘汰

* 被动方式。key超时后再执行修改指令时才会被删除。
* 主动方式。每秒10次执行：
  1. 随机检测20个设置有效期的keys
  2. 删除过期的keys
  3. 如果过期的数量超过25%，则重复步骤1。

