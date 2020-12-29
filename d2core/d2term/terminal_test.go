package d2term

import (
	"fmt"
	"testing"
)

func TestTerminal(t *testing.T) {
	term, err := NewTerminal()
	if err != nil {
		t.Fatal(err)
	}

	lenOutput := len(term.outputHistory)

	const expected1 = 2
	if lenOutput != expected1 {
		t.Fatalf("got %d expected %d", lenOutput, expected1)
	}

	if err := term.Execute("clear"); err != nil {
		t.Fatal(err)
	}

	if err := term.Execute("ls"); err != nil {
		t.Fatal(err)
	}

	lenOutput = len(term.outputHistory)

	const expected2 = 3
	if lenOutput != expected2 {
		t.Fatalf("got %d expected %d", lenOutput, expected2)
	}
}

func TestBind(t *testing.T) {
	term, err := NewTerminal()
	if err != nil {
		t.Fatal(err)
	}

	term.Clear()

	if err := term.Bind("hello", "world", []string{"world"}, func(args []string) error {
		const expected = "world"
		if args[0] != expected {
			return fmt.Errorf("got %s expected %s", args[0], expected)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	if err := term.Execute("hello world"); err != nil {
		t.Fatal(err)
	}
}

func TestUnbind(t *testing.T) {
	term, err := NewTerminal()
	if err != nil {
		t.Fatal(err)
	}

	if err := term.Unbind("clear"); err != nil {
		t.Fatal(err)
	}

	term.Clear()

	if err := term.Execute("ls"); err != nil {
		t.Fatal(err)
	}

	lenOutput := len(term.outputHistory)

	const expected = 2
	if lenOutput != expected {
		t.Fatalf("got %d expected %d", lenOutput, expected)
	}
}
