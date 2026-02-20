package Hello_modules

import (
	"fmt"
	"math/rand"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello, %v form hello1", name)
	return message
}

func RandomHello() string {
	greetings := []string{
		"Hey, you %v",
		"What roll with the chickens, %v",
		"What's up my dogs, %v",
		"Holis crayons, %v",
	}
	return greetings[rand.Intn(len(greetings))]
}

func LoQueQuiera() string {
	return "Lo que quiera"
}
