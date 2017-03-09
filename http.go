package fibweb

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// NewHTTPHandler create new http.Handler
func NewHTTPHandler(fibService FibonacciService) http.Handler {
	e := echo.New()

	e.GET("/api/fibonacci/numbers", func(c echo.Context) error {
		limitParam := c.QueryParam("limit")
		if limitParam == "" {
			limitParam = "5"
		}

		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		if limit < 0 {
			return c.String(http.StatusBadRequest, "Limit should be non negative")
		}

		nums, err := fibService.Seq(limit)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Service error: %v", err))
		}
		return c.JSON(http.StatusOK, nums)
	})

	return e
}
