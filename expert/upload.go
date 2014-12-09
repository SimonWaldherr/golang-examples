package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

var size int64 = 5 * 1024 * 1024
var html = template.Must(template.New("html").Parse(`
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8"/>
		<title>Golang File Upload</title>
	</head>
	<body>
		<form action="/upload" method="POST" enctype="multipart/form-data">
			<label for="file">File: </label>
			<input name="file" type="file"></input>
			<button type="submit">upload</button>
		</form>
	</body>
</html>
`))

func root(w http.ResponseWriter, r *http.Request) {
	err := html.Execute(w, nil)
	if err != nil {
		fmt.Print(err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	var path string
	if err := r.ParseMultipartForm(size); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusForbidden)
	}

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path = fmt.Sprintf("%s", fileHeader.Filename)
			buf, _ := ioutil.ReadAll(file)
			ioutil.WriteFile(path, buf, os.ModePerm)
		}
	}
	fmt.Printf("File \"%v\" uploaded\n", path)
}

func main() {
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/", root)
	fmt.Print(http.ListenAndServe(":8080", nil))
}
