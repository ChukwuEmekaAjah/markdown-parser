package md

import (
	"fmt"
	"testing"
)

func TestReplacer(t *testing.T) {
	str := "## ~~***ajah***~~ ~~**i*s**~~ cool ooo [title](https://www.example.com) ![alt text](image.jpg) # chah"

	str2 := " * * * * * "
	e := Replacer(str)

	fmt.Println("new str", e)
	fmt.Println("=============================")
	fmt.Println("\n\nnew str2", Replacer(str2))
	if e != "h1 ajah is cool ooo h1 chah" {
		t.Log("Invalid function", e)
		t.Fail()
	}

}
