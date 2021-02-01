package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Job struct {
	ID          string
	Title       string
	Location    string
	Description string
}
type Page struct {
	Title string
	Jobs  []Job
}

var jobs []Job

func (p *Page) save() error {
	filename := p.Title + ".json"
	file, _ := json.MarshalIndent(p.Jobs, "", " ")
	return ioutil.WriteFile(filename, file, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".json"
	file, _ := ioutil.ReadFile(filename)
	body := []Job{}

	err := json.Unmarshal([]byte(file), &body)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Jobs: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/send/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func sendHandler(w http.ResponseWriter, r *http.Request, title string) {

	keyword := r.FormValue("keyword")
	location := r.FormValue("location")
	var body []Job
	for i := 0; i < len(jobs); i++ {
		resK, errK := regexp.Match(strings.ToLower(keyword), []byte(strings.ToLower(jobs[i].Title)))
		resL, errL := regexp.Match(strings.ToLower(location), []byte(strings.ToLower(jobs[i].Location)))
		if errK == nil && errL == nil && resK && resL {
			body = append(body, Job{jobs[i].ID, jobs[i].Title, jobs[i].Location, jobs[i].Description})
		}
	}
	p := &Page{Title: title, Jobs: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(send|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	file, _ := ioutil.ReadFile("Jobs.json")
	mjobs := []Job{}

	err := json.Unmarshal([]byte(file), &mjobs)
	if err != nil {
		println(err)
		return
	}
	jobs = mjobs

	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/send/", makeHandler(sendHandler))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
