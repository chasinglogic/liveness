package server

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func httpCheck(resource string) error {
	// TODO: Should use context.Context of the HTTP request and should be more
	// configurable generally.
	log.Info().
		Str("resource", resource).
		Str("checkType", "http").
		Msg("testing resource")
	resp, err := http.Get(resource)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		err := fmt.Errorf("%s is unavailable: %d", resource, resp.StatusCode)
		return err
	}

	return nil
}
