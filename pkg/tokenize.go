package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	SYMBOL_INVALID = errors.New("Symbol/Word is invalid")
)

func isValidSize(s string, index, l int) bool {
	if index+l < len(s) {
		return true
	}
	return false
}
func msgError(lineEror, msg string, indexError int, line, index int) string {
	strErr := ""
	i := 0

	for i < len(lineEror)+index {
		if i == indexError {
			strErr += "^"
		} else {
			strErr += " "
		}
		i++
	}
	return fmt.Sprintf(
		`%v	at(%v:%v)
%v	%v`, lineEror, line+1, index+1, strErr, msg)
}
func msgErrorErr(lineEror, msg string, indexError int, line, index int) error {
	return errors.New(msgError(lineEror, msg, indexError, line, index))
}
func (skl *SkikLang) Tokenize() error {
	//skl.SourceCode = strings.ReplaceAll(skl.SourceCode,"\r\n","\n")
	line := 0
	index := 0
	lines := strings.Split(skl.SourceCode, "\n")
	for index < len(skl.SourceCode) {
		//fmt.Println(index)
		if skl.SourceCode[index] == '\n' {
			skl.Tokens = append(skl.Tokens, *NewToken(ENDLINE, "\n"))
			line++
		} else if skl.SourceCode[index] == '*' {
			skl.Tokens = append(skl.Tokens, *NewToken(ALL, "*"))

		} else if skl.SourceCode[index] == '{' {
			jsonStart := index
			jsonEnd := index
			strs := 0
			// {
			breck1 := 1
			// }
			breck2 := 0
			isCompleted := true
			for isCompleted {
				if skl.SourceCode[jsonEnd] == '"' {
					strs++
				}
				if skl.SourceCode[jsonEnd] == '{' && strs != 0 && strs%2 == 0 {
					breck1++
				}
				if skl.SourceCode[jsonEnd] == '}' && strs != 0 && strs%2 == 0 {
					if breck2+1 == breck1 {
						isCompleted = false
						breck2++
						break
					} else {
						breck2++
					}
				}
				if jsonEnd+1 < len(skl.SourceCode) {
					jsonEnd++
				} else {
					break
				}
			}

			if isCompleted {
				return msgErrorErr(lines[line], "Expectative JSON end( \"}\") but found EOF", jsonEnd+1, line, jsonEnd+1)
			}
			var p any
			err := json.Unmarshal([]byte(skl.SourceCode[jsonStart:jsonEnd+1]), &p)
			if err != nil {
				return err
			}
			index += jsonEnd - jsonStart
			skl.Tokens = append(skl.Tokens, *NewToken(JSON, p))
		} else if skl.SourceCode[index] == '[' {
			jsonStart := index
			jsonEnd := index
			strs := 0
			// {
			breck1 := 1
			// }
			breck2 := 0
			isCompleted := true
			for isCompleted {
				if skl.SourceCode[jsonEnd] == '"' {
					strs++
				}
				if skl.SourceCode[jsonEnd] == '[' && strs != 0 && strs%2 == 0 {
					breck1++
				}
				if skl.SourceCode[jsonEnd] == ']' && strs != 0 && strs%2 == 0 {
					if breck2+1 == breck1 {
						isCompleted = false
						breck2++
						break
					} else {
						breck2++
					}
				}
				if jsonEnd+1 < len(skl.SourceCode) {
					jsonEnd++
				} else {
					break
				}
			}

			if isCompleted {
				return msgErrorErr(lines[line], "Expectative LIST end( \"}\") but found EOF", jsonEnd+1, line, jsonEnd+1)
			}
			var p []any
			err := json.Unmarshal([]byte(skl.SourceCode[jsonStart:jsonEnd+1]), &p)
			if err != nil {
				return err
			}
			index += jsonEnd - jsonStart
			skl.Tokens = append(skl.Tokens, *NewToken(LIST, p))
		} else if skl.SourceCode[index] == ' ' {
			skl.Tokens = append(skl.Tokens, *NewToken(SPACE, " "))
		} else if isValidSize(skl.SourceCode, index, 2) && skl.SourceCode[index] == 'g' && skl.SourceCode[index+1] == 'e' && skl.SourceCode[index+2] == 't' {
			skl.Tokens = append(skl.Tokens, *NewToken(GET, "get"))
			index += 2
		} else if isValidSize(skl.SourceCode, index, 2) && skl.SourceCode[index] == 's' && skl.SourceCode[index+1] == 'e' && skl.SourceCode[index+2] == 't' {
			skl.Tokens = append(skl.Tokens, *NewToken(SET, "set"))
			index += 2
		} else if isValidSize(skl.SourceCode, index, 3) && skl.SourceCode[index] == 't' && skl.SourceCode[index+1] == 'r' && skl.SourceCode[index+2] == 'u' && skl.SourceCode[index+3] == 'e' {
			skl.Tokens = append(skl.Tokens, *NewToken(BOOLEAN, true))
			index += 3
		} else if isValidSize(skl.SourceCode, index, 3) && skl.SourceCode[index] == 's' && skl.SourceCode[index+1] == 'a' && skl.SourceCode[index+2] == 'v' && skl.SourceCode[index+3] == 'e' {
			skl.Tokens = append(skl.Tokens, *NewToken(SAVE, "save"))
			index += 3
		} else if isValidSize(skl.SourceCode, index, 5) && skl.SourceCode[index] == 'd' && skl.SourceCode[index+1] == 'e' && skl.SourceCode[index+2] == 'l' && skl.SourceCode[index+3] == 'e' && skl.SourceCode[index+4] == 't' && skl.SourceCode[index+5] == 'e' {
			skl.Tokens = append(skl.Tokens, *NewToken(DELETE, "delete"))
			index += 5
		} else if isValidSize(skl.SourceCode, index, 4) && skl.SourceCode[index] == 'f' && skl.SourceCode[index+1] == 'a' && skl.SourceCode[index+2] == 'l' && skl.SourceCode[index+3] == 's' && skl.SourceCode[index+4] == 'e' {
			skl.Tokens = append(skl.Tokens, *NewToken(BOOLEAN, false))
			index += 4
		} else if skl.SourceCode[index] == '"' {
			strStart := index
			strEnd := index + 1
			isCompleted := false

			//fmt.Println(strEnd)
			for strEnd < len(skl.SourceCode) && skl.SourceCode[strEnd] != '"' {
				//fmt.Println(string(skl.SourceCode[strEnd-1]))

				if strEnd+1 < len(skl.SourceCode) {

					strEnd++

					if skl.SourceCode[strEnd-1] == '\\' && skl.SourceCode[strEnd] == '"' {
						strEnd++
						if skl.SourceCode[strEnd] == '"' {
							isCompleted = true
						}
					} else {
						if skl.SourceCode[strEnd] == '"' {
							isCompleted = true
						}
					}
				} else {

					break
				}
			}

			if !isCompleted {
				return msgErrorErr(lines[line], "Expectative close string but found: EOF", index+(strEnd-strStart), line, index+(strEnd-strStart))
			}
			g := strings.ReplaceAll(skl.SourceCode[strStart+1:strEnd], `\"`, `"`)
			//g = skl.SourceCode[strStart+1 : strEnd]
			skl.Tokens = append(skl.Tokens, *NewToken(STRING, g))
			index += (strEnd - strStart)
		} else if isIntNumber(skl.SourceCode[index]) {
			numStart := index
			numEnd := index
			isFloat := false
			for numEnd < len(skl.SourceCode) && isIntNumber(skl.SourceCode[numEnd]) {

				if numEnd+1 < len(skl.SourceCode) {
					numEnd++
					if skl.SourceCode[numEnd] == '.' && isFloat == false {
						isFloat = true
						numEnd++
					} else if skl.SourceCode[numEnd] == '.' && isFloat == true {
						return msgErrorErr(lines[line], "Expectative number but found: '.'", numEnd, line, numEnd)

					}

				} else {
					break
				}
			}
			fl, err := strconv.ParseFloat(skl.SourceCode[numStart:numEnd], 64)
			if err != nil {
				return err

			}
			skl.Tokens = append(skl.Tokens, *NewToken(NUMBER, fl))
			index += (numEnd - numStart) - (1)
		} else {
			return msgErrorErr(lines[line], "Invalid char", index, line, index)
		}
		index++
	}
	return nil
}
