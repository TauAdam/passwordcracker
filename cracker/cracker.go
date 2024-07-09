package cracker

import (
	"crypto/sha1"
	"fmt"
	"passwordcracker/lib/wordlist"
)

func CrackSha1Hash(hash string, withSalt bool) (string, error) {
	for _, word := range wordlist.ParseWordlist() {
		if withSalt {

		} else if hashOfWord := HashString(word); hashOfWord == hash {
			return word, nil
		}
	}
	return "PASSWORD NOT IN DATABASE", nil
}

func HashString(str string) string {
	sum := sha1.Sum([]byte(str))
	return fmt.Sprintf("%x", sum)
}
