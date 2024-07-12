package common

type Message struct {
	Code string `json:"code"`
	Data string `json:"message"`
}

type Input struct {
	FuncName string
	Data     string
	Option   string
}
