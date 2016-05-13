package main

import (
	"fmt"
	"github.com/VoltDB/voltdb-client-go/voltdbclient"
	"os"
)

func main() {
	conn, err := voltdbclient.NewConnection("", "", "localhost:21212")
	if err != nil {
		fmt.Println("failed to connect to server")
		os.Exit(-1);
	}
	fmt.Println(conn)
}

