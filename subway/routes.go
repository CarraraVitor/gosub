package subway

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const TEMPLATES_PATH = "subway/templates"

type SubData struct {
	Nodes    []Node
	Lanes    []Lane
	LanesMap map[int]Lane
	Paths    [][]Node
	Add      func(int, int) int
	Parsed   []*Step
}

type ContextData struct {
	Location string
	Theme    string
	Subway   SubData
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		fmt.Sprintf("%s/main.html", TEMPLATES_PATH),
	))

    q := r.URL.Query()
    theme := q.Get("theme")
    if theme == "" || (theme != "light" && theme != "dark") {
        theme = "light"
    }

	lanes := ListAllLanes()
	sd := SubData{
		Lanes: lanes,
	}
	ctx := ContextData{
		Subway: sd,
        Theme: theme,
	}

    w.Header().Add("HX-Refresh", "true")

	err := tmpl.Execute(w, ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ListNodes(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	tmpl := template.Must(template.ParseFiles(
		fmt.Sprintf("%s/stationsdropdown.html", TEMPLATES_PATH),
	))

	params := r.URL.Query()
	l, err := strconv.Atoi(params.Get("SelectedLane"))
	if err != nil {
		log.Fatalf("[ERROR]Failed reading lane id: %s", err.Error())
	}
	nodes := FindNodesByLane(l)
	s := SubData{
		Nodes: nodes,
	}
	ctx := ContextData{
		Subway: s,
	}


	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Path(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO] %s %s \n", r.Method, r.URL.Path)
	err := r.ParseForm()
	if err != nil {
		log.Fatal("[ERROR] Failed at Parsing Form")
	}

	src, err := strconv.Atoi(r.Form.Get("StationFrom"))
	if err != nil {
		tmpl := template.Must(template.ParseFiles(
			fmt.Sprintf("%s/400.html", TEMPLATES_PATH),
		))
		ErrorContext := "Preencha o campo de partida"
		err = tmpl.Execute(w, ErrorContext)
		return
	}
	dest, err := strconv.Atoi(r.Form.Get("StationTo"))
	if err != nil {
		tmpl := template.Must(template.ParseFiles(
			fmt.Sprintf("%s/400.html", TEMPLATES_PATH),
		))
		ErrorContext := "Preencha o campo de destino"
		err = tmpl.Execute(w, ErrorContext)
		return
	}

	lanes := MapAllLanes()
	paths := FindPaths(src, dest)
	if len(paths) == 0 {
		fmt.Fprintf(w, "no paths found :(")
		return
	}
	parsed := ParsePath(paths[0])

	ctx := SubData{
		LanesMap: lanes,
		Parsed:   parsed,
		Add:      func(a int, b int) int { return a + b },
	}
	tmpl := template.Must(template.ParseFiles(
		fmt.Sprintf("%s/path.html", TEMPLATES_PATH),
	))
	err = tmpl.Execute(w, ctx)
	if err != nil {
		log.Printf("Failed executing template: \n%s\n", err.Error())
	}
}
