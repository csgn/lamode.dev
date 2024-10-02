package collector

import (
	"log"
	"net/http"
)

func addHandlers(
	mux *http.ServeMux,
	logger *log.Logger,
	config *Config,
	producer *Producer,
) {
	mux.Handle("/e", handleEvent(logger, config, producer))
	mux.Handle("/pixel", handlePixel(logger, config, producer))
}

func NewServer(logger *log.Logger, config *Config, producer *Producer) http.Handler {
	mux := http.NewServeMux()
	addHandlers(mux, logger, config, producer)
	var handler http.Handler = mux
	return handler
}
