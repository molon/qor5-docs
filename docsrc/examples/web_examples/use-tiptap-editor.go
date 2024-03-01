package web_examples

// @snippet_begin(HelloWorldTipTapSample)
import (
	"github.com/qor5/ui/tiptap"
	"github.com/qor5/web"
	. "github.com/theplant/htmlgo"
	"github.com/yosssi/gohtml"
)

func HelloWorldTipTap(ctx *web.EventContext) (pr web.PageResponse, err error) {

	defaultValue := ctx.R.FormValue("Content1")
	if len(defaultValue) == 0 {
		defaultValue = `
			<h1>Hello</h1>
			<p>
				This is a nice editor
			</p>
			<ul>
			  <li>
				<p>
				  123
				</p>
			  </li>
			  <li>
				<p>
				  456
				</p>
			  </li>
			  <li>
				<p>
				  789
				</p>
			  </li>
			</ul>
`
	}

	pr.Body = Div(
		tiptap.TipTapEditor().
			Attr("v-model", "locals.Content1"),
		Hr(),
		Pre(
			gohtml.Format(ctx.R.FormValue("Content1")),
		).Style("background-color: #f8f8f8; padding: 20px;"),
		Button("Submit").Style("font-size: 24px").
			Attr("@click", web.POST().EventFunc("refresh").Go()),
	)

	return
}

func refresh(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.Reload = true
	return
}

var HelloWorldTipTapPB = web.Page(HelloWorldTipTap).
	EventFunc("refresh", refresh)

const HelloWorldTipTapPath = "/samples/hello_world_tiptap"

// @snippet_end
