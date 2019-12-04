package task

type Task struct {
	Id          int    `json:"id"`
	Uuid        string `json:"uuid"`
	Description string `json:"description"`
	Entry       string `json:"entry"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Modified    string `json:"modified"`
	Update      string `json:update`
	Project     string `json:"project"`
	Status      string `json:"status"`
}
