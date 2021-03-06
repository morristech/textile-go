package gateway_test

import (
	"fmt"
	. "github.com/textileio/textile-go/gateway"
	"github.com/textileio/textile-go/repo/config"
	"testing"
)

func TestNewGateway(t *testing.T) {
	Host = &Gateway{}
	Host.Start(fmt.Sprintf("127.0.0.1:%d", config.GetRandomPort()))
}

func TestGateway_Addr(t *testing.T) {
	if len(Host.Addr()) == 0 {
		t.Error("get gateway address failed")
	}
}

func TestGateway_Stop(t *testing.T) {
	err := Host.Stop()
	if err != nil {
		t.Errorf("stop gateway failed: %s", err)
	}
}
