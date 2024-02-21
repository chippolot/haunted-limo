package api

import (
	"html/template"
	"net/http"
	"os"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	"github.com/chippolot/jokegen"
)

func Whammies(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory(jokegen.Slapstick)
	if err != nil {
		panic(err)
	}

	model := common.StoryModel{
		Title:              "whammies.",
		Story:              result.Story,
		BackgroundColor:    "#bb3a51",
		LogoFontLink:       "Courgette",
		LogoFontFamilyName: "Courgette",
	}

	baseTmplDir := os.Getenv("BASE_TEMPLATE_DIR")
	tmplPath := baseTmplDir + "data/templates/story.gohtml"
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, model)
}
