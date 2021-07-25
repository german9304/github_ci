package greeting

import (
	"testing"
)

func TestGreet(t *testing.T) {
	word := Hello("world")
	if word != "Hello world" {
		t.Errorf("got: %s, want: %s\n", word, "hello world")
	}
}
