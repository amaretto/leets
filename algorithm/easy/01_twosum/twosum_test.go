package main

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	Nums   []int
	Target int
	Result []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
	{[]int{1, 3, 4, 2}, 6, []int{2, 3}},
}

func TestTwoSum(t *testing.T) {
	for _, testCase := range testCases {
		result := TwoSum(testCase.Nums, testCase.Target)
		if !reflect.DeepEqual(result, testCase.Result) {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			TwoSum(testCase.Nums, testCase.Target)
		}
	}
}

func TestTwoSumHash(t *testing.T) {
	for _, testCase := range testCases {
		result := TwoSumHash(testCase.Nums, testCase.Target)
		if !reflect.DeepEqual(result, testCase.Result) {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func BenchmarkTwoSumHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			TwoSumHash(testCase.Nums, testCase.Target)
		}
	}
}

func TestTwoSumOnePass(t *testing.T) {
	for _, testCase := range testCases {
		result := TwoSumOnePass(testCase.Nums, testCase.Target)
		if !reflect.DeepEqual(result, testCase.Result) {
			t.Errorf("invalid result. testCaes:%v, actual:%v", testCase, result)
		}
	}
}

func BenchmarkTwoSumOnePass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			TwoSumOnePass(testCase.Nums, testCase.Target)
		}
	}
}
