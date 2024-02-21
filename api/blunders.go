package api

import (
	"html/template"
	"net/http"
	"os"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	"github.com/chippolot/jokegen"
)

func Blunders(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory(jokegen.Misunderstanding)
	if err != nil {
		panic(err)
	}

	model := common.StoryModel{
		Title:              "blunders.",
		Story:              result.Story,
		BackgroundColor:    "#154137",
		LogoFontLink:       "Fredoka:wght@600",
		LogoFontFamilyName: "Fredoka",
	}

	baseTmplDir := os.Getenv("BASE_TEMPLATE_DIR")
	tmplPath := baseTmplDir + "data/templates/story.gohtml"
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, model)
}
