package queuegate

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Talk(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
