package vertex

import (
	"fmt"

	"github.com/devmark88/unireg/models"
	"github.com/starwander/goraph"
)

// Draw => Draw graph
func Draw(nodes []models.Node, edges []models.Edge) *goraph.Graph {
	graph := goraph.NewGraph()
	for _, n := range nodes {
		err := graph.AddVertex(n.ID, n.Name)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("vertex '%v' added \n", n.ID)
		}
	}
	for _, e := range edges {
		err := graph.AddEdge(e.FromID, e.ToID, e.Weight, nil)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("edge %v => %v added\n", e.FromID, e.ToID)
		}
	}
	// dist, path, err := graph.Yen(from, to, 10)
	// if err != nil {
	// fmt.Println(err)
	// return nil, nil, err
	// }
	// fmt.Println(dist)
	// fmt.Println(path)
	return graph
}
