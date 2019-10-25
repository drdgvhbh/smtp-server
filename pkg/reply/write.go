package reply

import (
	"fmt"
	"net"
)

func write(c net.Conn, s string) (int, error) {
	n, err := c.Write([]byte(fmt.Sprintf("%s\r\n", s)))
	return n, err
}
