package server

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

type server struct {
	workers []chan int
	results []chan error
}

func New(resources []string) server {
	workers := make([]chan int, len(resources))
	results := make([]chan error, len(resources))

	log.Info().Int("number_of_workers", len(resources)).Msg("starting background workers")
	for i, resource := range resources {
		workChan := make(chan int)
		resultChan := make(chan error)
		workers[i] = workChan
		results[i] = resultChan
		go worker(resource, workChan, resultChan)
	}

	return server{
		workers,
		results,
	}
}

func (s *server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	numTests := len(s.workers)

	for i := 0; i < numTests; i++ {
		s.workers[i] <- 0
	}

	for i := 0; i < numTests; i++ {
		tr := <-s.results[i]
		if tr != nil {
			writer.WriteHeader(http.StatusServiceUnavailable)
			writer.Write([]byte(tr.Error()))
			return
		}
	}

	writer.Write([]byte{})
}
