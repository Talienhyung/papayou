package main

import (
	"log"
	"net/http"
	"text/template"
)

type Joueur struct {
	Distribue              bool
	Name                   string
	Score                  int
	NombreManchePerduSuite int
	NombreManchePerdu      int
	NombreMancheGagne      int
	NombreMancheGagneSuite int
}

type Structure struct {
	List   []Joueur
	Manche int
}

func main() {
	var structu Structure
	structu.Manche = 1
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Home(w, r, structu)
	})

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		Game(w, r, structu)
	})

	http.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, &structu)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.ListenAndServe(":8080", nil)
}

// Home handles HTTP requests for the home page and renders the appropriate HTML templates
func Home(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles(
		"index.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func Game(w http.ResponseWriter, r *http.Request, infos Structure) {
	template, err := template.ParseFiles(
		"init.html",
	)
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, infos)
}

func handler(w http.ResponseWriter, r *http.Request, infos *Structure) {
	var list []string
	list = append(list, r.FormValue("player1"))
	list = append(list, r.FormValue("player2"))
	list = append(list, r.FormValue("player3"))
	list = append(list, r.FormValue("player4"))
	list = append(list, r.FormValue("player5"))
	list = append(list, r.FormValue("player6"))
	list = append(list, r.FormValue("player7"))
	list = append(list, r.FormValue("player8"))

	for _, truc := range list {
		if truc != "" && truc != " " {
			infos.List = append(infos.List, Joueur{false, truc, 0, 0, 0, 0, 0})
		}
	}
	infos.List[0].Distribue = true

	http.Redirect(w, r, "/game", http.StatusSeeOther)
}
