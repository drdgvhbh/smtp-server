package reply

import (
	"fmt"
	"net"
	"smtp/pkg/reply/codes"
)

func SyntaxError(c net.Conn) error {
	_, err := write(c, fmt.Sprintf("%d", codes.SyntaxError))
	return err
}
