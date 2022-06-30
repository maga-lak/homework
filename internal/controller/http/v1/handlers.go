package v1

import (
	"context"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"homework/internal/controller/http/v1/request"
	"homework/internal/service"
	"io"
	"net/http"
)

type Server struct {
	authService *service.AuthService
	log         *zap.SugaredLogger
}

func NewServer(
	authService *service.AuthService,
	log *zap.SugaredLogger,
) *Server {
	return &Server{
		authService: authService,
		log:         log,
	}
}

func (s *Server) Auth(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		s.writeResponse(errors.New("internal error"), w, http.StatusInternalServerError)
		return
	}

	req := &request.AuthRequest{}
	if err = json.Unmarshal(data, req); err != nil {
		s.writeResponse(errors.New("json is not correct"), w, http.StatusUnprocessableEntity)
		return
	}
	//if err := req.Validate(); err != nil {
	//	s.writeResponse(fmt.Errorf("validation err: %s", err.Error()), w, http.StatusBadRequest)
	//	return
	//}

	ctx := context.Background()

	view, err := s.authService.Authorize(ctx, req)
	if err != nil {
		if errors.Is(err, service.UserFoundErr) || errors.Is(err, service.InvalidPasswordErr) {
			s.writeResponse(err.Error(), w, http.StatusBadRequest)
			return
		}
		s.writeResponse(errors.New("internal error"), w, http.StatusInternalServerError)
		return
	}

	s.writeResponse(view, w, http.StatusOK)
}

func (s *Server) writeResponse(response interface{}, w http.ResponseWriter, code int) {
	w.Header().Set("content-type", "application-json; charset=utf-8")
	w.WriteHeader(code)

	data, _ := json.Marshal(struct {
		Data interface{} `json:"data"`
	}{Data: response})
	if _, err := w.Write(data); err != nil {
		s.log.Errorf("failed to http response: %s", err.Error())
	}
}
