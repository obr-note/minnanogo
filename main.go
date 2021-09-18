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
	rv := reflect.ValueOf(p)
	fmt.Printf("type = %v\n", rv.Type())
	fmt.Printf("kind = %v\n", rv.Kind())
	fmt.Printf("interface = %v\n", rv.Interface())
}
