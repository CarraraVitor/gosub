package subway

import (
	"gosub/route"
	"net/http"
)

var SubwayRouter route.Router

func init() {
    SubwayRouter.Register(
        route.Route{
            Path: "GET /subway/style/", 
            Handler: http.StripPrefix("/subway/style/", http.FileServer(http.Dir("subway/style/"))).ServeHTTP,
        },
    )
	SubwayRouter.Register(
		route.Route{
			Path:    "GET /{$}",
			Handler: Home,
		},
	)
	SubwayRouter.Register(
		route.Route{
			Path:    "GET /subway/nodes",
			Handler: ListNodes,
		},
	)
	SubwayRouter.Register(
		route.Route{
			Path:    "POST /subway/path",
			Handler: Path,
		},
	)
}
