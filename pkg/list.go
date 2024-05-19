package pkg

func (skl *Executer) isListAddl(lang *SkikLang) bool {
	if len(lang.Tokens) >= 7 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == ADD_LEFT {
		return true
	}
	return false
}
func (skl *Executer) listAddr(lang *SkikLang) error {
	listName, exists := skl.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if !exists {
		return KEY_VALUE_NOT_FOUND
	}
	switch listName.(type) {
	case []any:
		{
			list := listName.([]any)
			list = append(list, lang.Tokens[6].Value)
			skl.Db.Storage.Datad[lang.Tokens[2].Value.(string)] = list
		}
	default:
		{
			return TYPE_INVALID
		}
	}

	return nil
}

// list "pepe" add "pep"

func (skl *Executer) isListAddr(lang *SkikLang) bool {
	if len(lang.Tokens) >= 7 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == ADD_RIGHT {
		return true
	}
	return false
}
func (skl *Executer) listAddl(lang *SkikLang) error {
	listName, exists := skl.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if !exists {
		return KEY_VALUE_NOT_FOUND
	}
	switch listName.(type) {
	case []any:
		{
			list := listName.([]any)
			h := []any{lang.Tokens[6].Value}
			h = append(h, list...)
			skl.Db.Storage.Datad[lang.Tokens[2].Value.(string)] = h
		}
	default:
		{
			return TYPE_INVALID
		}
	}

	return nil
}

// list "pepe" add "pep"
