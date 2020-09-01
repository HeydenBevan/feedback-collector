// Package linker will provide the handler to serve a Static Page to manage DomainLinks
// When calling this, you should be calling in the specific format:
// Read: /linker/
// Create: /linker/create/
// Update: /newlink/edit/{LinkID}
// Delete: /newlink/delete/{LinkID}
package linker

import (
	"html/template"
	"log"
	"net/http"
)

type linkData struct {
	Path   string
	Method string
	LinkID string
}

// Assigning meeting to some Index Values
const (
	MethodPath = 2
	PathLinkID = 3
)

// LinkHandler provides the Front-End UI to manage DomainLinks
func LinkHandler(w http.ResponseWriter, req *http.Request) {
	tpl := template.Must(template.ParseFiles("linker/index.gohtml"))
	// url := strings.Split(req.URL.Path, "/")
	// var lid string
	// if url[PathLinkID] != "" {
	// 	lid = url[PathLinkID]
	// }

	// var td linkData
	// td.Path = req.URL.Path
	// td.Method = url[MethodPath]
	// td.LinkID = lid

	// td := templateData{
	// 	req.URL.Path,
	// 	m,
	// 	lid,
	// }
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}
