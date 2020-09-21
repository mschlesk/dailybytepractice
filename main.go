package main

import (
	"fmt"

	"github.com/mschlesk/dailybytepractice/pkg/reversestring"
)

var result bool

func main() {
	m := map[bool]string{true: "✅", false: "❌"}

	welcome := "=== DailyBytes Runner ===\n"
	fmt.Println(welcome)

	fmt.Println("Running reversestring tests...")
	result = reversestring.Test()
	fmt.Printf("Results: %s\n\n", m[result])
}
