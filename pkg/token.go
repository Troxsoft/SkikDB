package pkg

const (
	GET    = iota // get
	SET           // set
	DELETE        // delete
	WHERE
	ALL     // *
	ENDLINE // \n
	STRING  // "a text"
	NUMBER  // 0,1,etc..
	SPACE   //a espace
	JSON    // {"json":true}
	BOOLEAN // true or false
	LIST    // [2,3,4,5,6,7,8,"hola mundo",true]
	SAVE

	LIST_SYMBOL // list
	EXIST
	TYPE
	ADD_LEFT  // addl
	ADD_RIGHT // addr
)

type Token struct {
	Type  uint
	Value any
}

func RemoveGarbageTokens(garbage []Token) []Token {
	n := []Token{}
	for i, v := range garbage {
		if i != 0 {
			if garbage[i].Type == SPACE && garbage[i-1].Type == SPACE {

			} else if garbage[i].Type == ENDLINE && garbage[i-1].Type == ENDLINE {

			} else {
				n = append(n, v)
			}
		} else {
			n = append(n, v)

		}
	}
	return n
}
func NewToken(tokenType uint, value any) *Token {
	return &Token{
		Type:  tokenType,
		Value: value,
	}
}
