package main

import (
	"fmt"
	"relationship/config"
	student "relationship/services/belongs-to/services"
)

func main() {

	fmt.Println("hello")
	config.Connection()
	student.CreateData()

}
