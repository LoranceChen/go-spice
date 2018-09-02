package main

import (
	"fmt"
	"strconv"
)

type Skills []string

type Human struct {
	name string
	age int
	phone string  // Human类型拥有的字段
}

type Employee struct {
	Human  // 匿名字段Human
	Human2
	Skills
	speciality string
	phone string  // 雇员的phone字段
	int
}

type Human2 struct {
	two string
	name string
}

func main() {
	Bob := Employee{
		Human{"Bob", 34, "777-444-XXXX"},
		Human2{"two", "name2"},
		Skills{"s1", "s2"},
		"Designer",
		"333-222",
		int(1),
	}

	c := Bob.int
	str :=strconv.Itoa(c)
	println("ccc: " + str)
	a := Bob.Skills[0]
	print(a)

	fmt.Println("Bob's work phone is:", Bob.phone)
	// 如果我们要访问Human的phone字段
	fmt.Println("Bob's personal phone is:", Bob.Human.phone)
}