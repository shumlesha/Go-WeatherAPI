package user

type ChangeCityRequest struct {
	City string `json:"city" binding:"required"`
}
