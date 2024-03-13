package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/albertosfernandes/api-go-rest/controllers"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func HandleRequest() (err error) {

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Start HTTP server.
	srv := &http.Server{
		Addr:         ":8000",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(),
	}
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	// Wait for interruption.
	select {
	case err = <-srvErr:
		// Error when starting HTTP server.
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	err = srv.Shutdown(context.Background())
	return
}

func newHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// Register handlers.
	handleFunc("/", controllers.Home)
	handleFunc("/api/vms", controllers.GetVms)
	// handleFunc("/api/vm/{id}", controllers.GetVm)
	// handleFunc("/api/vm", controllers.Postvm)
	// handleFunc("/api/vm/{id}", controllers.Deletevm)
	// handleFunc("/api/vm/{id}", controllers.Putvm)
	// handleFunc("/api/networks", controllers.GetNetworks)
	// handleFunc("/api/network/{id}", controllers.GetNetwork)
	// handleFunc("/api/network", controllers.PostNetwork)
	// handleFunc("/api/vm/{id}", controllers.DeleteNetwork)
	// handleFunc("/api/vm/{id}", controllers.PutNetwork)
	// handleFunc("/api/storages", controllers.GetStorages)
	// handleFunc("/api/storage/{id}", controllers.GetStorage)
	// handleFunc("/api/storage", controllers.PostStorage)
	// handleFunc("/api/storage/{id}", controllers.DeleteStorage)
	// handleFunc("/api/storage/{id}", controllers.PutStorage)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
