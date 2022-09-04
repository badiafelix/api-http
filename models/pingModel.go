package models

type Input struct {
	Url string `json:"url"`
}

type Result struct {
	Url         string
	Status_Code string
	Duration    string
}

type OutputContent struct {
	Status  string
	Message string
	Data    []Result
}
