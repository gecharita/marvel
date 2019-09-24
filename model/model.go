package model

// MarvelAPIResult bla
type MarvelAPIResult struct {
	Code            int64      `json:"code"`
	Status          string     `json:"status"`
	Etag            string     `json:"etag"`
	Copyright       string     `json:"copyright"`
	AttributionText string     `json:"attributionText"`
	AttributionHTML string     `json:"attributionHTML"`
	DataList        MarvelData `json:"data"`
}

// MarvelData bla
type MarvelData struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
	Total  int64 `json:"total"`
	Count  int64 `json:"count"`
	// results
}

// MarvelError bla
type MarvelError struct {
	Code   int64 `json:"code"`
	Status int64 `json:"status"`
}
