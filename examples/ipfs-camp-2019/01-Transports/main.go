package main

import (
	"github.com/webtransport/libp2p-go"
)

func main() {
	// TODO: add some libp2p.Transport options to this chain!
	transports := libp2p.ChainOptions()

	host, err := libp2p.New(transports)
	if err != nil {
		panic(err)
	}

	host.Close()
}
