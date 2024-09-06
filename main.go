package beeboo

import (
	"fmt"
	"net"
)

const (
	ENDPOINT = "::8080"
)

func main() {
	conn, err := net.Dial("tcp", ENDPOINT)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Write()
}
