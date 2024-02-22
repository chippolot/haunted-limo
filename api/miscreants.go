package api

import (
	"html/template"
	"net/http"
	"os"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	"github.com/chippolot/jokegen"
)

func Miscreants(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory(jokegen.Creature)
	if err != nil {
		panic(err)
	}

	model := common.StoryModel{
		Title:              "miscreants.",
		Story:              result.Story,
		BackgroundColor:    "#048c7c",
		LogoFontLink:       "Carter+One",
		LogoFontFamilyName: "Carter One",
		LogoFontStyle:      "normal",
		LogoFontWeight:     400,
		LogoFontSerif:      "system-ui",
	}

	baseTmplDir := os.Getenv("BASE_TEMPLATE_DIR")
	tmplPath := baseTmplDir + "data/templates/story.gohtml"
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, model)
}
