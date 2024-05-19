package pkg

import (
	"errors"
	"strings"
)

/*
skik lang example:
-
get all

get where key == "juan"
get key "juan"

*/

type SkikLang struct {
	Tokens     []Token
	SourceCode string
}

var (
	KEY_VALUE_NOT_FOUND = errors.New("key not found")
	TYPE_INVALID        = errors.New("invalid type")
)

func (skl *SkikLang) PreInit() {
	skl.SourceCode = strings.ReplaceAll(skl.SourceCode, "\r\n", "\n")
	skl.SourceCode = strings.TrimSpace(skl.SourceCode)
	skl.SourceCode += "\n"

}
func NewSkikLang(source string) *SkikLang {

	return &SkikLang{
		Tokens:     []Token{},
		SourceCode: source,
	}
}

//func (skl*SkikLang)
