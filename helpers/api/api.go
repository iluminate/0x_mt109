package api

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net/http"
)

func RenderError(ctx *fasthttp.RequestCtx, error string, statusCode int) {
	ctx.Write([]byte(error))
	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetConnectionClose()
}
func RenderSuccess(ctx *fasthttp.RequestCtx, body interface{}, statusCode int) {
	enc := json.NewEncoder(ctx.Response.BodyWriter())
	if err := enc.Encode(&body); err != nil {
		fmt.Printf("ERROR Marshal Response %v :", err)
		RenderError(ctx, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.Response.Header.Add("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	ctx.Response.Header.Add("Access-Control-Allow-Origin", "*")
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)
	ctx.SetConnectionClose()
}