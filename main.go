package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := &Point{X: 10, Y: 5}
	rv := reflect.ValueOf(p).Elem()
	// fmt.Printf("type = %v\n", rv.Type())
	// fmt.Printf("kind = %v\n", rv.Kind())
	fmt.Printf("interface = %v\n", rv.Interface())

	xv := rv.Field(0)
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after) = %d\n", xv.Int())
}
