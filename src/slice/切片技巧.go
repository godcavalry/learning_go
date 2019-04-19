package main

import "fmt"

func main() {
	/*
		append 向切片中追加元素
	*/
	a := []int{1, 2, 3}
	// 向a切片中追加多个元素
	a = append(a, 4, 5, 6)
	fmt.Println("向a切片中追加多个元素:", a)

	// 向a切片中追加切片
	b := []int{7, 8, 9}
	a = append(a, b...)
	fmt.Println("向a切片中追加切片:", a)

	/*
		copy深拷贝一个切片到另外一个切片中
	*/
	c1 := []string{"a", "b", "c", "d", "e", "f"}
	c2 := make([]string, len(c1))

	// 复制c1的元素到c2中，修改c2不影响c1的值
	copy(c2, c1)
	fmt.Println("c1的容量", cap(c1))
	fmt.Println("c2的容量", cap(c1))
	fmt.Println("c2修改前c1的数据：", c1)
	fmt.Println("c2修改前c2的数据：", c2)
	c2[0] = "test"
	fmt.Println("c2修改后c1的数据：", c1)
	fmt.Println("c2修改后c2的数据：", c2)

	// copy函数复制nil切片时，前后数据格式不一致
	var c3 []int
	c4 := make([]int, len(c3))
	copy(c4, c3)
	fmt.Println("c3是否为nil：", c3 == nil)
	fmt.Println("c4是否为nil：", c4 == nil)

	// 利用append复制切片
	c5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	c6 := append(c5[:0:0], c5...)
	fmt.Println("c5的容量", cap(c5))
	fmt.Println("c6的容量", cap(c6))
	fmt.Println("c5修改前")
	fmt.Println("c5的数据：", c5)
	fmt.Println("c6的数据：", c6)
	c5[0] = 4
	fmt.Println("c5修改后")
	fmt.Println("c5的数据：", c5)
	fmt.Println("c6的数据：", c6)
	// append函数复制nil切片时， 前后格式一致
	var c7 []int
	c8 := append(c7[:0:0], c7...)
	fmt.Println("c7是否为nil：", c7 == nil)
	fmt.Println("c8是否为nil：", c8 == nil)

	// 切片剪切
	str := []string{"a", "b", "c", "d", "e"}
	fmt.Println("str 剪切前：", str)
	str = append(str[:1], str[3:]...)
	fmt.Println("str 剪切后：", str)

	d1 := []int{1, 2, 3, 4, 5}
	fmt.Println("1号索引元素删除前：",d1)
	d1 = append(d1[:1], d1[2:]...)
	fmt.Println("1号索引元素删除后：",d1)
	fmt.Println(3/2-1)


}
