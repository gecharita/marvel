package model

// GreekAPIResponse bla
type GreekAPIResponse struct {
	ΤotalCount int64  `json:"totalCount"`
	Status     string `json:"status"`
	GodList    []God  `json:"persons"`
}

// God bla
type God struct {
	PersonID int64  `json:"personID"`
	Name     string `json:"name"`
}
