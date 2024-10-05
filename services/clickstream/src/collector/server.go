package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var tempData = []byte{
	0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00,
	0x01, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00,
	0xFF, 0xFF, 0xFF, 0x21, 0xF9, 0x04, 0x01, 0x00,
	0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x00,
	0x01, 0x00, 0x01, 0x00, 0x00, 0x02, 0x01, 0x44,
	0x00, 0x3B,
}

type Server struct {
	Addr          string
	AsyncProducer *AsyncProducer
}

func (s *Server) prompt() {
	promptString := `
       ┓┓        
    ┏┏┓┃┃┏┓┏╋┏┓┏┓
    ┗┗┛┗┗┗ ┗┗┗┛┛  %s
    - Serving at http://%s:%s
    - To close connection CTRL+C
    - %s
    `
	addr := strings.Split(s.Addr, ":")
	host := addr[0]
	if host == "" {
		host = "localhost"
	}

	port := addr[1]
	fmt.Printf(promptString, *env, host, port, time.Now())
	fmt.Println()
}

func (s *Server) handleHealthcheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) handleEvent() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			payload, err := io.ReadAll(r.Body)
			if err != nil || len(payload) == 0 {
				return
			}

			go func() {
				if err := s.AsyncProducer.SendMessage(payload, *kafkaTopic); err != nil {
					return
				}
			}()
		}

		w.WriteHeader(http.StatusOK)
	})
}

func (s *Server) handlePixel() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			values, err := url.ParseQuery(r.URL.Query().Encode())
			if err != nil || len(values) == 0 {
				w.WriteHeader(http.StatusOK)
				return
			}

			params := make(map[string]string)
			for key, value := range values {
				params[key] = value[0]
			}

			payload, err := json.Marshal(params)
			if err != nil || len(payload) == 0 {
				w.WriteHeader(http.StatusOK)
				return
			}

			go func() {
				if err := s.AsyncProducer.SendMessage(payload, *kafkaTopic); err != nil {
					return
				}
			}()
		}

		w.Header().Set("Content-Type", "image/gif")
		w.WriteHeader(http.StatusOK)
		w.Write(tempData)
	})
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/e", s.handleEvent())
	mux.Handle("/pixel", s.handlePixel())
	mux.Handle("/healthcheck", s.handleHealthcheck())

	return mux
}

func (s *Server) Run() error {
	httpServer := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler(),
	}

	s.prompt()
	return httpServer.ListenAndServe()
}

func (s *Server) Close() error {
	if err := s.AsyncProducer.Shutdown(); err != nil {
		return err
	}

	return nil
}
