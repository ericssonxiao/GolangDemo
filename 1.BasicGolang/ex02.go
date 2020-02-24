package main

import "fmt"

//array
func myArray() {
	arr := []int{1, 2, 3, 4, 5}
	//arr2 := []int{5, 4, 3, 2, 1}
	fmt.Println(arr)
}

func myString() {

}

func main() {
	fmt.Printf("hello %s\n", "world")
	fmt.Println("escaped\tchar")
	fmt.Println("hello\nworld")
	fmt.Println("C:\\Users\\Administrator")
	fmt.Println("xiao说\"你好\"")
	fmt.Println("xiao说\r你好")

	m1 := 1
	m2 := 2

	fmt.Println(int16(m1 + m2))

	ss := "This is a test"
	fmt.Println(ss)

	myArray()
	myString()
}
