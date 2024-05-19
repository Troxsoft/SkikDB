package pkg

func (ex *Executer) isListDeleteAll(lang *SkikLang) bool {
	if len(lang.Tokens) >= 7 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == DELETE && lang.Tokens[5].Type == SPACE && lang.Tokens[6].Type == ALL {
		return true
	}
	return false
}
func (skl *Executer) listDeleteAll(lang *SkikLang) error {
	listName, exists := skl.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if !exists {
		return KEY_VALUE_NOT_FOUND
	}
	switch listName.(type) {
	case []any:
		{
			skl.Db.Storage.Datad[lang.Tokens[2].Value.(string)] = []any{}
		}
	default:
		{
			return TYPE_INVALID
		}
	}

	return nil
}
func (ex *Executer) isListDeleteWhere(lang *SkikLang) bool {
	if len(lang.Tokens) >= 9 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == DELETE && lang.Tokens[5].Type == SPACE && lang.Tokens[6].Type == ALL && lang.Tokens[7].Type == SPACE && lang.Tokens[8].Type == WHERE {
		return true
	}
	if len(lang.Tokens) >= 9 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == DELETE && lang.Tokens[5].Type == SPACE && lang.Tokens[6].Type == NUMBER && lang.Tokens[7].Type == SPACE && lang.Tokens[8].Type == WHERE {
		return true
	}
	return false
}
func (ex *Executer) isListGetWhere(lang *SkikLang) bool {
	if len(lang.Tokens) >= 9 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == GET && lang.Tokens[5].Type == SPACE && lang.Tokens[6].Type == ALL && lang.Tokens[7].Type == SPACE && lang.Tokens[8].Type == WHERE {
		return true
	}
	if len(lang.Tokens) >= 9 && lang.Tokens[0].Type == LIST_SYMBOL && lang.Tokens[1].Type == SPACE && lang.Tokens[2].Type == STRING && lang.Tokens[3].Type == SPACE && lang.Tokens[4].Type == GET && lang.Tokens[5].Type == SPACE && lang.Tokens[6].Type == NUMBER && lang.Tokens[7].Type == SPACE && lang.Tokens[8].Type == WHERE {
		return true
	}
	return false
}
func (ex *Executer) listGetWhere(lang *SkikLang) (*[]any, error) {

	listName, exists := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if !exists {
		return nil, KEY_VALUE_NOT_FOUND
	}
	newList := []any{}
	switch listName.(type) {
	case []any:
		{

			k30, _ := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)].([]any)
			if lang.Tokens[6].Type == ALL {
				for i, v := range k30 {
					lolo90 := ex.whereListGet(v, i, lang)
					if lolo90 {
						newList = append(newList, v)
					}
				}
			} else {
				limit := 0.0
				for i, v := range k30 {
					if limit < lang.Tokens[6].Value.(float64) {
						lolo90 := ex.whereListGet(v, i, lang)
						if lolo90 {

							newList = append(newList, v)
							limit++
						}
					} else {
						break
					}

				}
			}

		}
	default:
		{
			return nil, TYPE_INVALID
		}
	}

	return &newList, nil
}
func (ex *Executer) listDeleteWhere(lang *SkikLang) (*[]any, error) {
	eliminates := []any{}
	listName, exists := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)]
	if !exists {
		return nil, KEY_VALUE_NOT_FOUND
	}
	switch listName.(type) {
	case []any:
		{
			newList := []any{}
			k30, _ := ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)].([]any)
			if lang.Tokens[6].Type == ALL {
				for i, v := range k30 {
					lolo90 := ex.whereListDelete(v, i, lang)
					if lolo90 {
						eliminates = append(eliminates, v)
					} else {
						newList = append(newList, v)
					}
				}
			} else {
				limit := 0.0
				for i, v := range k30 {
					if limit < lang.Tokens[6].Value.(float64) {
						lolo90 := ex.whereListDelete(v, i, lang)
						if lolo90 {
							eliminates = append(eliminates, v)
							limit++
						} else {
							newList = append(newList, v)
						}
					} else {
						break
					}

				}
			}

			ex.Db.Storage.Datad[lang.Tokens[2].Value.(string)] = newList
		}
	default:
		{
			return nil, TYPE_INVALID
		}
	}

	return &eliminates, nil
}
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
