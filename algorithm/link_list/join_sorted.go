package main

import "fmt"

// 将两个有序链表合并为一个有序链表
func main() {
	l1 := initLink(1, 3, 5, 7, 9)
	l2 := initLink(2, 4, 6)
	l1.print()
	l2.print()

	join(l1, l2).print()
}

type link struct {
	val  int
	next *link
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

func join(l1 *link, l2 *link) *link {
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

func (l *link) print() {
	for cur := l; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
