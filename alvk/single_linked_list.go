package main

import (
	"fmt"
)

type SingleNode struct {
	data interface{}
	next *SingleNode
}

type SingleLinkedList struct {
	head *SingleNode
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) Add(data interface{}) {
	node := &SingleNode{data, nil}
	if l.head == nil {
		l.head = node
	} else {
		temp := l.head
		l.head = node
		node.next = temp
	}
}

func (l *SingleLinkedList) For(a func(data interface{})) {
	node := l.head
	for node != nil {
		a(node.data)
		node = node.next
	}
}

func singleLinkedListAddTest() {
	list := NewSingleLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.For(func(data interface{}) { fmt.Println(data) })
}

func (l *SingleLinkedList) Reverse() {
	var prev *SingleNode = nil
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
}

func singleLinkedListReverseTest() {
	list := NewSingleLinkedList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)
	list.Reverse()
	list.For(func(data interface{}) { fmt.Println(data) })
}

func SingleLinkedListTest() {
	//singleLinkedListAddTest()
	singleLinkedListReverseTest()
}
