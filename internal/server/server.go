package server

import (
	"fmt"
	"net/http"
	"usd-uah-emal-subsriber/pkg/errs"
	"usd-uah-emal-subsriber/pkg/logging"

	"github.com/gorilla/mux"
)

type Server struct {
	r   *mux.Router
	log logging.Logger

	exchange     ExchangeService
	subscription SubscriptionService
}

func NewServer(log logging.Logger, exchange ExchangeService, subscription SubscriptionService) *Server {
	r := mux.NewRouter()
	return &Server{r: r, log: log, exchange: exchange, subscription: subscription}
}

func (s *Server) Run(host, port string) error {
	r := mux.NewRouter()
	r.HandleFunc("/api/rate", s.Rate).Methods("GET")
	r.HandleFunc("/api/subscribe", s.Subscribe).Methods("POST")

	if err := http.ListenAndServe(host+":"+port, r); err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	fmt.Printf("Server started on :%s\n", port)

	return nil
}

func (s *Server) setErrorResponse(w http.ResponseWriter, err error) {
	localeCode, statusCode := errs.TopError(err)
	s.log.Error("Error occurred", "localeCode", localeCode, "statusCode", statusCode, "error", err)
	http.Error(w, fmt.Sprintf(`{"error": "%s"}`, localeCode), statusCode)
}
