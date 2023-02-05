package cmd

import (
	"net/http"

	"github.com/Akmyrzza/simple-service/internal/adapters/repo/sqlite"
	"github.com/Akmyrzza/simple-service/internal/adapters/server/rest"
	"github.com/Akmyrzza/simple-service/internal/domain/core"
	"github.com/Akmyrzza/simple-service/internal/domain/usecases"
)

func Execute() {
	//sss
	app := struct {
		repo *sqlite.St
		core *core.St
		ucs  *usecases.St
	}{}

	confLoad()

	app.repo = sqlite.New()
	app.ucs = usecases.New()
	app.core = core.New(app.repo)
	app.ucs.SetCore(app.core)

	server := &http.Server{
		Addr: conf.HttpListen,
		Handler: rest.GetHandler(
			app.ucs,
		),
	}

	server.ListenAndServe()
}
