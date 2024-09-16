package handler

type CreateOpeningRequest struct{
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    *bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    float64   `json:"salary"`
}