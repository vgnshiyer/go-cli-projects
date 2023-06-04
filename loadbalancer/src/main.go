package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server interface {
	Address() string
	isAlive() bool
	serve(rw http.ResponseWriter, r *http.Request)
}

type myServer struct {
	address string
	proxy   *httputil.ReverseProxy
}

type loadbalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func initLoadbalancer(port string, servers []Server) *loadbalancer {
	return &loadbalancer{
		roundRobinCount: 0,
		port:            port,
		servers:         servers,
	}
}

func (s *myServer) Address() string { return s.address }

func (s *myServer) isAlive() bool { return true }

func (s *myServer) serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (lb *loadbalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	lb.roundRobinCount++
	// for !server.isAlive(){
	// 	lb.roundRobinCount++
	// 	server = lb.servers[lb.roundRobinCount % len(lb.servers)]
	// }
	if !server.isAlive() {
		server = lb.getNextAvailableServer()
	}
	return server
}

func (lb *loadbalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	target := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address: %q\n", target.Address())
	target.serve(rw, req)
}

func initServer(address string) *myServer {
	serverUrl, err := url.Parse(address)
	handleErr(err)

	return &myServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}

func main() {
	servers := []Server{
		initServer("https://vgnshiyer.github.io/"),
		initServer("https://vigneshfitness.com"),
		initServer("https://vigneshSDE.com"),
	}
	lb := initLoadbalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving at 'localhost:%s'\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
