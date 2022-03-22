package main

import (
	"errors"
	"fmt"
)

func main() {
	example1()
}

func example1() {
	name := "Tom"
	s, err := New(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s.Listen("english"))
	fmt.Println(s.Speak("english"))
	fmt.Println(s.Read("english"))
	fmt.Println(s.Write("english"))
}

//定义接口
type Study interface {
	Listen(msg string) string
	Speak(msg string) string
	Read(msg string) string
	Write(msg string) string
}

//结构体
type study struct {
	Name string
}

//一个 struct 实现了某个接口里的所有方法，就叫做这个 struct 实现了该接口 (必须是所有方法，否则编译不通过)
func (s *study) Listen(msg string) string {
	return s.Name + " 听 " + msg
}

func (s *study) Speak(msg string) string {
	return s.Name + " 说 " + msg
}

func (s *study) Read(msg string) string {
	return s.Name + " 读 " + msg
}

func (s *study) Write(msg string) string {
	return s.Name + " 写 " + msg
}

func New(name string) (Study, error) {
	if name == "" {
		return nil, errors.New("name required")
	}
	return &study{
		Name: name,
	}, nil
}

//调用器接口
type Invoker interface {
	Call(interface{})
}

//结构体类型
type Struct struct{}

//实现Invoker的Call
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

//函数定义为类型
type FuncCaller func(interface{})

//实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func exampe2() {
	//声明接口变量
	var invoker Invoker
	//实例化结构体,两种写法都可以，返回的都是指针类型
	//s := new(Struct)
	s := &Struct{}
	//将是实例化的结构体赋值到接口
	invoker = s
	//使用接口调用实例化结构体的方式Struct.Call
	invoker.Call("hello")
	//将匿名函数转为FuncCaller类型，在赋值给接口
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker.Call("hello")
}
