package p2p

import (
	"fmt"
	ma "github.com/multiformats/go-multiaddr"
	"testing"
)

func TestLocalForward(t *testing.T) {
	str := "/ip4/127.0.0.1/tcp/8080"
	listenAddress, err := ma.NewMultiaddr(str)
	if err != nil {
		t.Error(err)
	}

	obj := listenAddress.String()
	fmt.Println(obj)
}
