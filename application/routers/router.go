package routers

import (
	"0x_mt109/application/container"
	"0x_mt109/application/controllers"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type pathBuilder func(string) string

type httpRouter struct {
	contextPath string
	ActorHandler *controllers.ActorHandler
	handleCORS   func(fasthttp.RequestHandler) fasthttp.RequestHandler
}

func NewHttpRouter(contextPath string) *httpRouter {
	return &httpRouter{
		contextPath:  contextPath,
		ActorHandler: container.ActorHandler(),
	}
}

func (httpRouter *httpRouter) Handler() fasthttp.RequestHandler {

	pathV1 := httpRouter.pathVersion("v1")
	fmt.Printf("The base url is: [%s]\n", pathV1(""))
	router := fasthttprouter.New()
	router.RedirectTrailingSlash = true
	router.GET(pathV1("/"), baseHandler)
	router.GET(pathV1("/actors"), httpRouter.ActorHandler.GetAll)
	router.PUT(pathV1("/actors/:id"), httpRouter.ActorHandler.Update)
	router.POST(pathV1("/actors"), httpRouter.ActorHandler.Create)
	router.DELETE(pathV1("/actors/:id"), httpRouter.ActorHandler.Delete)
	if httpRouter.handleCORS != nil {
		return httpRouter.handleCORS(router.Handler)
	}
	return router.Handler
}

func (httpRouter *httpRouter) EnableCORS(origins string) *httpRouter {
	httpRouter.handleCORS = func(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			if string(ctx.Method()) == fasthttp.MethodOptions {
				ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
				ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
				ctx.Response.Header.Add("Access-Control-Allow-Origin", origins)
				ctx.SetStatusCode(fasthttp.StatusOK)
				ctx.SetConnectionClose()
				return
			}
			handler(ctx)
		}
	}
	return httpRouter
}

func baseHandler(ctx *fasthttp.RequestCtx) {
	ctx.Write([]byte("Api is Up!"))
}

func (httpRouter *httpRouter) pathVersion(version string) pathBuilder {
	return path(httpRouter.contextPath + "/api/" + version)
}

func path(basePath string) pathBuilder {
	return func(path string) string {
		return basePath + path
	}
}