package render

import (
	"GoInActionAssignment/form"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// TemplateData stores data to be used in Templates
type TemplateData struct {
	Data map[string]interface{}
	Form *form.Form
}

//To waypoint to which template are executed.
//If successfully executed, handler function will not trigger
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData) error {

	ts, err := template.ParseFiles(fmt.Sprintf("./templates/%s", tmpl),
		"./templates/base.layout.html",
		"./templates/header.layout.html")

	log.Println("Logging: Redirecting to", tmpl)
	if err != nil {
		return fmt.Errorf("ParseTemplate: Unable to find template pages: %w", err)
	}

	if err := ts.Execute(w, td); err != nil {
		return fmt.Errorf("ParseTemplate: Unable to execute template: %w", err)
	}

	return nil
}
