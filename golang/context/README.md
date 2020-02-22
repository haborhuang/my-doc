可参考源码

# 定义

```
type Context interface{
    // Deadline返回context有效期截止时间，ok代表是否设置了有效期
    Deadline() (deadline time.Time, ok bool)

    // Done返回只读channel，用来接收context已取消的信号。
    // 信号通过close channel实现。
    Done() <-chan struct{}

    // context被取消的原因。
    Err() error

    // context中获取指定value
    Value(key interface{}) interface{}
}
```

通常来说使用Done即可判断context是否已取消。

# 生成context

* context.Background()。通常用于main函数、测试函数、初始化或最顶层调用处。
* context.TODO()。当不清楚用什么context的时候，用此方式代替。
* context.WithCancel(parent Context)。返回新建的子context及其cancel方法。
* WithDeadline(parent Context, d time.Time)。返回新建的子context及其cancel方法。超过指定期限后，该子context自动取消。
* WithTimeout(parent Context, timeout time.Duration)。WithDeadline的语法糖。

# 实现原理

## emptyCtx 

```
// 是Background和TODO的数据结构。未设置有效期、无法取消、无key-value。
type emptyCtx int

// Done返回的是nil channel。意味着直接读取会永久阻塞。
func (*emptyCtx) Done() <-chan struct{} {
	return nil
}
```

## cancelCtx

cancelCtx是取消context的核心。其定义如下

```
type cancelCtx struct {
	Context // 父context

	mu       sync.Mutex            // protects following fields
	done     chan struct{}         // Done channel
	children map[canceler]struct{} // 子context
	err      error                 // 取消原因
}
```

WithCancel会返回一个cancelCtx，其取消方法如下：

```
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	if err == nil {
		panic("context: internal error: missing cancel error")
	}
    // 已取消
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return 
	}

    // 关闭Done channel以发送信号
	c.err = err
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}

    // cancel所有子context
	for child := range c.children {
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

    // 从父节点中移除
	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

WithCancel中还调用了propagateCancel，该函数是级联取消的关键：

```
func propagateCancel(parent Context, child canceler) {
    // 父节点已取消，则什么都不做
	if parent.Done() == nil {
		return // parent is never canceled
	}

	if p, ok := parentCancelCtx(parent); ok {
        // 如果父节点是cancelCtx
		p.mu.Lock()
		if p.err != nil {
            // 如果父节点已取消，则直接取消子节点
			child.cancel(false, p.err)
		} else {
            // 将子节点与父节点关联
			if p.children == nil {
				p.children = make(map[canceler]struct{})
			}
			p.children[child] = struct{}{}
		}
		p.mu.Unlock()
	} else {
        // 如果父节点不是cancelCtx，创建goroutine等待context被取消
		go func() {
			select {
			case <-parent.Done():
                // 如果父节点先取消，则取消子节点
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}
```

## timerCtx

```
type timerCtx struct {
	cancelCtx // 复用cancelCtx的父context，子context映射及Done channel
	timer *time.Timer // 有效期计时器

	deadline time.Time
}
```

WithDeadline会返回一个timerCtx，其取消方法如下：

```
func (c *timerCtx) cancel(removeFromParent bool, err error) {
    // 取消cancelCtx
	c.cancelCtx.cancel(false, err) 

    // 从父节点中移除当前节点
	if removeFromParent {
		removeChild(c.cancelCtx.Context, c)
	}

    // 停止计时器
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop() 
		c.timer = nil
	}
	c.mu.Unlock()
}
```

WithDeadline逻辑：

```
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
    // 父节点有效期较早时，与WithCancel等效
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
		return WithCancel(parent)
	}
	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	propagateCancel(parent, c)

    // 有效期已过，直接取消
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true, DeadlineExceeded) 
		return c, func() { c.cancel(false, Canceled) }
	}

    // 创建计时器，过有效期后调用取消
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {
			c.cancel(true, DeadlineExceeded)
		})
	}
	return c, func() { c.cancel(true, Canceled) }
}
```

## valueCtx

```
type valueCtx struct {
	Context // 父context
	key, val interface{} 
}

// value查找算法。
func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}
```

通过Value实现可以发现context并不适合存储大量的key-value，最坏情况下查找效率为O(n)。