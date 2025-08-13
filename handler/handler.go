package handler

import (
	z "github.com/Oudwins/zog"
	"github.com/Oudwins/zog/zhttp"
	"github.com/labstack/echo/v4"
)

func ParseReq(c echo.Context, schema *z.StructSchema, destPtr any) (sanitizedErrs map[string][]string, ok bool) {
	if errs := schema.Parse(zhttp.Request(c.Request()), destPtr); errs != nil {
		sanitizedErrs = z.Issues.SanitizeMap(errs)
		if sanitizedErrs["$first"][0] == "" {
			sanitizedErrs["$first"][0] = "Invalid request"
			sanitizedErrs["$root"][0] = "Invalid request"
		}
		return sanitizedErrs, false
	}

	return nil, true
}
