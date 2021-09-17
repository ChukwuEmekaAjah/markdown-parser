package md

import "testing"

var linesUl []string = []string{
	"- Hello cutie",
	"- Hello world ",
	"-  : I am a bad ass",
	"-Hi world",
	"- : Are you a bad ass",
	": Yes I am",
	"Hello girls",
}

var linesOl []string = []string{
	": I am smart",
	"## ahelo",
	"Hello cutie",
	"Hello world ",
	": I am a bad ass",
	"Hi world",
	": Are you a bad ass",
	"1. : Yes I am",
	"2. Hello girls",
	": I am smart",
	"## ahelo",
}

var resultUl []string = []string{
	"<ul><li> Hello cutie</li>",
	"<li> Hello world </li>",
	"<li>  : I am a bad ass</li>",
	"<li>Hi world</li>",
	"<li> : Are you a bad ass</li></ul>",
	": Yes I am",
	"Hello girls",
}

var resultOl []string = []string{
	": I am smart",
	"## ahelo",
	"Hello cutie",
	"Hello world ",
	": I am a bad ass",
	"Hi world",
	": Are you a bad ass",
	"<ol><li> : Yes I am</li>",
	"<li> Hello girls</li></ol>",
	": I am smart",
	"## ahelo",
}

func TestParseUnorderedList(t *testing.T) {

	e := ParseUnorderedList(linesUl)

	for i := range e {
		if e[i] != resultUl[i] {
			t.Log("Incorrect list item parsing", e[i], resultUl[i])
			t.Fail()
		}
	}
}

func TestParseOrderedList(t *testing.T) {

	e := ParseOrderedList(linesOl)

	for i := range e {
		if e[i] != resultOl[i] {
			t.Log("Incorrect list item parsing", e[i], resultOl[i])
			t.Fail()
		}
	}
}
