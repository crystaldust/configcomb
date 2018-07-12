package main

import (
	"fmt"
	"net/http"

	"github.com/juzhen/k8s-client-test/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/virtualservices", handlers.ListVirtualServices)
	mux.HandleFunc("/destinationrules", handlers.ListDestinationRules)
	mux.HandleFunc("/deployments", handlers.GetDeployments)
	mux.HandleFunc("/ratelimit", handlers.HandleRateLimit)

	fmt.Println("listen on port 8882")
	if err := http.ListenAndServe(":8882", mux); err != nil {
		fmt.Println(err)
	}

}
