package hello

import "testing"

func TestHello(t *testing.T) {
	expect := "hello, world"
	if received := Hello(); received != expect {
		t.Errorf("Hello() = %q, want %q", received, expect)
	}
}
