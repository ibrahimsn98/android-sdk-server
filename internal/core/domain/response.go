package domain

type Response struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}
