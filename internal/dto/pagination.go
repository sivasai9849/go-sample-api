package dto

type PaginatedResponse struct {
    Data       interface{} `json:"data"`
    Page       int         `json:"page"`
    Limit      int         `json:"limit"`
    TotalItems int64       `json:"total_items"`
}