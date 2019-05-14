package request

// Graph => Graph req
type Graph struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Nodes []Node `json:"nodes"`
	Edge  []Edge `json:"edges"`
}

// Node => Node req
type Node struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// Edge => Edge req
type Edge struct {
	ID     string `json:"id"`
	From   uint   `json:"from"`
	To     uint   `json:"to"`
	Weight uint   `json:"weight"`
}
