package UserController

import (
	"howToUseMod/modules/users"
	"log"
)

func Show(Id int) users.User {
	user, err := users.GetUser(Id)

	if err != nil {
		log.Fatalln(err)
	}

	return user
}
