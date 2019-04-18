package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"../../pers/educonn-monorepo/services/account"
	"github.com/lileio/lile"
)

var s = AccountServer{}
var cli account.AccountClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		account.RegisterAccountServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = account.NewAccountClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
