package apiserver

import (
	"encoding/json"
	"fmt"
	"hack/model"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	BindAddr string
	Router   *mux.Router
	Model    *model.Model
}

func New(Addr string) *APIServer {
	return &APIServer{
		Addr,
		mux.NewRouter(),
		model.New("host=localhost user=postgres password=postgres dbname=tree sslmode=disable"),
	}
}

func (s *APIServer) Start() error {
	if err := s.Model.Open(); err != nil {
		return err
	}
	fs := http.StripPrefix("/static", http.FileServer(http.Dir("static/")))
	s.Router.Handle("/", fs)
	s.Router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "HomePage")
	})
	s.Router.HandleFunc("/inputdata", inputDataByForms)
	s.Router.HandleFunc("/test", test).Methods("GET")

	s.Router.HandleFunc("/inputdata/save", s.saveArticle)

	go http.ListenAndServe(s.BindAddr, s.Router)
	return nil
}

type v struct {
	L int `json:"l"`
}

func test(w http.ResponseWriter, r *http.Request) {
	t := v{
		L: 100,
	}
	tmpl, _ := template.ParseFiles("test.html")
	tmpl.Execute(w, nil)
	json.NewEncoder(w).Encode(t)

}

func inputDataByForms(w http.ResponseWriter, r *http.Request) {

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, nil)

}

func (s *APIServer) saveArticle(w http.ResponseWriter, r *http.Request) {
	data := model.Line{
		MicrocategoryId: r.FormValue("category"),
		LocationId:      r.FormValue("location"),
		Price:           r.FormValue("price"),
	}
	fmt.Print(r.Body)
	s.Model.Replace(&data)
}
