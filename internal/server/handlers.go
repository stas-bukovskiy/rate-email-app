package server

import (
	"fmt"
	"net/http"
	"usd-uah-emal-subsriber/internal/localecode"
	"usd-uah-emal-subsriber/pkg/errs"
)

func (s *Server) Rate(w http.ResponseWriter, r *http.Request) {
	rate, err := s.exchange.Rate()
	if err != nil {
		s.setErrorResponse(w, err)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf(`{"rate": %.2f}`, rate)))
	if err != nil {
		s.setErrorResponse(w, errs.F("failed to write response: %w", err))
		return
	}
}

func (s *Server) Subscribe(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		s.setErrorResponse(w, errs.F("email is required").SetCode(localecode.CodeEmailRequired).
			SetStatusCode(http.StatusBadRequest))
		return
	}

	if err := s.subscription.CreateSubscription(r.Context(), email); err != nil {
		s.setErrorResponse(w, err)
		return
	}

	if _, err := w.Write([]byte(`{"message": "Subscribed"}`)); err != nil {
		s.setErrorResponse(w, errs.F("failed to write response: %w", err))
		return
	}
}
