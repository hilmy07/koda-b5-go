package internals

import "fmt"

func Segitigasiku(n int) string {
	result := ""

	for i := 1; i <= n; i++ {
		for k := 1; k <= i; k++ {
			result = fmt.Sprintf("%s%s", result, " *")
		}
		result += "\n"
	}

	return result
}