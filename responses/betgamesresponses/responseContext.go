package betgamesresponses

import "github.com/labstack/echo"

var context echo.Context

func SetResponseContext(c echo.Context) {
	context = c
}

func GetResponseContext() echo.Context {
	return context
}
