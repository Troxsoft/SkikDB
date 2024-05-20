package api

import (
	"encoding/json"
	"errors"

	"io"
	"net"
)

type ServerSkikDB struct {
	address  string
	password string
}

var (
	INVALID_PASSWORD = errors.New("Invalid password")
	INVALID_SERVER   = errors.New("Server isn`t SkikDB server ")
)

type res1 struct {
	Value bool `json:"value"`
}

func (r *ServerSkikDB) Query(code string) any {
	h, err := net.Dial("tcp", r.address)

	if err != nil {
		return err
	}
	defer h.Close()
	_, err = h.Write([]byte(r.password + "#" + code))
	if err != nil {
		return err
	}

	data, err := io.ReadAll(h)
	if err != nil {
		return err
	}
	//msg := string(data)
	var res map[string]any
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &res)
	if err != nil {
		panic(err)
	}
	return res
}
func NewServerSkikDB(address, password string) (*ServerSkikDB, error) {

	return &ServerSkikDB{
		address:  address,
		password: password,
	}, nil
}
