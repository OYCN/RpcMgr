package config

type Ret struct {
	Status bool `json:"status"`
	Action string `json:"action"`
	Data interface{} `json:"data"`
}

type RedirectAction struct {
	Url string `json:"url"`
	Args []string `json:"args"`
}
