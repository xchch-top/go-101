package main

import (
	"fmt"
)

func main() {
	str := "a string value"
	pointer := &str
	anotherString := *&str
	// a string value
	fmt.Println(str)
	// 0xc00003c230
	fmt.Println(pointer)
	// a string value
	fmt.Println(anotherString)
	str = "changed string"
	// changed string
	fmt.Println(str)
	// 0xc00003c230
	fmt.Println(pointer)
	// a string value
	fmt.Println(anotherString)

	para := ParameterStruct{Name: "aaa"}
	// {aaa}
	fmt.Println(para)
	changeParameter(&para, "bbb")
	// {bbb}
	fmt.Println(para)
	cannotChangeParameter(para, "ccc")
	// {bbb}
	fmt.Println(para)
}

type ParameterStruct struct {
	Name string
}

func changeParameter(para *ParameterStruct, value string) {
	para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}
