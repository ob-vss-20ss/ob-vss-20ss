package stringutil

import (
	"fmt"
	"testing"
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
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

// Example
func ExampleReverse() {
	fmt.Print(Reverse("Hallo Welt"))
	// output:
	// tleW ollaH
}

// Benchmark
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("ljhdsagjlhfasdgljfadsgfljhasgfljahgfdsajlhgafsdljh")
	}
}
