package main

import (
	"fmt"
)

func main() {
	//基础()
	f := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(两数之和(f, 9))
	//var a *单链表
	//a = 头插法()
	var a *单链表
	a = 头插法()
	//遍历单链表(*a)
	//fmt.Println(*a)
	遍历单链表(a)
}

//func 基础() {
//	var a int = 1
//	b := 2
//	var c int
//	var d [3]int
//	e := [3]int{1, 2, 3}
//	fmt.Println(a, b, c, d, e)
//
//	f := []int{4, 5, 6, 7} //切片
//	for k, v := range f {  //利用range来遍历数组或者切片
//		fmt.Println(k, v)
//	}
//	for i := 0; i < len(f); i++ {
//		fmt.Println(f[i])
//	}
//	fmt.Println(f[0:]) //可以直接打印数组的某个段
//	add(&a, b)
//	fmt.Println(a)
//}
//func add(x *int, y int) { //指针
//	*x = *x + y
//}
func 两数之和(nums []int, target int) (int, int) {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return i, j
			}
		}
		count = i
	}
	// if i==j && len(nums){} 提示j未被定义
	if count == len(nums) {
		return -1, -1
	} else {
		return 0, 0
	}
} //可以利用哈希表来求

type 单链表 struct {
	data int
	netx *单链表
}

//func 遍历单链表(x *单链表) {
//	head := x
//	if head.netx != nil {
//		fmt.Println(head.data)
//		head = head.netx
//	}
//}
func 遍历单链表(x *单链表) {
	var head *单链表
	head = x.netx
	for head != nil {
		fmt.Println(head.data)
		head = head.netx
	}
}
func 头插法() *单链表 {
	var head, a *单链表
	head = new(单链表) //少掉这行，会报错invalid memory address or nil pointer dereference
	var x int
	for x != -1 {
		a = new(单链表)
		fmt.Scan(&x)
		if x == -1 {
			break
		}
		a.data = x
		a.netx = head.netx
		head.netx = a
	}
	return head
}

//func 数组() []int {
//	var i, j int
//	j = 0
//	//var k []int
//	var k []int = make([]int, 5)
//	for i != -1 {
//		fmt.Scan(&i)
//		k[j] = i
//		j = j + 1
//	}
//	return k
//}

//func 两数相加(x 单链表, y 单链表) 单链表 {

//}
