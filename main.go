package main

import (
	"fmt"
	"reflect"
)

func main() {
	var1 := 3
	var2 := var1
	var3 := &var2

	fmt.Println("variables initalised")
	fmt.Printf("var1  : %d, type : %s, memory adress : %p\n", var1, reflect.TypeOf(var1), &var1)
	fmt.Printf("var2  : %d, type : %s, memory adress : %p\n", var2, reflect.TypeOf(var2), &var2)
	fmt.Printf("var3  : %d, type : %s, memory adress : %p\n", *var3, reflect.TypeOf(*var3), &*var3)

	var2++

	fmt.Println("\nvariables changed")
	fmt.Printf("var1  : %d, type : %s, memory adress : %p\n", var1, reflect.TypeOf(var1), &var1)
	fmt.Printf("var2  : %d, type : %s, memory adress : %p\n", var2, reflect.TypeOf(var2), &var2)
	fmt.Printf("var3  : %d, type : %s, memory adress : %p\n", *var3, reflect.TypeOf(*var3), &*var3)

	*var3 = 6
	// var3 = 6 <- won't compile
	fmt.Println("\ndeference made")
	fmt.Printf("var1  : %d, type : %s, memory adress : %p\n", var1, reflect.TypeOf(var1), &var1)
	fmt.Printf("var2  : %d, type : %s, memory adress : %p\n", var2, reflect.TypeOf(var2), &var2)
	fmt.Printf("var3  : %d, type : %s, memory adress : %p\n", *var3, reflect.TypeOf(*var3), &*var3)

	fmt.Printf("var3 type : %s", reflect.TypeOf(*var3))
}
