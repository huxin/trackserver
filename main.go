// http server that serve the same track image and record all the information
package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var (
	logFile = "access.log"
	logger  *os.File
	err     error

	flagPort = flag.String("port", "80", "port to listen on")
)

func createTrackRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/{path:.*}", tracker)
	return router
}

// tracker tracks emails
// stores url and other information in the log
func tracker(w http.ResponseWriter, r *http.Request) {
	t := fmt.Sprintf("%d", time.Now().Unix())
	tStr := fmt.Sprintf("[%s]", time.Now().Format(time.RFC3339))
	addr := r.RemoteAddr
	uri := r.RequestURI
	ua := r.UserAgent()
	referer := r.Referer()
	url := r.URL.String()

	output := strings.Join([]string{
		tStr, t, addr, uri, url, referer, ua,
	}, ",")

	// serve the output

	fmt.Println(output)
	logger.Write([]byte(output + "\n"))

	http.ServeFile(w, r, "pixel.png")
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func main() {
	flag.Parse()
	addr := ":" + *flagPort

	logger, err = os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer logger.Close()

	fmt.Println("Listen on:", addr)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	http.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)}, createTrackRouter())
}
