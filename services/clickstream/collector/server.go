package collector

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Config struct {
	Host    string
	Port    string
}

func handleEvent(
	logger *log.Logger,
	config *Config,
	producer *Producer,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			payload, err := io.ReadAll(r.Body)
			if err != nil || len(payload) == 0 {
				return
			}

			producer.Send(payload)
		}

		w.WriteHeader(http.StatusOK)
	})
}

func handlePixel(
	logger *log.Logger,
	config *Config,
	producer *Producer,
) http.Handler {
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

			producer.Send(payload)
		}

		tempData := []byte{
			0x47, 0x49, 0x46, 0x38, 0x39, 0x61, 0x01, 0x00,
			0x01, 0x00, 0x80, 0x00, 0x00, 0x00, 0x00, 0x00,
			0xFF, 0xFF, 0xFF, 0x21, 0xF9, 0x04, 0x01, 0x00,
			0x00, 0x00, 0x00, 0x2C, 0x00, 0x00, 0x00, 0x00,
			0x01, 0x00, 0x01, 0x00, 0x00, 0x02, 0x01, 0x44,
			0x00, 0x3B,
		}

		w.Header().Set("Content-Type", "image/gif")
		w.WriteHeader(http.StatusOK)
		w.Write(tempData)
	})
}

func NewServer(logger *log.Logger, config *Config, producer *Producer) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/e", handleEvent(logger, config, producer))
	mux.Handle("/pixel", handlePixel(logger, config, producer))

	return mux
}
