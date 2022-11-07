package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (Part: %s) - index: waiting (%d) <> found (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text    string
		part    string
		waiting int
	}{
		{"Alberto Parente", "Juliana Cavalcante", 0},
		{"", "", 0},
		{"Hello", "hello", -1},
		{"Alberto", "Parente", 2},
	}

	for _, test := range tests {
		t.Logf("Ok: %v", test)
		actual := strings.Index(test.text, test.part)

		if actual != test.waiting {
			t.Errorf(msgIndex, test.text, test.part, test.waiting, actual)
		}
	}
}
