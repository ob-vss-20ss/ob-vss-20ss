package stringutil_test

import (
	"fmt"
	"testing"

	"github.com/ob-vss-20ss/ob-vss-20ss/stringutil"
)

// Test
func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello", "olleH"},
		{"hallo", "ollah"},
	}

	for _, c := range cases {
		got := stringutil.Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// Example
func ExampleReverse() {
	fmt.Print(stringutil.Reverse("Hallo Welt"))
	// output:
	// tleW ollaH
}

// Benchmark
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringutil.Reverse("ljhdsagjlhfasdgljfadsgfljhasgfljahgfdsajlhgafsdljh")
	}
}
