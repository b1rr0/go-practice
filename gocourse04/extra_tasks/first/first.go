package first

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)

	for _, word := range strings.Split(s, " ") {
		res[word]++
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
