package d2lexer

import (
	"testing"
)

func TestName(t *testing.T) {
	lexer := New([]byte("correct horse battery staple andromeda13142 n1n2n4"))

	expected := []Token{
		{"correct", Name},
		{"horse", Name},
		{"battery", Name},
		{"staple", Name},
		{"andromeda13142", Name},
		{"n1n2n4", Name},
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
		{"12", Number},
		{"2325", Number},
		{"53252", Number},
		{"312", Number},
		{"3411", Number},
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
		{"(", Symbol},
		{"(", Symbol},
		{"+", Symbol},
		{"-", Symbol},
		{"==", Symbol},
		{">", Symbol},
		{">", Symbol},
		{">=", Symbol},
		{"!=", Symbol},
		{"<=", Symbol},
		{"<=", Symbol},
		{"<", Symbol},
		{"*", Symbol},
		{"/", Symbol},
		{"/", Symbol},
		{"*", Symbol},
		{")", Symbol},
		{"?", Symbol},
		{"(", Symbol},
		{":", Symbol},
		{":", Symbol},
		{".", Symbol},
		{",", Symbol},
		{".", Symbol},
		{":", Symbol},
		{"?", Symbol},
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
		{"correct", Name},
		{"horse", String},
		{"battery staple", String},
		{"andromeda13142 ", String},
		{"n1n2n4", Name},
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
		{"skill", Name},
		{"(", Symbol},
		{"Sacrifice", String},
		{".", Symbol},
		{"blvl", Name},
		{")", Symbol},
		{">", Symbol},
		{"3", Number},
		{"?", Symbol},
		{"min", Name},
		{"(", Symbol},
		{"50", Number},
		{",", Symbol},
		{"lvl", Name},
		{")", Symbol},
		{":", Symbol},
		{"skill", Name},
		{"(", Symbol},
		{"Sacrifice", String},
		{".", Symbol},
		{"lvl", Name},
		{")", Symbol},
		{"*", Symbol},
		{"ln12", Name},
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
