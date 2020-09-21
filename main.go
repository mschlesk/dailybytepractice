package main

import (
	"fmt"

	"github.com/mschlesk/dailybytepractice/pkg/reversebyte"
)

func main() {
	m := map[bool]string{true: "Pass", false: "Fail"}

	welcome := "=== DailyBytes Runner ===\n"
	fmt.Println(welcome)

	fmt.Printf("Running reverse byte tests... %s\n", m[reversebyte.Test()])
}
