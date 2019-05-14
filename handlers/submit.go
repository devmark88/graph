package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/devmark88/unireg/xml"

	"github.com/devmark88/unireg/config"

	"github.com/labstack/echo"
)

// AddGraph => add new graph to database
func AddGraph(c echo.Context, app *config.AppContext) (err error) {
	// g := new(request.Graph)
	if b, err := ioutil.ReadAll(c.Request().Body); err == nil {
		err := xml.Validate(string(b))
		if err != nil {
			return c.JSON(403, err)
		}
		return c.XML(http.StatusOK, string(b))
	}
	return err
}
