package server

import (
	"net"
	"strings"
	"time"
)

func tcpCheck(resource string) error {
	s := strings.Split(resource, ":")
	_, host, port := s[0], s[1], s[2]

	timeout := time.Second
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		return err
	}

	if conn != nil {
		defer conn.Close()
	}

	return nil
}
