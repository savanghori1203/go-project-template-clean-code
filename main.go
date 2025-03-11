package main

import (
	"fmt"
)

// func init() {
// 	_, err := os.OpenFile("./courses.json", os.O_WRONLY|os.O_CREATE, 0644)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	fmt.Println("Course CRUD in go lang")

	router := router()
	router.Run(":3000")
}
