package main

import (
	"encoding/json"
	"fmt"
	"go_rpc_dirac"
	"go_rpc_dirac/codec"
	"log"
	"net"
	"time"
)


func startServer(addr_ chan string) {
	con, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("Network Error:", err)
	}
	log.Println("Start rpc-server on:" ,con.Addr())
	addr_ <- con.Addr().String()
	go_rpc_dirac.Accept(con)
}

func main() {
	log.SetFlags(0)
	addr := make(chan string)
	go startServer(addr)


	conn, _ := net.Dial("tcp", <-addr)
	defer func() { _ = conn.Close() }()

	time.Sleep(time.Second)

	_ = json.NewEncoder(conn).Encode(go_rpc_dirac.DefaultOption)
	cc := codec.NewGobCodec(conn)

	for i := 0; i < 5; i++ {
		h := &codec.Header{
			ServiceMeth: "Foo.Sum",
			Seq:           uint64(i),
		}
		_ = cc.Write(h, fmt.Sprintf("go_rpc_dirac req %d", h.Seq))
		_ = cc.ReadHeader(h)
		var reply string
		_ = cc.ReadBody(&reply)
		log.Println("reply:", reply)
	}
}