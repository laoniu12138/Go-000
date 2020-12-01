package util

// PageRes 通用分页返回
type PageRes struct {
	Rows  interface{} `json:"rows"`
	Total int         `json:"total"`
}
