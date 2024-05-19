package pkg

import (
	"errors"
	//"fmt"
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
	//fmt.Printf("%s\n", jsonIdentToStr(lang.Tokens))
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
	} else if ex.isDeleteWhere(lang) {
		f := ex.deleteWhere(lang)
		return jsonToStr(map[string]any{
			"ok":         true,
			"eliminates": f,
		})
	} else if ex.isGetWhere(lang) {
		f := ex.getWhere(lang)
		return jsonToStr(map[string]any{
			"ok":     true,
			"values": f,
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
		}) // delete * where value != "aña"
	} else if ex.isDeleteAll(lang) {
		ex.deleteAll(lang)
		return jsonToStr(map[string]any{
			"ok": true,
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
	} else if ex.isExists(lang) {
		return jsonToStr(map[string]any{
			"ok":     true,
			"exists": ex.exists(lang),
		})
	} else if ex.isType(lang) {
		f := ex.typee(lang)
		if f == "Undefined" {
			return jsonToStr(map[string]any{
				"ok":        false,
				"errorInfo": "key/value not exists",
			})
		} else {
			return jsonToStr(map[string]any{
				"ok":   true,
				"type": f,
			})
		}
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
	} else if ex.isListAddl(lang) {
		f := ex.listAddl(lang)
		if f != nil {
			return jsonToStr(map[string]any{
				"ok":        false,
				"errorInfo": f.Error(),
			})
		}
		return jsonToStr(map[string]any{
			"ok": true,
		})
	} else if ex.isListAddr(lang) {
		f := ex.listAddr(lang)
		if f != nil {
			return jsonToStr(map[string]any{
				"ok":        false,
				"errorInfo": f.Error(),
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
