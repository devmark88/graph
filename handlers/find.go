package handlers

import (
	"fmt"

	"github.com/devmark88/unireg/response"

	"github.com/devmark88/unireg/requests"

	"github.com/devmark88/unireg/vertex"

	"github.com/devmark88/unireg/models"

	"github.com/devmark88/unireg/config"

	"github.com/labstack/echo"
)

// FindPath => find path in graph
func FindPath(c echo.Context, app *config.AppContext) (err error) {
	res := response.FindResponse{}
	u := new(requests.FindRequest)
	if err = c.Bind(u); err != nil {
		return
	}
	gID := c.FormValue("graphId")
	fmt.Println("Finding for " + gID)
	var nodes []models.Node
	var edges []models.Edge
	app.DB.Where("graph_id = ?", gID).Find(&nodes)
	ids := []uint{}
	for _, n := range nodes {
		ids = append(ids, n.ID)
	}
	app.DB.Where("from_id in (?) OR to_id in (?)", ids, ids).Find(&edges)

	for _, q := range u.Queries {
		graph := vertex.Draw(nodes, edges)
		dist, path, err := graph.Yen(q.Paths.Start, q.Paths.End, 3)
		if err != nil {
			c.JSON(400, err)
		}
		r := response.FindAnswer{}
		r.Paths = response.PathAnswer{
			From: q.Paths.Start,
			To:   q.Paths.End,
			Path: path,
		}
		cp := float64(0)
		cpi := -1
		for idx, ch := range dist {
			fmt.Println(ch)
			fmt.Printf("CH < CP ==> %v < %v", float64(ch), float64(cp))
			if float64(ch) < float64(cp) {
				cp = ch
				cpi = idx
			}
		}
		fmt.Println(cpi)
		r.Cheapest = response.CheapestAnswer{}
		r.Cheapest = response.CheapestAnswer{
			From: q.Paths.Start,
			To:   q.Paths.End,
			Path: path[cpi],
		}
		res.Answers = append(res.Answers, r)
	}
	return c.JSON(200, res)
}
