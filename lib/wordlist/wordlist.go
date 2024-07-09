package wordlist

import (
	"os"
	"strings"
)

func ParseWordlist() []string {
	b, err := os.ReadFile("./top-10k-passwords.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
func ParseSaltList() []string {
	b, err := os.ReadFile("./known-salts.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
