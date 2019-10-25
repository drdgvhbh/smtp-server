package reply

import (
	"fmt"
	"net"
)

const terminator = "\r\n"

func ServerReady(c net.Conn, domainName string) error {
	_, err := c.Write([]byte(fmt.Sprintf("%d %s%s", Ready, domainName, terminator)))
	return err
}
