package server

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

func getTestType(resource string) string {
	// TODO: this is probably too simple.
	segments := strings.Split(resource, ":")
	return segments[0]
}

func worker(resource string, trigger chan int, result chan error) {
	for {
		log.Debug().
			Str("resource", resource).
			Msg("worker ready for liveness checks")

		// Wait to be asked to do a check
		<-trigger

		log.Debug().
			Str("resource", resource).
			Msg("starting liveness check")

		var testResult error

		testType := getTestType(resource)
		switch testType {
		case "http":
			fallthrough
		case "https":
			testResult = httpCheck(resource)
		case "tcp":
			testResult = tcpCheck(resource)
		default:
			testResult = fmt.Errorf("Unsupported check type %s: %s", testType, resource)
		}

		if testResult != nil {
			log.Error().
				Str("resource", resource).
				Msg(testResult.Error())
		}

		result <- testResult
	}
}
