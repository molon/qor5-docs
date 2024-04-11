package examples_presets

import (
	"net/http"

	"github.com/qor5/admin/presets"
	. "github.com/qor5/ui/vuetify"
	"github.com/qor5/web"
	"github.com/qor5/x/perm"
	h "github.com/theplant/htmlgo"
	"gorm.io/gorm"
)

// @snippet_begin(PresetsPermissionsSample)
type User struct {
	ID       uint
	Username string
}

type Group struct {
	ID   uint
	Name string
}

func PresetsPermissions(b *presets.Builder) (
	cust *presets.ModelBuilder,
	cl *presets.ListingBuilder,
	ce *presets.EditingBuilder,
	dp *presets.DetailingBuilder,
	db *gorm.DB,
) {
	cust, cl, ce, dp, db = PresetsDetailPageCards(b)
	b.URIPrefix(PresetsPermissionsPath)

	b.ProfileFunc(func(ctx *web.EventContext) h.HTMLComponent {
		return VMenu(
			web.Slot(
				VBtn("").
					Icon(true).
					Attr("v-bind", "props").
					Children(
						VIcon("mdi-account"),
					).Class("ml-2"),
			).Name("activator").Scope("{  props }"),

			VList(
				VListItem(
					VListItemTitle(h.Text("Logout")),
				),
			),
		)
	})

	perm.Verbose = true
	b.Permission(perm.New().
		Policies(
			perm.PolicyFor("editor").WhoAre(perm.Allowed).ToDo(perm.Anything).On(perm.Anything),
			perm.PolicyFor("editor").WhoAre(perm.Denied).ToDo(presets.PermRead...).On("*user_management*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo(presets.PermCreate, presets.PermDelete).On("*customers*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo(presets.PermCreate, presets.PermUpdate).On("*companies*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo(presets.PermUpdate).On("*customers:*:company_id*"),
			perm.PolicyFor("editor").WhoAre(perm.Denied).
				ToDo("*bulk_actions:delete").On("*:customers*"),
		).
		SubjectsFunc(func(r *http.Request) []string {
			return []string{"editor"}
		}))

	err := db.AutoMigrate(&User{}, &Group{})
	if err != nil {
		panic(err)
	}

	b.MenuGroup("User Management").SubItems("user", "group")
	b.Model(&User{})
	b.Model(&Group{})
	return
}

const PresetsPermissionsPath = "/samples/presets-permissions"

// @snippet_end
