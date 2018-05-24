package dal

import (
	"fmt"

	"../domainModels"
)

type ILinkedList interface {
	Head() *domainModels.Point
	Last() *domainModels.Point
	Count() uint32
	Next() *domainModels.Point
	AddToEnd(value domainModels.Point) error
}

func NewILinkedList() ILinkedList {
	return &linkedList{nil, nil, nil, 0}
}

func NewILinkedListWithData(points []domainModels.Point) ILinkedList {
	list := &linkedList{nil, nil, nil, 0}

	for _, p := range points {
		list.AddToEnd(p)
	}
	return list
}

type node struct {
	value domainModels.Point
	next  *node
}

type linkedList struct {
	head    *node
	current *node
	last    *node
	count   uint32
}

func (this *linkedList) Head() *domainModels.Point {
	if this == nil {
		return nil
	}
	if this.head == nil {
		return nil
	}
	return &this.head.value
}

func (this *linkedList) Last() *domainModels.Point {
	if this == nil {
		return nil
	}
	if this.last == nil {
		return nil
	}
	return &this.last.value
}
func (this *linkedList) Count() uint32 {
	if this == nil {
		return 0
	}
	return this.count
}

func (this *linkedList) Next() *domainModels.Point {
	if this == nil {
		return nil
	}
	if this.head == nil {
		return nil
	}
	if this.current == nil {
		this.current = this.head
	}
	result := this.current.value
	this.current = this.current.next
	return &result
}

func (this *linkedList) AddToEnd(value domainModels.Point) error {
	if this == nil {
		return fmt.Errorf("linked list is nil!")
	}
	this.count += 1
	if this.head == nil {
		this.head = &node{value, nil}
		this.last = this.head
		this.current = this.head
		return nil
	}

	this.last.next = &node{value, nil}
	this.last = this.last.next
	return nil
}
