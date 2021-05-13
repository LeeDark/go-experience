package main

import (
	"flag"
	"log"

	"github.com/tomogoma/generator"
)

func init() {
	flag.Parse()
}

func main () {
	apiKey := generateApiKey()
	log.Println("API Key:", apiKey)
}

func generateApiKey() string {
	charSet, _ := generator.NewCharSet(generator.AlphaNumericChars)
	randChars, _ := charSet.SecureRandomBytes(32)
	return string(randChars)
}
