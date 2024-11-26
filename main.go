package main

import (
	AsciiArtWeb "AsciiArtWeb/ascii-art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	TextInput   string
	BannerInput string
	ArtResult   string
}

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// `http.HandleFunc` handles GET and POST requests only via `/` and `/ascii-art`.
// server is started on `http://localhost:8080`.
func main() {
	http.Handle("/templates/", http.FileServer(http.Dir(".")))

	http.HandleFunc("/", indexHandlerFunc)
	http.HandleFunc("/about", indexHandlerFunc)
	http.HandleFunc("/ascii-art", indexHandlerFunc)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

// This function loads the HTML template (index.html) using tmpl.ParseFiles().
// This handles both GET and POST requests only.
// GET loads and renders the input form using the HTML template.
// POST reads user's input text (`textInputName`) and selected style (`bannerName`).
// Then calls `AsciiArt()` to convert the input text to ASCII art.
func indexHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var err error
	pageData := PageData{}

	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" && r.URL.Path != "/about" {
		renderErrorPage(w, http.StatusNotFound, "The page you are looking for does not exist.")

		// http.Error(w, "ERROR 404. Incorrect URL.", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		pageData.TextInput = r.FormValue("textInputName")

		if pageData.TextInput == "" { // if input is empty
			renderErrorPage(w, http.StatusBadRequest, "Text input is empty.")
			return
		}

		pageData.BannerInput = r.FormValue("bannerName")

		pageData.ArtResult, err = AsciiArt(pageData.TextInput, pageData.BannerInput)
		if err != nil { // this handles all err we had on the terminal for ascii-art
			if err.Error() == "Character/s beyond standard ASCII printable characters code 32 to 126." {
				renderErrorPage(w, http.StatusBadRequest, err.Error())
				return
			} else {
				renderErrorPage(w, http.StatusInternalServerError, err.Error())
				return
			}
		}

		tmpl.ExecuteTemplate(w, "index.html", pageData)

	} else if r.Method == http.MethodGet {
		if r.URL.Path == "/about" {
			tmpl.ExecuteTemplate(w, "about.html", nil)
		} else {
			tmpl.ExecuteTemplate(w, "index.html", nil)
		}

	} else { // if other method is used
		renderErrorPage(w, http.StatusMethodNotAllowed, "Method Not Allowed.")
		return
	}
}

// this handles the ascii-art generator
func AsciiArt(inputStr, bannerType string) (string, error) {

	var artSrcFilename string

	switch bannerType {
	case "standard":
		artSrcFilename = "standard.txt"
	case "shadow":
		artSrcFilename = "shadow.txt"
	case "thinkertoy":
		artSrcFilename = "thinkertoy.txt"
	}

	for _, c := range inputStr {
		if c != '\n' && c != '\r' {
			if c < ' ' || c > '~' {
				return "", fmt.Errorf("Character/s beyond standard ASCII printable characters code 32 to 126.")
			}
		}
	}

	return (AsciiArtWeb.ConvToArt(inputStr, artSrcFilename))

}

// Function to render the error page
func renderErrorPage(w http.ResponseWriter, errorCode int, errorMessage string) {
	w.WriteHeader(errorCode) // Set the HTTP status code
	dataErr := struct {
		Code    int
		Message string
	}{
		Code:    errorCode,
		Message: errorMessage,
	}
	tmpl.ExecuteTemplate(w, "error.html", dataErr)
}
