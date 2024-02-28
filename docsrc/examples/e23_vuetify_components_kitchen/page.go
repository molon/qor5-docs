package e23_vuetify_components_kitchen

// @snippet_begin(VuetifyComponentsKitchen)

import (
	"github.com/qor5/docs/docsrc/utils"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	h "github.com/theplant/htmlgo"
)

var globalCities = []string{"Tokyo", "Hangzhou", "Shanghai"}

type formVals struct {
	Cities1 []string
	Cities2 []string

	MyItem string
}

var fv = formVals{
	Cities1: []string{
		"TK",
		"LD",
	},

	Cities2: []string{
		"Hangzhou",
		"Shanghai",
	},

	MyItem: "VItem2",
}

func VuetifyComponentsKitchen(ctx *web.EventContext) (pr web.PageResponse, err error) {

	var chips h.HTMLComponents
	for _, city := range globalCities {
		chips = append(chips,
			VChip(h.Text(city)).
				Closable(true).
				Attr("@click:close", web.POST().EventFunc("removeCity").Query("city", city).Go()),
		)
	}

	pr.Body = VContainer(
		utils.PrettyFormAsJSON(ctx),
		web.Scope(
			h.H1("Chips delete"),
			chips,

			h.H1("Chips group"),
			h.H2("Cities1"),
			VChipGroup(
				VChip(
					h.Text("Hangzhou"),
					VIcon("mdi-star").End(true),
				).Value("HZ"),
				VChip(h.Text("Shanghai")).Value("SH").Filter(true).Label(true),
				VChip(h.Text("Tokyo")).Value("TK").Filter(true),
				VChip(h.Text("New York")).Value("NY"),
				VChip(h.Text("London")).Value("LD"),
			).SelectedClass("bg-indigo").
				Attr("v-model", "locals.Cities1").
				Multiple(true),
			h.H2("Cities2"),
			VAutocomplete().
				Items(globalCities).
				Chips(true).
				ClosableChips(true).
				Multiple(true).
				Attr("v-model", "locals.Cities2"),

			h.H1("Items Group"),

			VItemGroup(
				VContainer(
					VRow(
						VCol(
							VItem(
								VCard(
									VCardTitle(h.Text("Item1")),
								).
									Height(200).
									Attr(":class", "['d-flex align-center', selectedClass]").
									Attr("@click", "toggle"),
							).Value("VItem1").Attr("v-slot", "{isSelected, selectedClass, toggle}"),
						),

						VCol(
							VItem(
								VCard(
									VCardTitle(h.Text("Item2")),
								).
									Height(200).
									Attr(":class", "['d-flex align-center', selectedClass]").
									Attr("@click", "toggle"),
							).Value("VItem2").Attr("v-slot", "{isSelected, selectedClass, toggle}"),
						),
					),
				),
			).
				SelectedClass("bg-primary").
				Attr("v-model", "locals.MyItem"),

			VBtn("Submit").
				OnClick("submit"),
		).VSlot("{ locals, plaidForm }").Init(h.JSONString(fv)),
	)
	return
}

func submit(ctx *web.EventContext) (r web.EventResponse, err error) {
	fv = formVals{}
	ctx.MustUnmarshalForm(&fv)

	r.Reload = true
	return
}

func removeCity(ctx *web.EventContext) (r web.EventResponse, err error) {
	city := ctx.R.FormValue("city")
	var newCities = make([]string, 0)
	for _, c := range globalCities {
		if c != city {
			newCities = append(newCities, c)
		}
	}
	globalCities = newCities
	r.Reload = true
	return
}

var VuetifyComponentsKitchenPB = web.Page(VuetifyComponentsKitchen).
	EventFunc("removeCity", removeCity).
	EventFunc("submit", submit)

const VuetifyComponentsKitchenPath = "/samples/vuetify-components-kitchen"

// @snippet_end
