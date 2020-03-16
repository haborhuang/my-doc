package main

import "fmt"

// 二叉搜索树 转为 升序排列的双向链表
func main() {
	root := &tnode{
		val: 5,
		left: &tnode{
			val: 3,
			left: &tnode{
				val: 2,
				left: &tnode{
					val: 1,
				},
			},
			right: &tnode{
				val: 4,
			},
		},
		right: &tnode{
			val: 8,
			left: &tnode{
				val: 6,
				right: &tnode{
					val: 7,
				},
			},
			right: &tnode{
				val: 9,
			},
		},
	}

	dl := conv(root)
	dl.print()
	dl.printReverse()
}

type tnode struct {
	val   int
	left  *tnode
	right *tnode
}

type doubleLinkList struct {
	head *dlnode
	tail *dlnode
}

type dlnode struct {
	val  int
	prev *dlnode
	next *dlnode
}

func (l *doubleLinkList) print() {
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func (l *doubleLinkList) printReverse() {
	for cur := l.tail; cur != nil; cur = cur.prev {
		fmt.Print(cur.val)
		if cur.prev != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

// 采用中序遍历策略遍历二叉树
func conv(root *tnode) doubleLinkList {
	var dl doubleLinkList
	var l *dlnode
	cur := root
	st := newStack()

	for cur != nil || !st.isEmpty() {
		if cur != nil {
			st.push(cur)   // cur压栈
			cur = cur.left // cur指向左子树
			continue
		}

		// 获取栈顶元素并访问
		cur = st.pop()
		{
			if dl.head == nil {
				l = &dlnode{
					val: cur.val,
				}
				dl.head = l
			} else {
				// 双向链表添加节点
				l.next = &dlnode{
					val:  cur.val,
					prev: l,
				}
				// 双向链表节点指针后移
				l = l.next
			}
			dl.tail = l
		}

		cur = cur.right // cur指向右子树
	}
	return dl
}

func newStack() *stack {
	return &stack{}
}

type stack struct {
	st []*tnode
}

func (s *stack) isEmpty() bool {
	return len(s.st) == 0
}

func (s *stack) push(n *tnode) {
	s.st = append(s.st, n)
}

func (s *stack) pop() *tnode {
	if s.isEmpty() {
		return nil
	}
	n := s.st[len(s.st)-1]
	s.st = s.st[:len(s.st)-1]
	return n
}
