package vdn

import "testing"

func TestIsValidDomain(t *testing.T) {
	for i, tc := range []struct {
		domain string
		ok     bool
	}{
		{"www.google.com", true},
		{".www.google.com", false},
		{"google.mp3", false},
		{"google", false},
		{"com", false},
		{"g--oogle.com", true},
		{"google.com", true},
		{"google-.com", false},
		{"-google.com", false},
		{"google.-com", false},
		{"google..com", false},
	} {
		if IsValidDomain(tc.domain) != tc.ok {
			t.Errorf("test case %d:(%s, %v) failed\n", i, tc.domain, tc.ok)
		}
	}
}
