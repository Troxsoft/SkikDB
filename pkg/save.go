package pkg

func (ex *Executer) isSave(lang *SkikLang) bool {

	if len(lang.Tokens) >= 1 && lang.Tokens[0].Type == SAVE {
		return true
	}
	return false
}

func (ex *Executer) save(lang *SkikLang) error {
	return ex.Db.Save()
}
