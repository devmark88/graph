package xml

import (
	"fmt"
	"strconv"

	"github.com/devmark88/unireg/models"

	"github.com/beevik/etree"
)

// Parse => Parse xml content
func (x *XML) Parse() (*models.Graph, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(x.Body); err != nil {
		return nil, err
	}
	graph := doc.SelectElement("graph")
	ch := graph.ChildElements()
	e2 := ch[2].FullTag()
	fmt.Println(len(ch))
	if e2 != "nodes" {
		if e2 == "edges" {
			return nil, fmt.Errorf("please put 'nodes' element before 'edges' element")
		}
		return nil, fmt.Errorf("cannot find node path in proper place. please see the sample xml file")
	}
	nodes := doc.FindElement("graph/nodes")
	edges := doc.FindElement("graph/edges")
	gID := doc.FindElement("graph/id")
	gName := doc.FindElement("graph/name")

	nodesM := []models.Node{}
	edgesM := []models.Edge{}

	g := models.Graph{}
	id, _ := strconv.ParseUint(gID.Text(), 10, 32)
	g.ID = uint(id)
	g.Name = gName.Text()
	for _, n := range nodes.ChildElements() {
		node := models.Node{}
		elem := n.FindElement("id")
		nid, _ := strconv.ParseUint(elem.Text(), 10, 32)
		node.ID = uint(nid)
		node.Name = n.FindElement("name").Text()
		nodesM = append(nodesM, node)
	}
	for _, e := range edges.ChildElements() {
		edge := models.Edge{}

		elem := e.FindElement("id")
		id, _ := strconv.ParseUint(elem.Text(), 10, 32)
		edge.ID = uint(id)

		elem = e.FindElement("from")
		id, _ = strconv.ParseUint(elem.Text(), 10, 32)
		edge.From = uint(id)

		elem = e.FindElement("to")
		id, _ = strconv.ParseUint(elem.Text(), 10, 32)
		edge.To = uint(id)

		elem = e.FindElement("cost")
		fid, _ := strconv.ParseFloat(elem.Text(), 64)
		edge.Weight = fid

		edgesM = append(edgesM, edge)
	}
	g.Edges = edgesM
	g.Nodes = nodesM

	return &g, nil
}
