package main

import "reflect"

func main() {

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	v.SetFloat(7.1)

	var xx float64 = 3.4
	p := reflect.ValueOf(&xx)

	v2 := p.Elem()

	v2.SetFloat(7.1)
}

