package reply

import (
	"fmt"
	"net"
	"smtp/pkg/reply/codes"
)

func ServerReady(c net.Conn, domainName string) error {
	_, err := write(c, fmt.Sprintf("%d %s", codes.Ready, domainName))
	return err
}

func Ok(c net.Conn) error {
	_, err := write(c, fmt.Sprintf("%d", codes.Ok))
	return err
}
