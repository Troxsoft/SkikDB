package pkg

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
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

	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	ex := NewExecuter(db)
	fmt.Printf(`
Server listen on %s
	`, address)
	for {
		conn, err := ser.Accept()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			//conn.Close()
			conn.Close()
			continue
		}
		goroutineLimit := make(chan struct{}, 10000) // Limitar a 100 goroutines simultáneas
		// Handle the connection in a new goroutine
		go func(conn3 net.Conn) {
			wg.Add(1)
			goroutineLimit <- struct{}{} // Bloquear si hay 100 goroutines activas
			msg_, err := bufio.NewReader(conn3).ReadString('\n')
			defer wg.Done()
			defer func() { <-goroutineLimit }() // Liberar una goroutine del límite
			msg := string(msg_)
			if err != nil {
				fmt.Println("ERROR: ", err.Error(), msg)

				conn3.Close()
				return
			}

			if query, pass := verifyPassword(msg, password); pass {
				fmt.Printf("MSG: %s\n", query)
				mu.Lock()
				res := ex.Execute(query)
				mu.Unlock()
				// Enviar una respuesta al cliente
				conn3.Write([]byte(res))
			} else {
				fmt.Println("INVALID PASSWORD")
				conn3.Write([]byte("{\"ok\":false,\"errorInfo\":\"invalid password\"}"))
			}

			conn3.Close()
		}(conn)
	}
	db.Server.Close()
	return nil
}
