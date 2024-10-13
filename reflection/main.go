package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{"Alice", 30}

    // Get Type information
    t := reflect.TypeOf(p)
    fmt.Println("Type:", t)

    // Get Value information
    v := reflect.ValueOf(p)
    fmt.Println("Value:", v)
    

    // Iterate over struct fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        tag := field.Tag.Get("json")
        fmt.Printf("%s: %v (tag: %s)\n", field.Name, value.Interface(), tag)
    }

    // Create a new struct using reflection
    newStruct := reflect.New(t).Elem()
    newStruct.Field(0).SetString("Bob")
    newStruct.Field(1).SetInt(25)
	fmt.Println("newStruct type:", reflect.TypeOf(newStruct))
    fmt.Println("New struct:", newStruct.Interface())
	fmt.Println("New struct type:", reflect.TypeOf(newStruct.Interface().(Person)))
	fmt.Println("p type:", reflect.TypeOf(p))
}