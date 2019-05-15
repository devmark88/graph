package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/devmark88/unireg/config"
	"github.com/devmark88/unireg/xml"

	"github.com/labstack/echo"
)

// AddGraph => add new graph to database
func AddGraph(c echo.Context, app *config.AppContext) (err error) {
	x := xml.XML{}
	if b, err := ioutil.ReadAll(c.Request().Body); err == nil {
		x.Body = string(b)
		err := x.Validate()
		g, err := x.Parse()
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
		}
		if result := app.DB.Create(g); result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.GetErrors())
		}
		// graph := models.Graph{}
		// graph.ID = g.ID
		// graph.Name = g.Name
		// if result := app.DB.Table("graphs").Create(graph); result.Error != nil {
		// 	return c.JSON(http.StatusBadRequest, result.GetErrors())
		// }
		// nodes := g.Nodes
		// if result := app.DB.Create(nodes); result.Error != nil {
		// 	return c.JSON(http.StatusBadRequest, result.GetErrors())
		// }
		// edges := g.Edges
		// if result := app.DB.Create(edges); result.Error != nil {
		// 	return c.JSON(http.StatusBadRequest, result.GetErrors())
		// }
		return c.XML(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
}
