package main

import "reflect"

type Person struct {
	name string
	age  int
}

func Set(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	}
}
