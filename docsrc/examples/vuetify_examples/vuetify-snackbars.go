package vuetify_examples

// @snippet_begin(VuetifySnackBarsSample)
import (
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

func VuetifySnackBars(ctx *web.EventContext) (pr web.PageResponse, err error) {

	pr.Body = VContainer(
		VBtn("Show Snack Bar").OnClick("showSnackBar"),
		web.Portal().Name("snackbar"),
		snackbar("bottom", "success"),
	)

	return
}

func showSnackBar(ctx *web.EventContext) (er web.EventResponse, err error) {
	er.UpdatePortals = append(er.UpdatePortals,
		&web.PortalUpdate{
			Name: "snackbar",
			Body: snackbar("top", "red"),
		},
	)

	return
}

func snackbar(pos string, color string) *web.ScopeBuilder {
	return web.Scope(
		VSnackbar().Location(pos).Timeout(-1).Color(color).
			Attr("v-model", "locals.show").
			Children(
				h.Text("Hello, I am a snackbar"),
				web.Slot(
					VBtn("").Variant("text").
						Attr("@click", "locals.show = false").
						Children(VIcon("mdi-close")),
				).Name("actions"),
			),
	).VSlot("{ locals }").Init(`{ show: true }`)
}

var VuetifySnackBarsPB = web.Page(VuetifySnackBars).
	EventFunc("showSnackBar", showSnackBar)

const VuetifySnackBarsPath = "/samples/vuetify-snackbars"

// @snippet_end