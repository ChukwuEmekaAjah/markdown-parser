package md

import (
	"testing"
)

// lists parsing
var lines []string = []string{
	"> Hello cutie",
	"> Hello world ",
	">  : I am a bad ass",
	"> >Hi world",
	">: Are you a bad ass",
	": Yes I am",
	">Hello girls",
	": I am smart",
	"## ahelo",
	"Hello cutie",
	"Hello world ",
	": I am a bad ass",
	"Hi world",
	"> : Are you a bad ass",
	"> > >1. : Yes I am",
	">2. Hello girls",
	": I am smart",
	"## ahelo",
}

var results []string = []string{
	"<blockquote> Hello cutie",
	" Hello world ",
	"  : I am a bad ass",
	"<blockquote>Hi world</blockquote>",
	": Are you a bad ass</blockquote>",
	": Yes I am",
	"<blockquote>Hello girls</blockquote>",
	": I am smart",
	"## ahelo",
	"Hello cutie",
	"Hello world ",
	": I am a bad ass",
	"Hi world",
	"<blockquote> : Are you a bad ass",
	"<blockquote> >1. : Yes I am</blockquote>",
	"2. Hello girls</blockquote>",
	": I am smart",
	"## ahelo",
}

func TestParseBlockQuotes(t *testing.T) {

	e := ParseBlockQuotes(lines)

	for i := range e {
		if e[i] != results[i] {
			t.Log("Incorrect block quote parsing", e[i], results[i])
			t.Fail()
		}
	}
}
