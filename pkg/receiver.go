package pkg

import (
	"net"
	"smtp/pkg/reply"
)

type SMTPServer struct {
	DomainName string
}

func (s SMTPServer) Receive(c net.Conn) error {
	reply.ServerReady(c, s.DomainName)
	return c.Close()
}
