package main

import (
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {

	// can NOT use copy to reset field value
	//var x float64 = 3.4
	//v := reflect.ValueOf(x)
	//v.SetFloat(7.1)

	var xx float64 = 3.4
	p := reflect.ValueOf(&xx)

	v2 := p.Elem()

	v2.SetFloat(7.1)
}

