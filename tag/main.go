package main

import (
	"log"
	"reflect"
)

func main() {
	num := 20

	reflectTest01(num)
}

func reflectTest01(b interface{}) {
	refType := reflect.TypeOf(b)
	log.Println("类型信息 ", refType)

	refVal := reflect.ValueOf(b)
	log.Println("值信息 ", refVal)

	num2 := refVal.Int() + 2
	log.Println("num2 ", num2)
}

type Person struct {
	Name string
	Age uint
}