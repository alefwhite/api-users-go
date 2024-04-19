package handler

import (
	"encoding/json"
	"fmt"
	"github.com/alefwhite/api-users-go/internal/dto"
	"github.com/alefwhite/api-users-go/internal/handler/httperr"
	"github.com/alefwhite/api-users-go/internal/handler/validation"
	"log/slog"
	"net/http"
)

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Body == http.NoBody {
		slog.Error("body is empty", slog.String("package", "userhandler"))

		w.WriteHeader(http.StatusBadRequest)

		msg := httperr.NewBadRequestError("body is required")

		_ = json.NewEncoder(w).Encode(msg)

		return
	}

	var req dto.LoginDTO
	if r.Body != nil {
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			slog.Error("error to decode body", err, slog.String("package", "userhandler"))

			w.WriteHeader(http.StatusBadRequest)

			msg := httperr.NewBadRequestError("error to decode body")

			_ = json.NewEncoder(w).Encode(msg)

			return
		}
	}

	httpErr := validation.ValidateHttpData(req)
	if httpErr != nil {
		slog.Error(fmt.Sprintf("error to validate data: %v", httpErr), slog.String("package", "userhandler"))

		w.WriteHeader(httpErr.Code)

		_ = json.NewEncoder(w).Encode(httpErr)

		return
	}

	token, err := h.service.Login(r.Context(), req)
	if err != nil {
		if err.Error() == "user not found" || err.Error() == "invalid password" {
			w.WriteHeader(http.StatusUnauthorized)

			msg := httperr.NewUnauthorizedRequestError("invalid credentials")

			_ = json.NewEncoder(w).Encode(msg)

			return
		}

		w.WriteHeader(http.StatusBadRequest)

		msg := httperr.NewBadRequestError(err.Error())

		_ = json.NewEncoder(w).Encode(msg)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(token)
}
