package model

// PageInfo common page info.
type PageInfo struct {
	Count       int         `json:"count"`
	CurrentPage int         `json:"currentPage,omitempty"`
	Item        interface{} `json:"item"`
}
