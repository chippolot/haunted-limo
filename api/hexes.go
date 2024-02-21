package api

import (
	"html/template"
	"net/http"
	"os"

	common "github.com/chippolot/haunted-limo/api/_pkg"
	"github.com/chippolot/jokegen"
)

func Hexes(w http.ResponseWriter, r *http.Request) {
	// Prep data provider
	connectionString := common.GetMySQLConnectionString()
	dataProvider := common.MakeSQLDataProvider(connectionString)
	defer dataProvider.Close()

	// Get most recent story
	result, err := dataProvider.GetMostRecentStory(jokegen.Hex)
	if err != nil {
		panic(err)
	}

	model := common.StoryModel{
		Title:              "hexes.",
		Story:              result.Story,
		BackgroundColor:    "#372445",
		LogoFontLink:       template.URL("Vesper+Libre:wght@400"),
		LogoFontFamilyName: "Vesper Libre",
		LogoFontStyle:      "normal",
		LogoFontWeight:     400,
		LogoFontSerif:      "serif",
	}

	baseTmplDir := os.Getenv("BASE_TEMPLATE_DIR")
	tmplPath := baseTmplDir + "data/templates/story.gohtml"
	tmpl := template.Must(template.ParseFiles(tmplPath))
	tmpl.Execute(w, model)
}
