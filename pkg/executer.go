package pkg

import (
	"errors"
)

type Executer struct {
	Db *DB
}

func NewExecuter(db *DB) *Executer {

	return &Executer{
		Db: db,
	}
}

var (
	InvalidCode = errors.New("The code is invalid :/")
)

func (ex *Executer) Execute(code string) string {
	lang := NewSkikLang(code)
	lang.PreInit()
	err := lang.Tokenize()
	lang.Tokens = RemoveGarbageTokens(lang.Tokens)
	if err != nil {
		return jsonToStr(map[string]any{
			"ok":        false,
			"errorInfo": err.Error(),
		})
	}
	if ex.isSet(lang) {
		ex.set(lang)
		return jsonToStr(map[string]any{
			"ok": true,
		})
	} else if ex.isGet(lang) {
		f := ex.get(lang)
		if f == nil {
			return jsonToStr(map[string]any{
				"ok": false,
			})
		} else {
			return jsonToStr(map[string]any{
				"ok":    true,
				"value": f,
			})
		}

	} else if ex.isGetAll(lang) {
		return jsonToStr(map[string]any{
			"ok":     true,
			"values": ex.getAll(lang),
		})
	} else if ex.isDelete(lang) {
		err := ex.delete(lang)
		if err != nil {
			return jsonToStr(map[string]any{
				"ok":        false,
				"errorInfo": err.Error(),
			})
		}
		return jsonToStr(map[string]any{
			"ok": true,
		})
	} else if ex.isSave(lang) {
		e := ex.save(lang)
		if e != nil {
			return jsonToStr(map[string]any{
				"ok":        false,
				"errorInfo": e.Error(),
			})
		}
		return jsonToStr(map[string]any{
			"ok": true,
		})
	}
	return jsonToStr(map[string]any{
		"ok": false,
	})
}
