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
		res := new(xml.ParseResponse)
		err = x.Parse(res)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
		}
		if result := app.DB.Create(&res.Graph); result.Error != nil {
			return c.JSON(http.StatusBadRequest, result.GetErrors())
		}
		for _, e := range res.Edges {
			if result := app.DB.Create(&e); result.Error != nil {
				return c.JSON(http.StatusBadRequest, result.GetErrors())
			}
		}
		return c.XML(http.StatusOK, "OK")
	}
	return c.JSON(http.StatusBadRequest, echo.Map{"message": err})
}
