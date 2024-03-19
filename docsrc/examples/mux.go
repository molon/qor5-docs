package examples

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qor5/admin/presets"
	"github.com/qor5/docs/docsrc"
	"github.com/qor5/docs/docsrc/assets"
	"github.com/qor5/docs/docsrc/examples/admin_examples"
	"github.com/qor5/docs/docsrc/examples/presets_mux"
	"github.com/qor5/docs/docsrc/examples/web_vuetify_mux"
	"github.com/theplant/docgo"
	"net/http"
)

func Mux(mux *http.ServeMux, prefix string) http.Handler {
	web_vuetify_mux.Mux(mux, prefix)
	samplesMux := SamplesHandler(prefix)
	mux.Handle("/samples/",
		middleware.Logger(
			middleware.RequestID(
				samplesMux,
			),
		),
	)

	mux.Handle("/", docgo.New().
		MainPageTitle("QOR5 Document").
		Assets("/assets/", assets.Assets).
		DocTree(docsrc.DocTree...).
		Build(),
	)

	return mux
}

func SamplesHandler(prefix string) http.Handler {
	mux := http.NewServeMux()
	web_vuetify_mux.SamplesHandler(mux, prefix)
	presets_mux.SamplesHandler(mux, prefix)
	addGA := web_vuetify_mux.AddGA

	c23 := presets.New().AssetFunc(addGA)
	admin_examples.WorkerExampleMock(c23)
	mux.Handle(
		admin_examples.WorkerExamplePath+"/",
		c23,
	)

	c24 := presets.New().AssetFunc(addGA)
	admin_examples.ActionWorkerExampleMock(c24)
	mux.Handle(
		admin_examples.ActionWorkerExamplePath+"/",
		c24,
	)

	c27 := presets.New().AssetFunc(addGA)
	admin_examples.InternationalizationExample(c27)
	mux.Handle(
		admin_examples.InternationalizationExamplePath+"/",
		c27)
	c28 := presets.New().AssetFunc(addGA)
	admin_examples.LocalizationExampleMock(c28)
	mux.Handle(
		admin_examples.LocalizationExamplePath+"/",
		c28,
	)

	c29 := presets.New().AssetFunc(addGA)
	admin_examples.PublishExample(c29)
	mux.Handle(
		admin_examples.PublishExamplePath+"/",
		c29)

	return mux
}