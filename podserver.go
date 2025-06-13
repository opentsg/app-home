// Copyright ©2022-2025 Mr MXF   info@mrmxf.com
// BSD-3-Clause License   https://opensource.org/license/bsd-3-clause/

package main

// package main is a simple static web server that serves a hugo site.
//

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/mrmxf/clog/gommi"
	"github.com/mrmxf/clog/slogger"
	"github.com/opentsg/app-home/podserver"
	"github.com/opentsg/app-home/www"
)

// the default port to serve data is 8080
var Port = "8080"

var urlPrefix="/"
var mountPath="."

	var Logger *slog.Logger
func main() {
	podserver.InitialiseTemplates(www.StaticFs)
	
	r, _ := gommi.Bare()
	r.NewEmbedFileServer(www.StaticFs, urlPrefix, mountPath)
	r.Get("/",podserver.RouteLanding)
	// run the server in a thread
	slog.Info(fmt.Sprintf("Listening on port %s", Port))
	http.ListenAndServe("0.0.0.0:"+Port, r)
}

func init() {
	slogger.UsePrettyLogger(slog.LevelError)
	containerDataPath, runningInsideKoContainer := os.LookupEnv("KO_DATA_PATH")
	if runningInsideKoContainer {
		slog.Info(fmt.Sprintf("ko static data in folder %s", containerDataPath))
	}
}

