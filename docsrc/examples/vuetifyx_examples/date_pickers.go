package vuetifyx_examples

// @snippet_begin(VuetifyxDatetimePickers)

import (
	"context"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

type VXDateBuilder struct {
	tag *h.HTMLTagBuilder
}

func Vxdatepicker(children ...h.HTMLComponent) (r *VXDateBuilder) {
	r = &VXDateBuilder{
		tag: h.Tag("vx-datepicker").Children(children...),
	}
	return
}
func (b *VXDateBuilder) MarshalHTML(ctx context.Context) (r []byte, err error) {
	return b.tag.MarshalHTML(ctx)
}
func VuetifyxDatePickers(ctx *web.EventContext) (pr web.PageResponse, err error) {
	pr.Body = h.Div(
		Vxdatepicker(),
	)
	return
}

var DatePickersPB = web.Page(VuetifyxDatePickers)

const DatePickersPath = "/samples/vuetifyx_date_pickers"

// @snippet_end