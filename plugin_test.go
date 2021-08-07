package plugin

import (
	"github.com/spiral/roadrunner/v2/plugins/logger"
	"net"
	netRpc "net/rpc"
	"os"
	"os/signal"
	"sync"
	"syscall"

	endure "github.com/spiral/endure/pkg/container"
	"github.com/spiral/roadrunner/v2/plugins/config"
	rpcPlugin "github.com/spiral/roadrunner/v2/plugins/rpc"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	goridgeRpc "github.com/spiral/goridge/v3/pkg/rpc"
)

// testConfig provides a yaml representation of what the config is for the test
func testConfig() string {
	return `
rpc:
  listen: tcp://127.0.0.1:6001
plugin:
  enable: true
`
}

// TestPluginInit creates a service container with the RPC, Logger and the Plugin to run test
func TestPluginInit(t *testing.T) {
	cont, err := endure.NewContainer(nil, endure.SetLogLevel(endure.ErrorLevel))
	if err != nil {
		t.Fatal(err)
	}

	cfg := &config.Viper{}
	cfg.Type = "yaml"
	cfg.ReadInCfg = []byte(testConfig())

	// configure the container with the right plugins
	err = cont.RegisterAll(
		cfg,
		&rpcPlugin.Plugin{},
		&logger.ZapLogger{},
		&Plugin{},
	)
	assert.NoError(t, err)

	err = cont.Init()
	if err != nil {
		t.Fatal(err)
	}

	ch, err := cont.Serve()
	assert.NoError(t, err)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	stopCh := make(chan struct{}, 1)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case e := <-ch:
				assert.Fail(t, "error", e.Error.Error())
				err = cont.Stop()
				if err != nil {
					assert.FailNow(t, "error", err.Error())
				}
			case <-sig:
				err = cont.Stop()
				if err != nil {
					assert.FailNow(t, "error", err.Error())
				}
				return
			case <-stopCh:
				// timeout
				err = cont.Stop()
				if err != nil {
					assert.FailNow(t, "error", err.Error())
				}
				return
			}
		}
	}()

	time.Sleep(time.Second)

	// execute tests against the rpc server/plugin
	t.Run("messageRpcTest", messageRPCTest)
	stopCh <- struct{}{}
	wg.Wait()
}

// messageRPCTest tests if the message rpc endpoint works as expected
func messageRPCTest(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	assert.NoError(t, err)
	client := netRpc.NewClientWithCodec(goridgeRpc.NewClientCodec(conn))

	var returned Payload
	input := Payload{Message: "test result"}
	err = client.Call("plugin.Message", input, &returned)
	assert.NoError(t, err)
	assert.Equal(t, input, returned)
}
