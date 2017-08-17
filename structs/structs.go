package structs

type JSONBlockchain struct {
	USD map[string]interface{} `json:"USD"`
	BRL map[string]interface{} `json:"BRL"`
	EUR map[string]interface{} `json:"EUR"`
}
