package cracker

import (
	"os"
	"strings"
)

func CrackSha1Hash(hash string, withSalt bool) (string, error) {
	return "PASSWORD NOT IN DATABASE"
}

func parseWordlist() []string {
	b, err := os.ReadFile("./top-10k-passwords.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
func parseSaltList() []string {
	b, err := os.ReadFile("./known-salts.txt")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(b), "\n")
}
