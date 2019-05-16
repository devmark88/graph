package requests

type FindRequest struct {
	Queries []FindQuery `json:"queries"`
}

type FindQuery struct {
	Paths    PathQuery     `json:"paths"`
	Cheapest CheapestQuery `json:"cheapest"`
}
type PathQuery struct {
	Start uint `json:"start"`
	End   uint `json:"end"`
}
type CheapestQuery struct {
	Start uint `json:"start"`
	End   uint `json:"end"`
}
