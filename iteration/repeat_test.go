package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T)  {

	t.Run("Repeat('a', 5) should return 'aaaaa' ", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("Repeat('f', 5) should return 'fffff' ", func(t *testing.T) {
		repeated := Repeat("f", 5)
		expected := "fffff"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})

	t.Run("Repeat('a', 3) should return 'aaa'", func(t *testing.T) {
		repeated := Repeat("a", 3)
		expected := "aaa"

		if repeated != expected {
			t.Errorf("expected '%s' but got '%s'", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("d", 10)
	fmt.Println(repeated)
	// Output: dddddddddd
}
