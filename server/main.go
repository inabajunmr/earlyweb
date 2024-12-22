package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
)

func main() {
	service := ":80"
	addr, _ := net.ResolveTCPAddr("tcp4", service)
	listener, _ := net.ListenTCP("tcp", addr)
	for {
		conn, _ := listener.Accept()
		reader := bufio.NewReader(conn)
		line, _ := reader.ReadString('\n')
		regex, _ := regexp.Compile(`^GET ([^?]+)\??(.*)`)
		matches := regex.FindStringSubmatch(line)

		path := matches[1]
		var keyword string
		if len(matches) >= 3 {
			keyword = matches[2]
		}

		fmt.Println("path:" + path)
		fmt.Println("keyword:" + keyword)

		conn.Write([]byte(`
<TITLE>This is my homepage</TITLE>
<H1>Hello my homepage</H1>
<H2>Search my article</H2>
<ISINDEX>
		`))
		conn.Close()
	}
}
