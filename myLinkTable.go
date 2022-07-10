package main

import "fmt"

//链表结点定义
type ListNode struct {
	val  int
	next *ListNode
}

//链表定义
type LinkList struct {
	head   *ListNode
	length int
}

//创建链表
func CreateLinkList() LinkList {
	return LinkList{&ListNode{}, 0}
}

//在指定位置添加元素
func (this *LinkList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.head.next = &ListNode{val, this.head.next}
		this.length++
		return
	} else if index == this.length {
		p := this.head
		for p.next != nil {
			p = p.next
		}
		p.next = &ListNode{val, nil}
		this.length++
		return
	} else if index > this.length {
		return
	}

	p := this.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := &ListNode{val, nil}
	q.next = p.next
	p.next = q
	this.length++
}

//删除指定位置元素
func (this *LinkList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}
	p := this.head
	for i := 0; i < index; i++ {
		p = p.next
	}
	q := p.next
	p.next = q.next
	q.next = nil
	this.length--
}

//遍历链表
func (this *LinkList) PrintList() {
	p := this.head.next
	for p != nil {
		fmt.Print(p.val)
		fmt.Print(" ")
		p = p.next
	}
	fmt.Println()
}

//test
func main() {
	linklist := CreateLinkList()
	linklist.AddAtIndex(0, 1)
	linklist.AddAtIndex(0, 3)
	linklist.AddAtIndex(1, 4)
	linklist.AddAtIndex(1, 2)
	linklist.PrintList()
	linklist.DeleteAtIndex(1)
	linklist.PrintList()
}
