package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/template"
	"time"
)

type Person []struct {
	Name struct {
		Title  string `json:"title"`
		First  string `json:"first"`
		Middle string `json:"middle"`
		Last   string `json:"last"`
	} `json:"name"`
	Address struct {
		Street   string `json:"street"`
		Zip      string `json:"zip"`
		Location string `json:"location"`
		Country  string `json:"country"`
	} `json:"address"`
	Contact struct {
		Phone  string `json:"phone"`
		Mobile string `json:"mobile"`
		Email  string `json:"email"`
		Web    string `json:"web"`
	} `json:"contact"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

func shortMiddleName(input string) string {
	if len(input) > 0 {
		return input[0:1] + "."
	}
	return ""
}

func getDate() string {
	return time.Now().Format("2006-01-02")
}

func main() {
	jsonstr, latexstr := getExampleData()

	person := Person{}

	err := json.Unmarshal([]byte(jsonstr), &person)
	if err != nil {
		log.Println(err)
	}

	fmap := template.FuncMap{
		"shortMiddleName": shortMiddleName,
		"getDate":         getDate,
		"getPlace": func() string {
			return "Munich"
		},
	}

	t := template.New("letter").Delims("<<", ">>").Funcs(fmap)

	_, err = t.Parse(latexstr)
	if err != nil {
		log.Println(err)
	}

	f, err := os.OpenFile("./template_latex.tex", os.O_WRONLY|os.O_CREATE, 0755)
	if err == nil {
		err = t.Execute(f, person)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(err)
	t.Execute(os.Stdout, person)
}

func getExampleData() (string, string) {
	jsonstr := `
[
  {
    "Name": {
      "Title": "Dr.",
      "First": "John",
      "Middle": "F.",
      "Last": "Doe"
    },
    "Address": {
      "Street": "Musterweg 2",
      "Zip": "12345",
      "Location": "Bielefeld",
      "Country": "Germany"
    },
    "Gender": "m",
    "Birthday": "1970-01-01"
  },
  {
    "Name": {
      "Title": "Prof. Dr.",
      "First": "Jane",
      "Middle": "A.",
      "Last": "Doe"
    },
    "Address": {
      "Street": "Musterweg 3",
      "Zip": "12345",
      "Location": "Bielefeld",
      "Country": "Germany"
    },
    "Gender": "f",
    "Birthday": "1980-01-03"
  },
  {
    "Name": {
      "Title": "Dr.",
      "First": "Sascha",
      "Middle": "",
      "Last": "Doe"
    },
    "Address": {
      "Street": "Musterweg 4",
      "Zip": "12345",
      "Location": "Bielefeld",
      "Country": "Germany"
    },
    "Gender": "d",
    "Birthday": "1978-01-02"
  }
]`
	latexstr := `
\documentclass[
    sender,
    paper=a4,
    version=last,
    fontsize=12pt,
    DIV=13,
    BCOR=0mm]{scrlttr2}
\parskip4mm
\parindent0mm
\usepackage[english,ngerman]{babel}
\usepackage[utf8]{inputenc}
\usepackage{csquotes}

\usepackage{lmodern}
\renewcommand*\familydefault{\sfdefault}
\usepackage[T1]{fontenc}

\usepackage{changepage}
\changepage{+3cm}{}{}{}{}{}{}{}{-5cm}
\LoadLetterOption{sender}

\begin{document}
<< range . >>
\newpage
\setkomavar*{enclseparator}{Appendix}
\setkomavar{subject}{Subject: This is an Example}
\setkomavar{date}{<< getDate >>}
\setkomavar{place}{<< getPlace >>}

\begin{letter}{
    << .Name.Title >> << .Name.First >> << .Name.Middle | shortMiddleName >> << .Name.Last >> \\
    << .Address.Street >>\\
    << .Address.Zip >> << .Address.Location >>
}
\opening{Dear Recipient}

\selectlanguage{english}
Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. 

\closing{kind regards}
<< end >>
\end{letter}

\end{document}
`

	return jsonstr, latexstr
}
