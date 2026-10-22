// ============================================================
//  TEST SUITE - Do not modify
// ============================================================

package main

import "fmt"

type testCase struct {
	input    []int
	expected int
}

func main() {
	tests := []testCase{
		{[]int{1, 2, 3, 4}, 20},
		{[]int{-5, 10, -3, 7}, 34},
		{[]int{0, 1, -1, 2}, 6},
		{[]int{}, 0},
		{[]int{-1, -2, -3}, 0},
		{[]int{100, 200, 300}, 1200},
	}

	fmt.Println()
	fmt.Println("🧪 Running tests...")
	fmt.Println()

	passed := 0
	for i, t := range tests {
		result := TransformArray(t.input)
		ok := result == t.expected
		if ok {
			passed++
		}
		status := "❌"
		if ok {
			status = "✅"
		}
		fmt.Printf("%s Test %d: got %d, expected %d\n", status, i+1, result, t.expected)
	}

	fmt.Printf("\n📊 %d/%d passed\n", passed, len(tests))
	if passed == len(tests) {
		fmt.Println("🎉 All passed!")
	} else {
		fmt.Println("⚠️ Some failed. Keep going!")
	}
}
