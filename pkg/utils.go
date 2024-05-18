package pkg

import "encoding/json"

func isIntNumber(num byte) bool {
	if num == '0' || num == '1' || num == '2' || num == '3' || num == '4' || num == '5' || num == '6' || num == '7' || num == '8' || num == '9' {
		return true
	}
	return false
}
func jsonIdentToStr(js any) string {
	d, err := json.MarshalIndent(js,"","	")
	if err != nil {
		panic(err)
	}
	return string(d)
}
func jsonToStr(js any) string {
	d, err := json.Marshal(js)
	if err != nil {
		panic(err)
	}
	return string(d)
}
