package common

import (
	"passwordcracker/cracker"
	wordlist2 "passwordcracker/lib/wordlist"
)

type Case struct {
	ExpectedPassword string
	Hash             string
}

func GenerateTestCases(n int) []Case {
	wordlist := wordlist2.ParseWordlist(".")
	var cases []Case
	for i := 0; i < n; i++ {
		hash := cracker.HashString(wordlist[i])
		cases = append(cases, Case{
			ExpectedPassword: wordlist[i],
			Hash:             hash,
		})
	}
	return cases
}
