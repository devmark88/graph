package handlers

import (
	"fmt"

	"github.com/devmark88/unireg/requests"

	"github.com/devmark88/unireg/vertex"

	"github.com/devmark88/unireg/models"

	"github.com/devmark88/unireg/config"

	"github.com/labstack/echo"
)

type response struct {
}

// FindPath => find path in graph
func FindPath(c echo.Context, app *config.AppContext) (err error) {
	u := new(requests.FindRequest)
	resPath := make(map[int]interface{})
	resCheap := make(map[int]interface{})
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

	for idx, q := range u.Queries {
		graph := vertex.Draw(nodes, edges)
		dist, path, err := graph.Yen(q.Paths.Start, q.Paths.End, 3)

		if err != nil {
			c.JSON(400, err)
		}
		resPath[idx] = path
		resCheap[idx] = dist
	}

	return c.JSON(200, map[string]interface{}{"paths": resPath, "cheapest": resCheap})
}
