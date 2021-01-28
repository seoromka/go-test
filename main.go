package main

import "fmt"
import "math/rand"
import "time"
import "strings"

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func kratn(number int, divider int) bool {
	return number % divider == 0
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var result strings.Builder
	number := random(1, 100)
	criteria := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	fmt.Printf("Number: %d\n", number)

	for divider, value := range criteria {
		if kratn(number, divider) {
			result.WriteString(value)
		}
	}

	if result.Len() == 0 {
		result.WriteString(fmt.Sprintf("%d", number))
	}
	
	fmt.Printf("Result: %s\n", result.String())
}
