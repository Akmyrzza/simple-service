package cmd

import (
	"github.com/Akmyrzza/simple-service/internal/adapters/repo/sqlite"
	"github.com/Akmyrzza/simple-service/internal/adapters/server"
	"github.com/Akmyrzza/simple-service/internal/adapters/server/rest"
	"github.com/Akmyrzza/simple-service/internal/domain/core"
	"github.com/Akmyrzza/simple-service/internal/domain/usecases"
)

func Execute() {
	app := struct {
		repo       *sqlite.St
		core       *core.St
		ucs        *usecases.St
		restApiSrv *server.St
	}{}

	confLoad()

	app.repo = sqlite.New()
	app.ucs = usecases.New()
	app.core = core.New(app.repo)
	app.ucs.SetCore(app.core)
	app.restApiSrv = server.Start(conf.HttpListen, rest.GetHandler(app.ucs))

}
