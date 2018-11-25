// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/RoboCup-SSL/ssl-game-controller/internal/app/controller"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gobuffalo/packr"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strings"

	"github.com/RoboCup-SSL/ssl-game-controller/restapi/operations"
	"github.com/RoboCup-SSL/ssl-game-controller/restapi/operations/card"
)

//go:generate swagger generate server --target .. --name SslGameController --spec ../swagger.yml

var gameController = controller.NewGameController()

func configureFlags(api *operations.SslGameControllerAPI) {
}

func configureAPI(api *operations.SslGameControllerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	gameController.Run()

	api.CardAddCardHandler = card.AddCardHandlerFunc(func(params card.AddCardParams) middleware.Responder {
		return middleware.NotImplemented("operation card.AddCard has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler

	box := packr.NewBox("../dist")
	uiHandler := http.FileServer(box)
	if !box.Has("index.html") {
		log.Print("Backend-only version started. Run the UI separately or get a binary that has the UI included")
	}

	return handleCORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/control") {
			gameController.ApiServer.WsHandler(w, r)
		} else if strings.HasPrefix(r.URL.Path, "/api") || strings.HasPrefix(r.URL.Path, "/swagger") {
			handler.ServeHTTP(w, r)
		} else {
			uiHandler.ServeHTTP(w, r)
		}
	}))
}
