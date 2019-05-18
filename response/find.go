package response

type FindResponse struct {
	Answers []FindAnswer `json:"answers"`
}

type FindAnswer struct {
	Paths    PathAnswer     `json:"paths"`
	Cheapest CheapestAnswer `json:"cheapest"`
}
type PathAnswer struct {
	From uint        `json:"from"`
	To   uint        `json:"to"`
	Path interface{} `json:"path"`
}
type CheapestAnswer struct {
	From uint        `json:"from"`
	To   uint        `json:"to"`
	Path interface{} `json:"path"`
}
