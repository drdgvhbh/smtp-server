package pkg

import (
	"bufio"
	"github.com/alecthomas/participle"
	"github.com/pkg/errors"
	"log"
	"net"
	"smtp/pkg/parser"
	"smtp/pkg/reply"
)

type SMTPServer struct {
	DomainName string
	Parser *participle.Parser
}

type state struct {
	clientFQDN *string
}

func (s SMTPServer) Receive(c net.Conn) {
	err := reply.ServerReady(c, s.DomainName); if err != nil {
		log.Println(errors.Wrapf(err, "failed to emit server ready message"))
	}
	scanner := bufio.NewScanner(c)

	state := state{}
	for {
		for scanner.Scan() {
			line := scanner.Text()
			g := parser.SMTPGrammar{}
			err := s.Parser.ParseString(line, &g); if err != nil {
				_ = reply.SyntaxError(c)
				break
			}
			if g.Hello != nil {
				state.clientFQDN = g.Hello.FQDN
				_ = reply.Ok(c)
			}
		}
	}
}
