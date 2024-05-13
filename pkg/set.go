package pkg

func (ex *Executer) isSet(lang *SkikLang) bool {

	if len(lang.Tokens) >= 5 && lang.Tokens[0].Type == SET && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE {
		return true
	}
	return false
}

func (ex *Executer) set(lang *SkikLang) {
	ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)] = lang.Tokens[4].Value
}
