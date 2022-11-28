package dtos

type CreateGuitar struct {
	Manufacturer string `json:"manufacturer" binding:"required"`
	Model        string `json:"model" binding:"required"`
	Year         int    `json:"year" binding:"required"`
}
