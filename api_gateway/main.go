package api_gateway

import "flag"

func main() {
	port := flag.String("port", "8080", "port to serve on")
	flag.Parse()

	server := NewServer(*port)
	server.Start()
}
