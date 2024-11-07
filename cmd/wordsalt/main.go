package main

import (
	"fmt"

	"github.com/iamcryptoki/wordsalt"
)

func main() {
	keys := wordsalt.GenerateWordPressKeys()

	for _, name := range keys {
		fmt.Printf("define('%s', '%s');\n", name, keys[name])
	}
}
