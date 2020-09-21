package reversestring

import "fmt"

/*
This question is asked by Google. Given a string, reverse all of its characters and return the resulting string.

Ex: Given the following strings...

"Cat", return "taC"
"The Daily Byte", return "etyB yliaD ehT"
"civic", return "civic"
*/

func reverse(s string) string {
	return ""
}

// Test runs given and extracurricular test cases against proposed solution
func Test() bool {
	pass := true
	tests := map[string]string{"Cat": "taC", "The Daily Byte": "etyB yliaD ehT", "civic": "civic"}

	for test, expected := range tests {
		fmt.Printf("\treverse(%s) == %s?", test, expected)

		if reverse(test) == expected {
			fmt.Println(" ✔️")
		} else {
			fmt.Println(" ❌")
			pass = false
		}

	}
	return pass
}
