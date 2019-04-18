package main

import (
	_ "net/http/pprof"

	"github.com/lukasjarosch/educonn-monorepo/services/account"
	"github.com/lukasjarosch/educonn-monorepo/services/account/account/cmd"
	"github.com/lukasjarosch/educonn-monorepo/services/account/server"
	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.AccountServer{}

	lile.Name("account")
	lile.Server(func(g *grpc.Server) {
		account.RegisterAccountServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
