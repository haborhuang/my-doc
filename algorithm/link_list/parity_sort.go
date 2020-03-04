package main

import "fmt"

// 奇数位元素升序，偶数位元素降序的单向链表，对链表排序
func main() {
	l := initLink(5, 40, 15, 20, 25, 10)

	// 思路一：按奇偶位拆分，反转偶数位链表，最后拼接两个链表
	// l1, l2 := splitByParity(l)
	// l2 = reverse(l2)
	// join(l1, l2).print()

	// 优化：在拆分时，顺便进行偶数位链表的反转
	l1, l2 := splitByParityAndReverseByEven(l)
	join(l1, l2).print()
}

type link struct {
	val  int
	next *link
}

func (l *link) print() {
	for cur := l; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

// 根据奇偶性拆分，并反转偶数位链表
func splitByParityAndReverseByEven(l *link) (*link, *link) {
	var (
		h1 *link // 奇数位链表头
		l1 *link // 当前奇数位链表节点

		h2 *link // 偶数位链表头
		t2 *link // 偶数位链表尾
	)

	for i, cur := 0, l; cur != nil; {
		i++
		if i%2 == 0 {
			if h2 == nil {
				// 定位到最后一个节点
				h2, t2 = cur, cur
				cur = cur.next
				continue
			}
			// cur.next指向当前head，head更新为cur
			// cur后移
			cur, cur.next, h2 = cur.next, h2, cur
		} else {
			if h1 == nil {
				h1, l1 = cur, cur
			} else {
				l1.next = cur
				l1 = l1.next
			}
			cur = cur.next
		}
	}
	// 清理尾部
	if l1 != nil {
		l1.next = nil
	}
	if t2 != nil {
		t2.next = nil
	}

	h1.print()
	h2.print()

	return h1, h2
}

// 根据奇偶性拆分
func splitByParity(l *link) (*link, *link) {
	var l1, l2, h1, h2 *link
	for i, cur := 0, l; cur != nil; cur = cur.next {
		i++
		if i%2 == 0 {
			if h2 == nil {
				h2, l2 = cur, cur
			} else {
				l2.next = cur
				l2 = l2.next
			}
		} else {
			if h1 == nil {
				h1, l1 = cur, cur
			} else {
				l1.next = cur
				l1 = l1.next
			}
		}
	}
	if l1 != nil {
		l1.next = nil
	}
	if l2 != nil {
		l2.next = nil
	}

	return h1, h2
}

func initLink(nums ...int) *link {
	if len(nums) == 0 {
		return nil
	}
	head := &link{
		val: nums[0],
	}
	p := head
	for i := 1; i < len(nums); i++ {
		p.next = &link{
			val: nums[i],
		}
		p = p.next
	}
	return head
}

func reverse(l *link) *link {
	if l == nil {
		return nil
	}
	head, next := l, l.next
	head.next = nil
	for next != nil {
		next, next.next, head = next.next, head, next
	}

	return head
}

func join(l1 *link, l2 *link) *link {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var joined, cur *link
	p1, p2 := l1, l2
	if p1.val < p2.val {
		cur = p1
		p1 = p1.next
	} else {
		cur = p2
		p2 = p2.next
	}

	joined = cur
	for ; p1 != nil && p2 != nil; cur = cur.next {
		if p1.val < p2.val {
			cur.next = p1
			p1 = p1.next
		} else {
			cur.next = p2
			p2 = p2.next
		}
	}

	p := p1
	if p == nil {
		p = p2
	}

	for ; p != nil; p = p.next {
		cur.next = p
		cur = cur.next
	}

	return joined
}
