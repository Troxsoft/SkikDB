package pkg

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
)

// func replaceWordsIfNotString(text, oldWord, newWord string) string {
// 	// Patrón para encontrar palabras entre comillas
// 	quotesPattern := regexp.MustCompile(`"[^"]*"`)

// 	// Encontrar todas las palabras entre comillas y sus posiciones
// 	positions := quotesPattern.FindAllStringIndex(text, -1)

// 	// Lista para almacenar partes del texto
// 	parts := []string{}
// 	start := 0

// 	for _, position := range positions {
// 		// Agregar parte antes de la palabra entre comillas
// 		parts = append(parts, text[start:position[0]])
// 		// Agregar la palabra entre comillas sin cambios
// 		parts = append(parts, text[position[0]:position[1]])
// 		start = position[1]
// 	}

// 	// Agregar la parte final del texto
// 	parts = append(parts, text[start:])

// 	// Unir las partes, ignorando las que están entre comillas
// 	textWithoutQuotes := strings.Join(parts, "")

// 	// Reemplazar palabras en la parte sin comillas
// 	wordPattern := regexp.MustCompile(fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(oldWord)))
// 	replacedText := wordPattern.ReplaceAllString(textWithoutQuotes, newWord)

// 	// Recombinar las partes del texto
// 	for _, position := range positions {
// 		replacedText = replacedText[:position[0]] + text[position[0]:position[1]] + replacedText[position[0]:]
// 	}

// 	// Manejar palabras con \ antes de comillas
// 	escapedPattern := regexp.MustCompile(fmt.Sprintf(`\\(["']%s["'])`, regexp.QuoteMeta(oldWord)))
// 	finalText := escapedPattern.ReplaceAllString(replacedText, fmt.Sprintf(`\\%s`, newWord))

//		return finalText
//	}
func (executer *Executer) newExpLangGET(exp string, key string, value any) *govaluate.EvaluableExpression {
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
func (executer *Executer) whereGet(keyName string, value any, lang *SkikLang) bool {

	where := 0
	for i, v := range lang.Tokens {
		if v.Type == WHERE {
			where = i
		}
	} // get * where value <=2
	fmt.Println(lang.Tokens[where].Value.(string))
	d := executer.newExpLangGET(lang.Tokens[where].Value.(string), keyName, value)
	e, err := d.Evaluate(map[string]interface{}{
		"key":   keyName,
		"value": value,
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
