package handlers

import (
	"net/http"
	"server/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// External Interractions
func (service *HTTPHandler) SaveAddress(ctx *gin.Context) {
	address := domain.AddressDto{}
	if err := ctx.ShouldBindJSON(&address); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": service.CompileErrors(err) })
		return
	}
	user, err := service.CurrentUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error() })
		return
	}

	profile, err := service.ExternalServicesAdapter.ReadProfile((user.ID).String())
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Profile record not found", "status": false})
		return
	}

	if err := service.ExternalServicesAdapter.SaveAddress(
		domain.Address{
			StreetNo:   address.StreetNo,
			StreetName: address.StreetName,
			Province:   address.Province,
			State:      address.State,
			ProfileID:  profile.ID,
		}); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": "User created successfully", "status": true})
}

func (service *HTTPHandler) ReadAddress(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Get an Address", "status": true })
}

func (service *HTTPHandler) ReadAddresses(ctx *gin.Context) {
	addresses, err := service.ExternalServicesAdapter.ReadAddresses()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{ "error": err.Error(), "status": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": addresses, "status": true })
}

func (service *HTTPHandler) PatchAddress(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Patch an Address"})
}

func (service *HTTPHandler) DeleteAddress(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": "Deleted an address successfully"})
}
