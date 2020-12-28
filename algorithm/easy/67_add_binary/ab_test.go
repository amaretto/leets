package main

import "testing"

var testCases = []struct {
	InputA string
	InputB string
	Output string
}{
	{"11", "1", "100"},
	{"1010", "1011", "10101"},
	{"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011", "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000"},
}

func TestAddBinary(t *testing.T) {
	for _, testCase := range testCases {
		result := addBinary(testCase.InputA, testCase.InputB)
		if result != testCase.Output {
			t.Errorf("invalid result. testCase:%v, actual:%v", testCase, result)
		}
	}
}
