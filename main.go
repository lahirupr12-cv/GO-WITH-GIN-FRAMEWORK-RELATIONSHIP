package main

import (
	"fmt"
	"relationship/config"

	student "relationship/services/belongs-to/services"
	customer "relationship/services/one-to-one/services"
	user "relationship/services/one-to-many/services"
)

func main() {

	fmt.Println("hello")
	config.Connection()

	//belongs to reltionship
	student.CreateData()

	//one to one relationship
	customer.CreateData()

	//one to many relationship
	user.CreateData()

}
