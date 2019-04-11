package main


import ( "log"
		 "time"
		 "net/http"
)

func StartServer() {

	// creates a handler (router or multiplexer)
	h := createHandler()

	// creates a HTTP server with default parameters
	s := createServer()

	// associate handler to server
	s.Handler = h

	// instanciates a HTTP server wrapped in a log fatal
	log.Fatal(s.ListenAndServe())
}

func StopServer() {

}

func createServer() (server *http.Server){


	// create a http server instance
	server = &http.Server {

		Addr: SERVER_ADDR,
		IdleTimeout:  100 * time.Millisecond,
		ReadTimeout:  100 * time.Millisecond,
		WriteTimeout: 100 * time.Millisecond,

	}

	return
}

