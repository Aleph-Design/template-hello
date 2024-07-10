package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/aleph-design/hello/pkg/config"
	"github.com/aleph-design/hello/pkg/models"
)

var functions = template.FuncMap{}

// Make config available to render package
var app *config.AppConfig

func NewRender(a *config.AppConfig) {
	app = a
}

// @ tmpl
// -	requested template name: "home.page.tmpl" is key in tmpCache
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	// define the essential structure to store go templates
	var tmpCache map[string]*template.Template
	if app.Production {
		tmpCache = app.TemplateCache
	} else {
		tmpCache, _ = CreateTemplateCache()
	}
	fmt.Println("\nCall template: ", tmpl)

	ts := tmpCache[tmpl]
	fmt.Println("\ntemplate set for: ", tmpl, "\nis: ", ts.DefinedTemplates())

	// get requested template from cache
	tmp, ok := tmpCache[tmpl]
	if !ok {
		// die when no pages available
		log.Fatal("could not load template from cache")
	}

	// tmp.Execute(os.Stdout, td)

	buf := new(bytes.Buffer)
	err := tmp.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// https://pkg.go.dev/text/template#hdr-Associated_templates

func CreateTemplateCache() (map[string]*template.Template, error) {

	tmpCache := map[string]*template.Template{}

	// get all files names that match the given pattern
	// so far this concerns only the *.page.tmpl files.
	fNames, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tmpCache, err
	}
// fmt.Println("\nfNames: ", fNames)

	// range through slice fNames
	for _, file := range fNames {
		fName := filepath.Base(file) // strip the path
	// fmt.Println("index: ", idx, fName)
	// index:  0 about.page.tmpl
	// index:  1 home.page.tmpl

		// create a tempate set of type *template.Template for file: 'file'
		// this will hold all templates associated with this 'file'
		tmpSet, err := template.New(fName).Funcs(functions).ParseFiles(file)
		if err != nil {
			return tmpCache, err
		}
	// fmt.Println("tmpSet: ", idx, tmpSet.DefinedTemplates())
	// tmpSet: 0 ; defined templates are: "content", "about.page.tmpl"
	// tmpSet: 1 ; defined templates are: "home.page.tmpl", "content"

		if fName == "home.page.tmpl" {
			tmpSet, err = tmpSet.ParseGlob("./templates/*.partial.tmpl")
			if err != nil {
				fmt.Println("95 - ParseGlob error")
				return tmpCache, err
			}
		}

		/*
			Glob returns the names of all files matching the given pattern or nil
			The pattern may describe hierarchical names, assuming the [Separator] is '/').
		*/
		// Glob wil find files: 
		// ./templates/base.layout.tmpl, ./templates/title.layout.tmpl, ./templates/css.layout.tmpl
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tmpCache, err
		}
	// fmt.Println("matches: ", idx, matches)
	// matches: 0 [templates/base.layout.tmpl templates/css.layout.tmpl templates/title.layout.tmpl]
	// matches: 1 [templates/base.layout.tmpl templates/css.layout.tmpl templates/title.layout.tmpl]

		// when there is a match, we add the template sets associated with found files
		// to the current template set
		// Glob wil find files: 
		// ./templates/base.layout.tmpl, ./templates/title.layout.tmpl, ./templates/css.layout.tmpl
		if len(matches) > 0 {
			tmpSet, err = tmpSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				fmt.Println("122 - ParseGlob error")
				return tmpCache, err
			}
		}
		tmpCache[fName] = tmpSet // add template set to cache map[key] = filename

		fmt.Println("")
	}

	// tpl := tmpCache["home.page.tmpl"]
	// fmt.Println("tmpCache: ", tpl.DefinedTemplates())
	// The complete generated and ready to execute template for page: home.page.tmpl
	// tmpCache:  ; defined templates are: "css", "title.layout.tmpl", "content", 
	//																		 "home.page.tmpl", "js", "base.layout.tmpl", 
	//																		 "base", "title", "css.layout.tmpl"

	return tmpCache, nil
}
