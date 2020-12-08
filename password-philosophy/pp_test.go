package main

import "testing"

type isValidPasswordTest struct {
	input string
	want  bool
}

var isValidPasswordCases = []isValidPasswordTest{
	{
		input: "1-3 a: abcde",
		want:  true,
	},
	{
		input: "1-3 b: cdefg",
		want:  false,
	},
	{
		input: "2-9 c: ccccccccc",
		want:  true,
	},
}

func TestCountValidPasswords(t *testing.T) {
	want := 2
	got := CountValidPasswords([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})
	if got != want {
		t.Errorf("CountValidPasswords(2) = %v; want %v", got, want)
	}
}

func TestIsValidPassword(t *testing.T) {
	for _, test := range isValidPasswordCases {
		got := IsValidPassword(test.input)
		if got != test.want {
			t.Errorf("CountValidPasswords(2) = %v; want %v", got, test.want)
		}
	}
}
