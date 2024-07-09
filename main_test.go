package main

import (
	"passwordcracker/cracker"
	"testing"
)

func TestCrackSha1Hash(t *testing.T) {
	cases := []struct {
		expectedPassword string
		hash             string
	}{
		{"ytrewq", "22837024f941f67c2ff80c49e6bccf110c062149"},
		{"Batman", "32b26a271530f105cbc35cb653110e1a49d019b6"},
		{"123456", "7c4a8d09ca3762af61e59520943dc26494f8941b"},
		{"password", "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8"},
		{"12345678", "7c222fb2927d828af22f592134e8932480637c0d"},
		{"qwerty", "b1b3773a05c0ed0176787a4f1574ff0075f7521e"},
		{"123456789", "f7c3bc1d808e04732adf679965ccc34ca7ae3441"},
		{"12345", "8cb2237d0679ca88db6464eac60da96345513964"},
		{"1234", "7110eda4d09e062aa5e4a390b0a572ac0d2c0220"},
		{"111111", "3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d"},
		{"1234567", "20eabe5d64b0e216796e834f52d61fd0b70332fc"},
		{"dragon", "af8978b1797b72acfff9595a5a2a373ec3d9106d"},
		{"123123", "601f1889667efaebb33b8c12572835da3f027f78"},
		{"baseball", "a2c901c8c6dea98958c219f6f2d038c44dc5d362"},
		{"abc123", "6367c48dd193d56ea7b0baad25b19455e529f5ee"},
		{"football", "2d27b62c597ec858f6e7b54e7e58525e6a95e6d8"},
		{"monkey", "ab87d24bdc7452e55738deb5f868e1f16dea5ace"},
		{"letmein", "b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3"},
		{"696969", "cedf41fccb586dc39e1ce34bb482f0afe557b49f"},
		{"shadow", "ed9d3d832af899035363a69fd53cd3be8f71501c"},
		{"master", "4f26aeafdb2367620a393c973eddbe8f8b846ebd"},
		{"666666", "1411678a0b9e25ee2f7c8b2f7ac92b6a74b3f9c5"},
	}

	for _, tc := range cases {
		t.Run(tc.expectedPassword, func(t *testing.T) {
			actual, _ := cracker.CrackSha1Hash(tc.hash, false)
			if actual == tc.expectedPassword {
				t.Logf("PASS: %s", tc.hash)
			} else {
				t.Errorf("FAIL: %s", tc.hash)
			}
		})
	}
}
