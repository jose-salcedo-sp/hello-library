package greetings

import (
	"fmt"
	"math/rand"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello %s from Go!", name)
	return message
}

func RandomHello() string {
	greeting := []string{
		"Hey, you %v",
		"What roll with the chicken, %v",
		"Hello crayon, %v",
		"What's up my dogs %v",
		"Holis crayons %v",
	}
	return greeting[rand.Intn(len(greeting))]
}

func PersonalFunc() int {
	return rand.Intn(100)
}
