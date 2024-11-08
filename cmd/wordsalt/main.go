package main

import (
	"fmt"

	"github.com/iamcryptoki/wordsalt"
)

func main() {
	keys, err := wordsalt.GenerateWordPressKeys()
	if err != nil {
		fmt.Println("Error generating keys:", err)
		return
	}

	for _, name := range keys {
		fmt.Printf("define('%s', '%s');\n", name, keys[name])
	}
}
