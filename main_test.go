package main

import "testing"

var input string = "omama"

func TestSortWord(t *testing.T) {
	var expected string = "aaomm"
	var output string

	output = SortWord(input)
	
	t.Logf("Word : %s\n", input)
	t.Logf("Output : %s\n", output)
	t.Logf("Expected : %s\n", expected)

	if (output != expected) {
		t.Errorf("Error! expected (%s) got (%s)", expected, output)
	}

}

func BenchmarkSortWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SortWord(input)
	}
}