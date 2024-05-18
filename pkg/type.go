package pkg

func (ex *Executer) isType(lang *SkikLang) bool {

	if len(lang.Tokens) >= 3 && lang.Tokens[0].Type == TYPE && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING {
		return true
	}
	return false
}

func (ex *Executer) typee(lang *SkikLang) any {
	v, err := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if err == true {
		switch v.(type) {
		case float64:
			{
				return "Number"
			}
		case map[string]any:
			{
				return "JSON"
			}
		case []any:
			{
				return "List"
			}
		case string:
			{
				return "String"
			}
		}

	}
	return "Undefined"
}
