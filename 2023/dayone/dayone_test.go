package dayone

import (
	"testing"
)

func TestPartonetest(t *testing.T) {

	got := Part1("input_test_1.txt")
	want := 142

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestPartone(t *testing.T) {

	got := Part1("input.txt")
	want := 55017

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestParttwotest(t *testing.T) {

	got := Part2("input_test_2.txt")
	want := 281

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestParttwotestfailure(t *testing.T) {

	got := Part2("input_failing.txt")
	want := 12

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestParttwo(t *testing.T) {

	got := Part2("input.txt")
	want := 53539

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
