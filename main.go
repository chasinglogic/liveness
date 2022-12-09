package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/chasinglogic/liveness/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type resourceFlag []string

func (i *resourceFlag) String() string {
	return strings.Join(*i, ",")
}

func (i *resourceFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var resources resourceFlag

func main() {
	flag.Var(&resources, "resource", "A resource to watch for liveness, can be provided multiple times to watch multiple resources")
	structuredLogging := flag.Bool("structured-logging", false, "If provided will output structured JSON logs.")
	flag.Parse()

	if !*structuredLogging {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	s := server.New(resources)
	if err := http.ListenAndServe(":8080", &s); err != nil {
		fmt.Println("exited with error!", err.Error())
	}
}
