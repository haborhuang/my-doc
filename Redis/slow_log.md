
Slow Log特性用来记录执行时间超过指定值的查询，执行时间不包括网络io，仅为指令执行时间。Slow Log只在内存中维护。

* slowlog-log-slower-than参数指定阈值，单位为微秒。指定负数时关闭log，指定0时记录所有指令。
* slowlog-max-len参数指定日志大小。可以认为使用先入先出队列实现。
* 配置文件和CONFIG SET都可以设置参数

# 日志查看

SLOWLOG GET N指令用来查看最近N条日志内容。

输出示例：
```
redis 127.0.0.1:6379> slowlog get 2
1) 1) (integer) 14
   1) (integer) 1309448221
   2) (integer) 15
   3) 1) "ping"
2) 1) (integer) 13
   1) (integer) 1309448128
   2) (integer) 30
   3) 1) "slowlog"
      1) "get"
      2) "100"
```

4.0版本之后还会输出以下内容：
```
5) "127.0.0.1:58217"
6) "worker-123"
```

字段含义：
1. 日志记录id
2. 指令执行时间戳
3. 指令执行时间，微秒
4. 指令参数
5. 客户端ip和端口（仅4.0版本）
6. 客户端别名（仅4.0版本）

使用SLOWLOG LEN指令可获取Slow Log长度。使用SLOWLOG RESET指令可重置日志，日志信息将被丢弃。