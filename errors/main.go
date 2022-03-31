package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name, Email, role string
	age int
}

func (u User) Salary() (int, error) {
	if (u.role == "") {
		return 0, errors.New("Undefined role")
	}   

	switch u.role {
		case "Full-stack developer": 
			return 100, nil;
		case "Architect": 
			return 200, nil
	}
	return 0, errors.New("I don't know this role")
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func main() {
	user := User {
		Name: "Cláudio",
		role: "Unknow role",
		age: 30,
	}

	user.UpdateEmail("claudinhoandre@gmail.com")
	fmt.Println(user)
	fmt.Println(user.Name)
	salary, err := user.Salary()

	if (err != nil) {
		fmt.Println(err)
	} else {
		fmt.Println(salary)
	}

}
