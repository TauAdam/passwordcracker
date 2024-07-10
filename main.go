package main

import (
	"fmt"
	"passwordcracker/common"
)

func main() {
	testCases := common.GenerateTestCases(20)
	for _, testCase := range testCases {
		fmt.Printf("{%q, %q},\n", testCase.ExpectedPassword, testCase.Hash)
	}
}
