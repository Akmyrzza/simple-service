package cmd

import (
	"github.com/Akmyrzza/simple-service/internal/adapters/logger/zap"
	"github.com/Akmyrzza/simple-service/internal/adapters/repo/sqlite"
	"github.com/Akmyrzza/simple-service/internal/adapters/server"
	"github.com/Akmyrzza/simple-service/internal/adapters/server/rest"
	"github.com/Akmyrzza/simple-service/internal/domain/core"
	"github.com/Akmyrzza/simple-service/internal/domain/usecases"
)

func Execute() {
	app := struct {
		lg   *zap.St
		repo *sqlite.St
		core *core.St
		ucs  *usecases.St
		srv  *server.St
	}{}

	confLoad()

	app.lg = zap.New(conf.LogLevel, conf.Debug)
	app.repo = sqlite.New()
	app.core = core.New(app.repo)
	app.ucs = usecases.New()
	app.ucs.SetCore(app.core)
	app.srv = server.Start(conf.HttpListen, rest.GetHandler(app.ucs))

}
