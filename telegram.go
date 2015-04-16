package main

import (
	"fmt"
	"github.com/sdidyk/telegram/mtproto"
)

func main() {
	var err error

	m := new(mtproto.MTProto)

	err = m.Connect("149.154.175.50:443")
	if err != nil {
		fmt.Println("Connect failed", err)
		return
	}

	err = m.Handshake()
	if err != nil {
		fmt.Println("Handshake failed:", err)
		return
	}

	go m.ReadRoutine()

	_ = m.SendPacket(mtproto.Encode_TL_ping(1), false)
	_ = m.SendPacket(mtproto.Encode_TL_help_getConfig(), true)

	fmt.Println("[press enter to quit]")
	_, _ = fmt.Scanf("abc")

}
