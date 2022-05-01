package tool

import (
	"fmt"
	"reflect"
)

// 利用反射打印结构体
func SmartPrint(i interface{}) {
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)

	if vValue.Kind() == reflect.Ptr {
		vValue = vValue.Elem()
	}
	if vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}

	for i := 0; i < vValue.NumField(); i++ {
		fmt.Printf("%-30v%-50v \n", vType.Field(i).Name, vValue.Field(i))
	}
	fmt.Println()
}
