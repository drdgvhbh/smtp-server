package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/alecthomas/participle"
	"io/ioutil"
	"log"
	"net"
	"smtp/pkg"
	"smtp/pkg/parser"
)

const envFileName = ".env.toml"

func main() {
	tomlData, err := ioutil.ReadFile(envFileName)
	if err != nil {
		log.Fatalln( err, fmt.Sprintf("failed to read %s", envFileName))
	}
	var conf pkg.Config
	if _, err := toml.Decode(string(tomlData), &conf); err != nil {
		log.Fatalln(err, fmt.Sprintf("failed to parse %s", envFileName))
	}

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", conf.Port))
	if err != nil {
		log.Fatalln(err, "failed to open tcp socket")
	}
	defer func() {
		err = l.Close()
		if err != nil {
			log.Fatalln(err, "failed to close tcp socket")
		}
	}()

	lexer, err := parser.Lexer()
	if err != nil {
		log.Fatalln(err, "failed to create lexer")
	}
	p, err := participle.Build(&parser.SMTPGrammar{},
		participle.Lexer(lexer))
	if err != nil {
		log.Fatalln(err, "failed to create parser")
	}

	smtpServer := pkg.SMTPServer{
		DomainName: conf.DomainName,
		Parser: p,
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go smtpServer.Receive(c)
	}

}
