package pkg

func (ex *Executer) isGet(lang *SkikLang) bool {

	if len(lang.Tokens) >= 3 && lang.Tokens[0].Type == GET && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING {
		return true
	}
	return false
}

func (ex *Executer) get(lang *SkikLang) any {
	v, err := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if err == true {
		return v
	}
	return nil
}
func (ex *Executer) isGetAll(lang *SkikLang) bool {
	if len(lang.Tokens) >= 3 && lang.Tokens[0].Type == GET && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == ALL {
		return true
	}
	return false
}
func (ex *Executer) getAll(lang *SkikLang) map[string]any {

	return ex.Db.Storage.Datad
}
func (ex *Executer) isGetWhere(lang *SkikLang) bool {
	if len(lang.Tokens) >= 5 && lang.Tokens[0].Type == GET && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == ALL && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == WHERE {
		return true
	}
	if len(lang.Tokens) >= 5 && lang.Tokens[0].Type == GET && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == NUMBER && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == WHERE {
		return true
	}
	return false
}

/*
	example:

get * where(key startsWith "+18") retorna las claves que empiezan con +18
*/
func (ex *Executer) getWhere(lang *SkikLang) map[string]any {
	keysVals := make(map[string]any)
	if lang.Tokens[2].Type == ALL {
		for k, v := range ex.Db.Storage.Datad {
			b := ex.whereGet(k, v, lang)
			if b {
				keysVals[k] = v
			}
		}
	} else {
		lolol := 0.0
		for k, v := range ex.Db.Storage.Datad {
			if lolol < lang.Tokens[2].Value.(float64) {
				b := ex.whereGet(k, v, lang)
				if b {
					keysVals[k] = v
					lolol++
				}
			} else {
				break
			}

		}
	}

	return keysVals
}
