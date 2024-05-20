package pkg

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var (
	IS_SKIKDB_SERVER  = "#__IS_SKIKDB_SERVER__#0987654321"
	IS_VALID_PASSWORD = "#__IS_VALID_PASSWORD__#0987654321"
)

func verifyPassword(str, password string) (string, bool) {
	f := strings.SplitN(str, "#", 2)
	fmt.Printf("%v", f)
	if len(f) != 2 {
		return "", false
	}
	if f[0] == password {
		return f[1], true
	}
	return "", false
}
func (db *DB) StartServer(address string, password string) error {
	ser, err := net.Listen("tcp", address)

	db.Server = ser
	defer db.Server.Close()
	if err != nil {
		return err
	}
	ex := NewExecuter(db)
	fmt.Printf(`
Server listen on %s
	`, address)
	for {

		conn, err := db.Server.Accept()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			//conn.Close()
			continue
		}

		defer conn.Close()
		// Handle the connection in a new goroutine
		go func() {
			msg_, err := io.ReadAll(conn)
			msg := string(msg_)
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
				conn.Close()
				return
			}

			if query, pass := verifyPassword(msg, password); pass {
				fmt.Printf("MSG: %s\n", query)

				res := ex.Execute(query)

				// Enviar una respuesta al cliente
				conn.Write([]byte(res))
			} else {
				fmt.Println("INVALID PASSWORD")
				conn.Write([]byte("{\"ok\":false,\"errorInfo\":\"invalid password\"}"))
			}

		}()
	}
	return nil
}
