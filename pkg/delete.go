package pkg

import (
	"fmt"
)

func (ex *Executer) isDelete(lang *SkikLang) bool {

	if len(lang.Tokens) >= 3 && lang.Tokens[0].Type == DELETE && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING {
		return true
	}
	return false
}

func (ex *Executer) delete(lang *SkikLang) error {
	_, f := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if !f {
		return fmt.Errorf("Key:%s not exists :/", lang.Tokens[2].Value.(string))
	}
	delete(ex.Db.Storage.Datad, lang.Tokens[2].Value.(string))
	return nil
}
