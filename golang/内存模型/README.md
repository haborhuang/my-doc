参考 https://blog.csdn.net/ywh147/article/details/10942603

# happens before机制

* A happens before B 表示A修改的内存对B可见。
* A happens before B 并不代表A在B之前执行，因为代码的执行顺序和书写的逻辑顺序并不会完全一致。
* 单个goroutine，happens before顺序与代码顺序一致。

# init

* p包 import了 q包，则q的init happens before p包执行。
* main方法 happens after所有包的init。

# goroutine

* goroutine创建 happens before goroutine执行
* goroutine执行完成没有happens before保证

# channel规则

* A send on a channel happens before the corresponding receive from that channel completes.
* The closing of a channel happens before a receive that returns a zero value because the channel is closed.
* A receive from an unbuffered channel happens before the send on that channel completes.
* The kth receive on a channel with capacity C happens before the k+Cth send from that channel completes.

* 发送 happens before 读取完成
* channel关闭 happens before 读取到closed
* 无缓冲channel的读取 happens before 发送完成
* 第k次读取容量为C的channel happens before 第k+C次发送完成

# mutax

sync.Mutex 或 sync.RWMutex的unlock() 一定 happens before 后续lock()返回

# Once

单例f()的执行 happens before 所有once.Do(f)返回