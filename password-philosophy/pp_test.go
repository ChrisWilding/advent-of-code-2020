package main

import "testing"

type isValidPasswordTest struct {
	input   string
	wantOne bool
	wantTwo bool
}

var isValidPasswordCases = []isValidPasswordTest{
	{
		input:   "1-3 a: abcde",
		wantOne: true,
		wantTwo: true,
	},
	{
		input:   "1-3 b: cdefg",
		wantOne: false,
		wantTwo: false,
	},
	{
		input:   "2-9 c: ccccccccc",
		wantOne: true,
		wantTwo: false,
	},
}

func TestCountValidPasswordsOne(t *testing.T) {
	want := 2
	got := CountValidPasswords([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, IsValidPasswordOne)
	if got != want {
		t.Errorf("CountValidPasswordsOne() = %v; want %v", got, want)
	}
}

func TestIsValidPasswordOne(t *testing.T) {
	for _, test := range isValidPasswordCases {
		got := IsValidPasswordOne(test.input)
		if got != test.wantOne {
			t.Errorf("IsValidPasswordOne() = %v; want %v", got, test.wantOne)
		}
	}
}

func TestCountValidPasswordsTwo(t *testing.T) {
	want := 1
	got := CountValidPasswords([]string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}, IsValidPasswordTwo)
	if got != want {
		t.Errorf("CountValidPasswordsTwo() = %v; want %v", got, want)
	}
}

func TestIsValidPasswordTwo(t *testing.T) {
	for _, test := range isValidPasswordCases {
		got := IsValidPasswordTwo(test.input)
		if got != test.wantTwo {
			t.Errorf("IsValidPasswordTwo() = %v; want %v", got, test.wantTwo)
		}
	}
}
