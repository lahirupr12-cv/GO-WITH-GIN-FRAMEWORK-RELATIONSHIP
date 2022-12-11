package main

import (
	"fmt"
	"relationship/config"
	student "relationship/services/one-to-one/service"
)

func main() {

	fmt.Println("hello")
	config.Connection()
	student.CreateData()

}
