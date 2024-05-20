package api

import (
	"encoding/json"

	"github.com/Troxsoft/SkikDB/pkg"
)

type SkikDB struct {
	DB       *pkg.DB
	Executer *pkg.Executer
}

func (sk *SkikDB) Query(code string) any {
	j := sk.Executer.Execute(code)
	var g map[string]any
	json.Unmarshal([]byte(j), &g)
	return g
}
func NewSkikDB(filepath string) (*SkikDB, error) {
	j, err := pkg.NewDB(filepath)

	if err != nil {
		return nil, err
	}
	return &SkikDB{
		DB:       j,
		Executer: pkg.NewExecuter(j),
	}, nil
}
