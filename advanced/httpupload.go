package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var templates *template.Template
var uploadDir string

func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.Execute(w, data)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//GET requests will be redirected
	case "GET":
		// redirect to . (http://localhost:8080/)
		http.Redirect(w, r, ".", 301)

	//POST takes the uploaded file(s) and saves it to disk
	case "POST":
		//get the multipart reader for the request
		reader, err := r.MultipartReader()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//copy each part to dir
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}

			//if FileName() is empty, skip this part
			if part.FileName() == "" {
				continue
			}
			dst, err := os.Create(uploadDir + part.FileName())
			if err == nil {
				defer dst.Close()
				log.Println("file upload started:", part.FileName())
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if _, err := io.Copy(dst, part); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			} else {
				log.Println("file uploaded:", part.FileName())
			}
		}
		//display success message via html template
		display(w, "upload", "Upload successful.")
	default:
		//only allow GET- and POST-Requests, otherwise respond with HTTP-Code 405
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RawQuery == "" && r.Method == "GET" {
		display(w, "upload", nil)
	}
}

func main() {
	uploadDir = "./"
	templates, _ = template.New("tmpl").Parse(templateString)

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/", formHandler)

	http.ListenAndServe(":8080", nil)
}

//this is the HTML template - usually this would be stored in a separate file
var templateString = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<title>File Upload Example</title>
		<style>
			body {
					font-family: Sans-serif;
					padding-bottom: 20px;
					background-color: #ffffff;
			}
			h1 a {
				color: black;
				text-decoration: none;
			}
			h1 {text-align: left; margin-bottom: 20px;}
			.message {font-weight:bold}
			fieldset {width:400px}
		</style>
	</head>
	<body>
		<div class="container">
			<h1><a href="./">File Upload Example</a></h1>
			<div class="message">{{.}}</div>
			<form class="form-signin" method="post" action="/upload" enctype="multipart/form-data">
					<fieldset>
						<input type="file" name="myfiles" id="myfiles" multiple="multiple">
						<input type="submit" name="submit" value="Upload">
				</fieldset>
			</form>
		</div>
	</body>
</html>`
