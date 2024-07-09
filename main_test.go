package main

import (
	"passwordcracker/cracker"
	"testing"
)

func TestCrackSha1Hash(t *testing.T) {
	tests := []struct {
		name    string
		hash    string
		want    string
		wantErr bool
	}{
		{"Ytrewq", "some_sha1_hash_for_ytrewq", "ytrewq", false},
		{"Batman", "some_sha1_hash_for_batman", "batman", false},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := cracker.CrackSha1Hash(tt.hash, false)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrackSha1Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if actual != tt.want {
				t.Errorf("CrackSha1Hash() = %v, want %v", actual, tt.want)
			}
		})
	}
}
