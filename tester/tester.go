package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"proxy/diag"
	"proxy/policy/error"
	"proxy/socks4"
)

func main() {
	testSlice := make([]byte, 0, 8)
	testSlice = append(testSlice, 1, 2, 3)
	fmt.Printf("%v\n", testSlice)
	fmt.Println("Starting...")
	tracer := diag.NewTracer(os.Stdout)
	tracer.TraceLine("Starting")
	if listener, err := net.Listen("tcp", ":1080"); err == nil {
		if connection, err := listener.Accept(); err == nil {
			request := make([]byte, 1024)
			if n, _ := connection.Read(request); n > 0 {
				fmt.Printf("%d", n)
				fmt.Printf("Var : %X Command %X\n", request[0], request[1])
				fmt.Printf("Port: %X %X %d %d\n", request[2:3], request[3:4], binary.BigEndian.Uint16(request[2:4]), request[4:5])
				dstport := request[2:4]
				dstip := request[4:8]
				userid := request[8 : len(request)-1]
				dstports := string(dstport)
				dstips := string(dstip)
				userstring := string(userid)
				tracer.TraceLine("p" + dstports)
				tracer.TraceLine(dstips)
				tracer.TraceLine(userstring)
				response, _ := socks4.NewResponse(socks4.Granted)
				connection.Write(response.Bytes())
				fmt.Println(response.String())
				bindbytes := make([]byte, 64000, 64000)
				_, _ = connection.Read(bindbytes)
				//bindReq, _ := socks4.NewRequest(bindbytes)
				req1 := string(bindbytes[:bytes.IndexByte(bindbytes, 0)])
				fmt.Println(req1)
				fmt.Printf("%v\n", bindbytes)
				//fmt.Println(bindReq.String())

				if socksReq, err := socks4.NewRequest(request); err != nil {
					error.HandleAndExit(err)
				} else {
					fmt.Println(socksReq.String())
				}
			}
		} else {
			error.HandleAndExit(err)
		}
	} else {
		error.HandleAndExit(err)
	}
}
