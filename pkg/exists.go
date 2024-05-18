package pkg

func (ex *Executer) isExists(lang *SkikLang) bool {

	if len(lang.Tokens) >= 3 && lang.Tokens[0].Type == EXIST && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING {
		return true
	}
	return false
}
func (ex *Executer) exists(lang *SkikLang) bool {
	_, f := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)]

	return f
}
