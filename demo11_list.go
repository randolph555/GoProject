package main

import (
	"container/list"
	"fmt"
)

func main() {
	linkendList()
	funcLinkendList()
}

//双向链表List
func funcLinkendList() {
	l := list.New()

	//尾部添加元素
	l.PushBack("chanon")

	//头部添加元素
	l.PushFront(11)

	//尾部添加后保存元素句柄
	element := l.PushBack("fist")

	//在fish之后添加high元素
	l.InsertAfter("hign", element)

	//在fist之前添加noon
	l.InsertBefore("noon", element)

	//移除元素
	l.Remove(element)
	fmt.Println("l", l.Len(), l.Front().Value)

	//迭代列表元素
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

//声明 双向链表List
func linkendList() {
	list1 := list.New()
	var list2 list.List
	fmt.Println(list1, list2)
}
