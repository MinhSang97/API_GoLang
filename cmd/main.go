package main

import (
	"app/framework"
	"app/payload"
	"fmt"
)

func main() {

	framework.Route()
	fmt.Println(payload.AddStudentRequest{})

}
