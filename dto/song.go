package dto

//Song struct does a thing
type Song struct {
	Title      string `json:"title"`
	ID         string `json:"id"`
	Parts      string `json:"parts"`
	Difficulty int64  `json:"difficulty"`
}
