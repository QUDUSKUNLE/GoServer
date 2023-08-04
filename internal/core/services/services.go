package services

import (
	ports "server/internal/core/ports"
)

type ServicesHandler struct {
	External ports.ServerRepositoryExternalPorts
	Internal ports.ServerRepositoryInternalPorts
}

func ExternalServicesAdapter(repo ports.ServerRepositoryExternalPorts) *ServicesHandler {
	return &ServicesHandler{
		External: repo,
	}
}


func InternalServicesAdapter(repo ports.ServerRepositoryInternalPorts) *ServicesHandler {
	return &ServicesHandler{
		Internal: repo,
	}
}
