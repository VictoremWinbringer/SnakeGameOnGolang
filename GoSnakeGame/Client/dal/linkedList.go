package dal
import "../domainModels"

type ILinkedList interface {
	Head() (domainModels.Point, bool)
	Count() uint32
	Next() (domainModels.Point , bool)
	AddToEnd(value domainModels.Point) error
	}
	
	func NewILinkedList() ILinkedList {
		return &linkedList{nil,nil,nil,0}
	}
	
	type node struct {
	value domainModels.Point
	next *node
	}
	
	type linkedList struct {
		head *node
		current *node	
		last *node
		count uint32
	}

	func (this *linkedList) Head() (domainModels.Point, bool) {
		if this == nil{
			return 0, false
		}
		if this.head == nil{
			return 0, false
		}
		return this.head.value,true
		}

		func (this *linkedList) Count() uint32 {
              if this == nil{
	          	return 0
	          }
	          return this.count
		}

		func (this *linkedList) Next() (domainModels.Point, bool) {
              if this == nil {
              	return 0, false
              }
              if this.head == nil{
              	return 0, false
              }              
              if this.current == nil{
              this.current = this.head
              return 0, false
	          } 
		      result := this.current.value
		      this.current = this.current.next
		      return result, true
	    }

		func (this *linkedList) AddToEnd(value domainModels.Point) error{
			if this == nil {
				return fmt.Errorf("linked list is nil!")
			}
			this.count +=1			
			if this.head == nil{
				this.head = &node {value,nil}
				this.last = this.head
				this.current = this.head
				return nil
			}

			this.last.next = &node{value,nil}
			this.last = this.last.next
			return nil
		}