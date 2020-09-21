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
	if s == "" {
		return ""
	}

	rs := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		rs[i] = s[len(s)-i-1]
	}

	return string(rs)
}

// Test runs given and extracurricular test cases against proposed solution
func Test() bool {
	tests := map[string]string{
		"Cat":            "taC",
		"The Daily Byte": "etyB yliaD ehT",
		"civic":          "civic",
		"":               "",
	}

	pass := true
	for test, expected := range tests {
		fmt.Printf("\treverse(\"%s\") == \"%s\"?", test, expected)

		actual := reverse(test)
		if actual == expected {
			fmt.Println(" ✅")
		} else {
			fmt.Printf(" ❌\n\tReceived %s\n", actual)
			pass = false
		}
	}

	return pass
}
