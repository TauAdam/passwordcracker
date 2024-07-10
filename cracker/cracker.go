package cracker

import (
	"crypto/sha1"
	"fmt"
	"passwordcracker/lib/wordlist"
)

func CrackSha1Hash(wordlist []string, hash string, withSalt bool) (string, error) {
	for _, word := range wordlist {
		if withSalt {
			for _, salted := range HashWithSalt(word) {
				if salted == word {
					return salted, nil
				}
			}
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

func HashWithSalt(str string) []string {
	var s []string
	for _, salt := range wordlist.ParseSaltList() {
		s = append(s, HashString(str+salt))
		s = append(s, HashString(salt+str))
	}
	return s
}
