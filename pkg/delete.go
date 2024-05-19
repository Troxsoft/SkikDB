package pkg

import (
	"fmt"
)

func (ex *Executer) isDeleteWhere(lang *SkikLang) bool {
	if len(lang.Tokens) >= 5 && lang.Tokens[0].Type == DELETE && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == ALL && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == WHERE {
		return true
	}
	if len(lang.Tokens) >= 5 && lang.Tokens[0].Type == DELETE && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == NUMBER && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == WHERE {
		return true
	}
	return false
}

/*
	example:

get * where(key startsWith "+18") retorna las claves que empiezan con +18
*/
func (ex *Executer) deleteWhere(lang *SkikLang) map[string]any {
	keysVals := make(map[string]any)
	if lang.Tokens[2].Type == ALL {
		for k, v := range ex.Db.Storage.Datad {
			b := ex.whereDelete(k, v, lang)

			if b {
				keysVals[k] = v
				delete(ex.Db.Storage.Datad, k)
			}
		}
	} else {
		lolol := 0.0
		for k, v := range ex.Db.Storage.Datad {
			if lolol < lang.Tokens[2].Value.(float64) {
				b := ex.whereDelete(k, v, lang)
				if b {
					keysVals[k] = v
					lolol++
					delete(ex.Db.Storage.Datad, k)
				}
			} else {
				break
			}

		}
	}

	return keysVals
}

func (ex *Executer) isDeleteAll(lang *SkikLang) bool {
	if len(lang.Tokens) >= 3 && lang.Tokens[0].Type == DELETE && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == ALL {
		return true
	}
	return false
}
func (ex *Executer) deleteAll(lang *SkikLang) {

	ex.Db.Storage.Datad = map[string]any{}
}
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
