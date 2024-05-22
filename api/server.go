package api

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

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

func readAll(conn net.Conn) ([]byte, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, conn)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (r *ServerSkikDB) Query(code string) any {
	h, err := net.Dial("tcp", r.address)

	if err != nil {
		return err
	}
	defer h.Close()
	_, err = h.Write([]byte(fmt.Sprintf("%s#%s\n", r.password, code)))
	if err != nil {
		return err
	}

	data, err := bufio.NewReader(h).ReadString('\n')

	//msg := string(data)
	var res map[string]any
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return err
	}
	return res
}
func NewServerSkikDB(address, password string) (*ServerSkikDB, error) {

	return &ServerSkikDB{
		address:  address,
		password: password,
	}, nil
}
