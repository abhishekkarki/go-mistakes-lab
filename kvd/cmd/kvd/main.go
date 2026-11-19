// Command kvd runs the kvd key-value server.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/abhishekkarki/go-mistakes-lab/kvd/internal/api"
	"github.com/abhishekkarki/go-mistakes-lab/kvd/internal/store"
)

func main() {
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()

	h := api.NewHandler(store.New())

	log.Printf("kvd listening on %s", *addr)

	// TODO(day 22, #81): http.ListenAndServe has no timeouts and no graceful
	// shutdown. Replace it with an http.Server (ReadHeader/Read/Write/Idle
	// timeouts), signal.NotifyContext and srv.Shutdown.
	// TODO(day 31, #98): serve net/http/pprof on a separate admin port.
	if err := http.ListenAndServe(*addr, h); err != nil {
		log.Fatal(err)
	}
}
