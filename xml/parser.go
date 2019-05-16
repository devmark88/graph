package xml

import (
	"fmt"
	"strconv"

	"github.com/devmark88/unireg/models"

	"github.com/beevik/etree"
)

type ParseResponse struct {
	Graph models.Graph
	Edges []models.Edge
}

// Parse => Parse xml content
func (x *XML) Parse(p *ParseResponse) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromString(x.Body); err != nil {
		return err
	}
	graph := doc.SelectElement("graph")
	ch := graph.ChildElements()
	e2 := ch[2].FullTag()
	fmt.Println(len(ch))
	if e2 != "nodes" {
		if e2 == "edges" {
			return fmt.Errorf("please put 'nodes' element before 'edges' element")
		}
		return fmt.Errorf("cannot find node path in proper place. please see the sample xml file")
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
		node := models.Find(uint(id), nodesM)
		edge.FromID = node.ID

		elem = e.FindElement("to")
		id, _ = strconv.ParseUint(elem.Text(), 10, 32)
		node = models.Find(uint(id), nodesM)
		edge.ToID = node.ID

		elem = e.FindElement("cost")
		fid, _ := strconv.ParseFloat(elem.Text(), 64)
		edge.Weight = fid

		edgesM = append(edgesM, edge)
	}
	// g.Edges = edgesM
	g.Nodes = nodesM
	p.Graph = g
	p.Edges = edgesM
	return nil
}
