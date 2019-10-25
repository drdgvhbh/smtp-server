package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"net"
	"smtp/pkg"
)

const envFileName = ".env.toml"

func main() {
	tomlData, err := ioutil.ReadFile(".env.toml")
	if err != nil {
		log.Fatalln(err, "failed to read %s", ".env.toml")
	}
	var conf pkg.Config
	if _, err := toml.Decode(string(tomlData), &conf); err != nil {
		log.Fatalln(err, "failed to parse %s", envFileName)
		// handle error
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

	smtpServer := pkg.SMTPServer{
		DomainName: conf.DomainName,
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go smtpServer.Receive(c)
	}

}
