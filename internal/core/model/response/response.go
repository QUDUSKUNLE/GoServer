package response

import "server/internal/core/entity/errorcode"

type Response struct {
	Data       		interface{}      		 `json:"data"`
	Status     		bool             		 `json:"status"`
	ErrorCode  		errorcode.ErrorCode  `json:"errorCode"`
	ErrorMessage 	string               `json:"errorMessage"`
}
