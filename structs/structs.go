package structs

type JSONBlockchain struct {
	USD map[string]interface{} `json:"USD"`
	BRL map[string]interface{} `json:"BRL"`
	EUR map[string]interface{} `json:"EUR"`
}

type Currency struct {
	Idmode int64
	Nmmode string
}
