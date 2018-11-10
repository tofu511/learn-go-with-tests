package integers

import "testing"

func TestAdder(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("adding 1 and 2 should be 3", func(t *testing.T) {
		sum := Add(1,2)
		expect := 3
		assertCorrectMessage(t, sum, expect)
	})

	t.Run("adding 10 and 20 should be 30", func(t *testing.T) {
		sum := Add(10, 20)
		expect := 30
		assertCorrectMessage(t, sum, expect)
	})

}
