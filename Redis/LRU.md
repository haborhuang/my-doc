参考http://www.redis.cn/topics/lru-cache.html

需要配置maxmemory来限制内存占用上限。

# 淘汰策略

配置项为maxmemory-policy，策略有：
* noeviction。内存达到上限时，申请内存的指令会返回错误。
* allkeys-lru。lru淘汰。
* volatile-lru。设置有效期的keys中，lru淘汰。
* allkeys-random。随机淘汰。
* volatile-random。设置有效期的keys中，随机淘汰。
* volatile-ttl。设置有效期的keys中，回收有效期最早的key。

注：如果没有有效期的keys，volatile-lru、volatile-random和volatile-ttl与noeviction相似。

## 工作原理

1. client执行某一指令，且导致内存增加。
2. Redis校验内存使用情况。如果超过maxmemory限制则根据配置策略淘汰一些keys。