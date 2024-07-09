package main

import (
	"fmt"
	"passwordcracker/cracker"
	wordlist2 "passwordcracker/lib/wordlist"
)

func main() {
	testCases := generateTestCases(20)
	for _, testCase := range testCases {
		fmt.Printf("{%q, %q},\n", testCase.expectedPassword, testCase.hash)
	}
}

type Case struct {
	expectedPassword string
	hash             string
}

func generateTestCases(n int) []Case {
	wordlist := wordlist2.ParseWordlist()
	var cases []Case
	for i := 0; i < n; i++ {
		hash := cracker.HashString(wordlist[i])
		cases = append(cases, Case{
			expectedPassword: wordlist[i],
			hash:             hash,
		})
	}
	return cases
}
