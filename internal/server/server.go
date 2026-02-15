package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"

	"yourmodule/internal/commands"
	"yourmodule/internal/resp"
	"yourmodule/internal/storage"
)

type Server struct {
	Addr  string
	Store *storage.Store
}

func (s *Server) ListenAndServe() error {
	if s.Store == nil {
		return errors.New("nil Store")
	}

	ln, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return fmt.Errorf("listen %s: %w", s.Addr, err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// In a real server you might continue on temporary errors.
			return fmt.Errorf("accept: %w", err)
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)

	for {
		cmd, err := resp.ReadArray(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return
			}
			_, _ = conn.Write([]byte(resp.Error("Protocol error: " + err.Error())))
			return
		}

		out := commands.Dispatch(cmd, s.Store)

		if _, err := conn.Write([]byte(out)); err != nil {
			return
		}
	}
}
