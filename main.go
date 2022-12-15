package main

import (
	"fmt"
	"relationship/config"

	student "relationship/services/belongs-to/services"
	class "relationship/services/many-to-many/services"
	club "relationship/services/many-to-many/services"
	user "relationship/services/one-to-many/services"
	customer "relationship/services/one-to-one/services"
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

	//many to many relationship normal
	club.CreateData()

	/*many to many back reference*/
	//One student have mny classes
	class.CreateOneSide()

	//One class have many students
	class.CreateOneSide()

}
