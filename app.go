package main

import (
	"fmt"

	md "github.com/Bernardo258/10A/modules"
)

var test1 map[string]string

func main() {
	test1 = make(map[string]string)
	test1["test1"] = "beach"
	fmt.Println("Hello")
	//fmt.Println(md.pUser)
	fmt.Println(md.User)

}
