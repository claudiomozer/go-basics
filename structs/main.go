package main

import "fmt"

type User struct {
	Name, Email, role string
	age int
}

func (u User) Salary() int {
	switch u.role {
		case "Full-stack developer": return 100;
		case "Ar'chtect": return 200
	}
	return 0
}

func (u *User) UpdateEmail(email string) {
	u.Email = email
}

func main() {
	user := User {
		Name: "Cláudio",
		role: "Full-stack developer",
		age: 30,
	}

	user.UpdateEmail("claudinhoandre@gmail.com")
	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Salary())
}
