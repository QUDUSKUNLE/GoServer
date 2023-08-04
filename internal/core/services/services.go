package services

import (
	ports "server/internal/core/ports"
)

type ServicesHandler struct {
	External ports.ServerRepositoryExternalPorts
}

func ExternalServicesAdapter(repo ports.ServerRepositoryExternalPorts) *ServicesHandler {
	return &ServicesHandler{
		External: repo,
	}
}
