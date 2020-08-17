package d2lexer

import (
	"testing"
)

func TestName(t *testing.T) {
	lexer := New([]byte("correct horse battery staple andromeda13142 n1n2n4"))

	expected := []Token{
		{Name, "correct"},
		{Name, "horse"},
		{Name, "battery"},
		{Name, "staple"},
		{Name, "andromeda13142"},
		{Name, "n1n2n4"},
	}

	for _, want := range expected {
		got := lexer.NextToken()
		if got.Type != Name || got.Value != want.Value {
			t.Errorf("Got: %v, want %v", got, want)
		}
	}

	eof := lexer.NextToken()
	if eof.Type != EOF {
		t.Errorf("Did not reach EOF")
	}
}

func TestNumber(t *testing.T) {
	lexer := New([]byte("12 2325 53252 312 3411"))

	expected := []Token{
		{Number, "12"},
		{Number, "2325"},
		{Number, "53252"},
		{Number, "312"},
		{Number, "3411"},
	}

	for _, want := range expected {
		got := lexer.NextToken()
		if got.Type != Number || got.Value != want.Value {
			t.Errorf("Got: %v, want %v", got, want)
		}
	}

	eof := lexer.NextToken()
	if eof.Type != EOF {
		t.Errorf("Did not reach EOF")
	}
}

func TestSymbol(t *testing.T) {
	lexer := New([]byte("((+-==>>>=!=<=<=<*//*)?(::.,.:?"))

	expected := []Token{
		{Symbol, "("},
		{Symbol, "("},
		{Symbol, "+"},
		{Symbol, "-"},
		{Symbol, "=="},
		{Symbol, ">"},
		{Symbol, ">"},
		{Symbol, ">="},
		{Symbol, "!="},
		{Symbol, "<="},
		{Symbol, "<="},
		{Symbol, "<"},
		{Symbol, "*"},
		{Symbol, "/"},
		{Symbol, "/"},
		{Symbol, "*"},
		{Symbol, ")"},
		{Symbol, "?"},
		{Symbol, "("},
		{Symbol, ":"},
		{Symbol, ":"},
		{Symbol, "."},
		{Symbol, ","},
		{Symbol, "."},
		{Symbol, ":"},
		{Symbol, "?"},
	}

	for _, want := range expected {
		got := lexer.NextToken()
		if got.Type != Symbol || got.Value != want.Value {
			t.Errorf("Got: %v, want %v", got, want)
		}
	}

	eof := lexer.NextToken()
	if eof.Type != EOF {
		t.Errorf("Did not reach EOF")
	}
}

func TestString(t *testing.T) {
	lexer := New([]byte(`correct 'horse' 'battery staple' 'andromeda13142 ' n1n2n4`))

	expected := []Token{
		{Name, "correct"},
		{String, "horse"},
		{String, "battery staple"},
		{String, "andromeda13142 "},
		{Name, "n1n2n4"},
	}

	for _, want := range expected {
		got := lexer.NextToken()
		if got.Type != want.Type || got.Value != want.Value {
			t.Errorf("Got: %v, want %v", got, want)
		}
	}

	eof := lexer.NextToken()
	if eof.Type != EOF {
		t.Errorf("Did not reach EOF")
	}
}

func TestActualConstructions(t *testing.T) {
	lexer := New([]byte("skill('Sacrifice'.blvl) > 3 ? min(50, lvl) : skill('Sacrifice'.lvl) * ln12"))

	expected := []Token{
		{Name, "skill"},
		{Symbol, "("},
		{String, "Sacrifice"},
		{Symbol, "."},
		{Name, "blvl"},
		{Symbol, ")"},
		{Symbol, ">"},
		{Number, "3"},
		{Symbol, "?"},
		{Name, "min"},
		{Symbol, "("},
		{Number, "50"},
		{Symbol, ","},
		{Name, "lvl"},
		{Symbol, ")"},
		{Symbol, ":"},
		{Name, "skill"},
		{Symbol, "("},
		{String, "Sacrifice"},
		{Symbol, "."},
		{Name, "lvl"},
		{Symbol, ")"},
		{Symbol, "*"},
		{Name, "ln12"},
	}

	for _, want := range expected {
		got := lexer.NextToken()
		if got.Type != want.Type || got.Value != want.Value {
			t.Errorf("Got: %v, want %v", got, want)
		}
	}

	eof := lexer.NextToken()
	if eof.Type != EOF {
		t.Errorf("Did not reach EOF")
	}
}
