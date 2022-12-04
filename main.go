package main

import (
	"fmt"
	"reflect"
)

func Add(x, y int) int {
	return x + y
}

type A struct {
	m    int
	Name string
}

func (a A) SetName1(name string)  { a.Name = name }
func (a *A) SetName2(name string) { a.Name = name } // 修改字段的方法必须使用指针方法才会生效
func (a A) GetName1() string      { return a.Name } // 获取字段值的方法则无所谓
func (a *A) GetName2() string     { return a.Name }

func main() {
	// 直接调用
	aIns := &A{Name: "thy"}
	aIns.SetName1("thy1")
	fmt.Println(aIns.GetName1(), aIns.GetName2())
	aIns.SetName2("thy1")
	fmt.Println(aIns.GetName1(), aIns.GetName2())

	// 通过反射包调用
	aIns = &A{Name: "thy", m: 2}
	v := reflect.ValueOf(aIns) // 传入的是指针类型，所以返回的reflect.Value的Kind()是reflect.Ptr
	fmt.Println(v.Kind(), v.Elem().Kind())

	//fmt.Println(v.Elem().FieldByName("m").Interface()) // 触发panic

	// 指针类型方法只能通过指针类型调用
	v.Elem().MethodByName("SetName1").Call([]reflect.Value{reflect.ValueOf("thy1")})
	result1 := v.MethodByName("GetName1").Call([]reflect.Value{})
	result2 := v.MethodByName("GetName2").Call([]reflect.Value{})
	fmt.Println(result1[0].String(), result2[0].String())

	v.MethodByName("SetName2").Call([]reflect.Value{reflect.ValueOf("thy1")})
	result1 = v.MethodByName("GetName1").Call([]reflect.Value{})
	result2 = v.MethodByName("GetName2").Call([]reflect.Value{})
	fmt.Println(result1[0].String(), result2[0].String())
}
