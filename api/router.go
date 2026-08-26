package api

import "net/http"

func NewHandler(s *Server) http.Handler { return CORS(Logging(s.Handler())) }
func Routes() []string                  { return []string{"/health", "/records", "/consume", "/archive"} }
