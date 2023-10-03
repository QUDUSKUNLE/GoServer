package services

import (
	"os"
	ports "server/internal/core/ports"
)

type ServicesHandler struct {
	Internal ports.ServerRepositoryPorts
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))
var tokenKey = os.Getenv("TOKEN_TTL")

func ServicesAdapter(serverRepository ports.ServerRepositoryPorts) *ServicesHandler {
	return &ServicesHandler{
		Internal: serverRepository,
	}
}
