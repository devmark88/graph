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

		dbg := app.DB.Create(g)
		return c.XML(http.StatusOK, dbg)
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
}
