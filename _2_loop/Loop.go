package main

import "fmt"

func main() {
	users := [...]string{"One", "Two", "Three"}

	for i := 0; i < len(users); i++ {
		print(users[i])
	}

	for _, value := range users {
		println(value)
	}

	for index, value := range users {
		fmt.Println(index, value)
	}

}
