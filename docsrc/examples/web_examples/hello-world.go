package web_examples

// @snippet_begin(HelloWorldSample)
import (
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
)

func HelloWorld(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = H1("Hello World")
	return
}

var HelloWorldPB = web.Page(HelloWorld) // this is already a http.Handler

const HelloWorldPath = "/samples/hello_world"

// @snippet_end
