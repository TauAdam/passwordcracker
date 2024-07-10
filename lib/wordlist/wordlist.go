package wordlist

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ParseWordlist(dir string) []string {
	b, err := os.ReadFile(filepath.Join(dir, "./top-10k-passwords.txt"))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\n")
}
func ParseSaltList() []string {
	b, err := os.ReadFile("./known-salts.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), "\n")
}
