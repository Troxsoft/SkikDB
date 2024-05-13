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
