package server

import (
	"github.com/codegangsta/negroni"

	"github.com/dnl-jst/goard-daemon/v1/config"
	"github.com/dnl-jst/goard-daemon/v1/router"
)

func Listen() {
	n := negroni.New()
	n.UseHandler(router.New())
	n.Run(config.Std.Http.Listen)
}
