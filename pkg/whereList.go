package pkg

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
)

func (executer *Executer) newExpLangListDelete(exp string, value any) *govaluate.EvaluableExpression {
	d, err := govaluate.NewEvaluableExpressionWithFunctions(exp, map[string]govaluate.ExpressionFunction{
		"startsWith": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, fmt.Errorf("startsWith function requires exactly 2 arguments")
			}
			if str, ok := arguments[0].(string); ok {
				if prefix, ok := arguments[1].(string); ok {
					return strings.HasPrefix(str, prefix), nil
				}
			}
			return nil, fmt.Errorf("startsWith function requires two string arguments")
		},
		"endsWith": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, fmt.Errorf("endsWith function requires exactly 2 arguments")
			}
			if str, ok := arguments[0].(string); ok {
				if prefix, ok := arguments[1].(string); ok {
					return strings.HasSuffix(str, prefix), nil
				}
			}
			return nil, fmt.Errorf("endsWith function requires two string arguments")
		},
	})
	if err != nil {
		panic(err)
	}
	return d
}
func (executer *Executer) whereListDelete(value any, index int, lang *SkikLang) bool {

	where := 0
	for i, v := range lang.Tokens {
		if v.Type == WHERE {
			where = i
		}
	} // get * where value <=2
	//fmt.Println(lang.Tokens[where].Value.(string))
	d := executer.newExpLangListDelete(lang.Tokens[where].Value.(string), value)
	e, err := d.Evaluate(map[string]interface{}{
		"value": value,
		"index": index,
	})
	if err != nil {
		return false
	}
	b3, err3 := e.(bool)
	//fmt.Println(b3, "  ", e)
	if err3 != true {
		return false
	}
	return b3
}

func (executer *Executer) newExpLangListGet(exp string, value any) *govaluate.EvaluableExpression {
	d, err := govaluate.NewEvaluableExpressionWithFunctions(exp, map[string]govaluate.ExpressionFunction{
		"startsWith": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, fmt.Errorf("startsWith function requires exactly 2 arguments")
			}
			if str, ok := arguments[0].(string); ok {
				if prefix, ok := arguments[1].(string); ok {
					return strings.HasPrefix(str, prefix), nil
				}
			}
			return nil, fmt.Errorf("startsWith function requires two string arguments")
		},
		"endsWith": func(arguments ...interface{}) (interface{}, error) {
			if len(arguments) != 2 {
				return nil, fmt.Errorf("endsWith function requires exactly 2 arguments")
			}
			if str, ok := arguments[0].(string); ok {
				if prefix, ok := arguments[1].(string); ok {
					return strings.HasSuffix(str, prefix), nil
				}
			}
			return nil, fmt.Errorf("endsWith function requires two string arguments")
		},
	})
	if err != nil {
		panic(err)
	}
	return d
}
func (executer *Executer) whereListGet(value any, index int, lang *SkikLang) bool {

	where := 0
	for i, v := range lang.Tokens {
		if v.Type == WHERE { // json "pepe" set "info.nombre" "juan suarez"
			where = i
		}
	} // get * where value <=2
	//fmt.Println(lang.Tokens[where].Value.(string))
	d := executer.newExpLangListGet(lang.Tokens[where].Value.(string), value)
	e, err := d.Evaluate(map[string]interface{}{
		"value": value,
		"index": index,
	})
	if err != nil {
		return false
	}
	b3, err3 := e.(bool)
	//fmt.Println(b3, "  ", e)
	if err3 != true {
		return false
	}
	return b3
}
