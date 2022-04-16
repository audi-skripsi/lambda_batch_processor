package dto

type EventExtractionRequest struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type EventExtractionResponse struct {
	TotalExtracted int `json:"total_extracted"`
}
