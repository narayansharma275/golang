package main

import (
	"GOMONGO/config"
	"GOMONGO/dao"
	"GOMONGO/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func AllMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	da := dao.MoviesDAO{}
	movies, err := da.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func CreateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	da := dao.MoviesDAO{}

	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := da.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
}

func FindMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	da := dao.MoviesDAO{}
	movie, err := da.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, movie)
}

func UpdateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	da := dao.MoviesDAO{}
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := da.Update(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	da := dao.MoviesDAO{}
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := da.Delete(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	cfg := config.Config{}
	dao := dao.MoviesDAO{}
	cfg.Read()
	dao.Server = cfg.Server
	dao.Database = cfg.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", homePage)
	r.HandleFunc("/movies", AllMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies/{id}", FindMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies", UpdateMovieEndPoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndPoint).Methods("DELETE")
	r.HandleFunc("/movies", CreateMovieEndPoint).Methods("POST")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err.Error())
	}
}
