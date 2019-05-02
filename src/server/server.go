package server

import (
	"model"
)

type Server struct {
	model *model.Model
}

func NewServer(m *model.Model) (server *Server) {
	server = &Server{
		model: m,
	}

	return server
}
