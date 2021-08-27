package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"music/music"
	"music/storage"
)

type ServerDependencies struct {
	MusicService music.Service
}

func Start(deps ServerDependencies) error {
	server := http.Server{Addr: "localhost:8080"}

	server.Handler = router(deps.MusicService)
	fmt.Println("Starting server on 8080")
	return server.ListenAndServe()
}

func router(musicService music.Service) http.HandlerFunc {
	getTrackHandler := getTrack(musicService)
	postTrackHandler := postTrack(musicService)

	return func(w http.ResponseWriter, r *http.Request) {
		pathWord := strings.Split(r.URL.Path, "/")[1]
		switch pathWord {
		case "tracks":
			switch r.Method {
			case http.MethodGet:
				getTrackHandler(w, r)
			case http.MethodPost:
				postTrackHandler(w, r)
			default:
				notFoundHandler(w, r)
			}
		default:
			notFoundHandler(w, r)
		}
	}
}

func getTrack(service music.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		trackId := struct {
			Id int `json:"id"`
		}{}
		err := decoder.Decode(&trackId)
		if err != nil {
			errStatusCode(w, http.StatusBadRequest)
			return
		}

		track, err := service.GetTrack(r.Context(), trackId.Id)

		if err != nil {
			if errors.Is(err, storage.ErrorNotFound{}) {
				errStatusCode(w, http.StatusNotFound)
				return
			}
			errStatusCode(w, http.StatusInternalServerError)
			return
		}

		respondJson(w, track)
	}
}

func postTrack(service music.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		requestTrack := music.Track{}
		err := decoder.Decode(&requestTrack)
		if err != nil {
			errStatusCode(w, http.StatusBadRequest)
			return
		}

		addedTrack, err := service.AddTrack(r.Context(), requestTrack)
		if err != nil {
			errStatusCode(w, http.StatusInternalServerError)
		}
		respondJson(w, addedTrack)
	}
}

func respondJson(w http.ResponseWriter, jsonStruct interface{}) {
	encoder := json.NewEncoder(w)

	err := encoder.Encode(jsonStruct)
	if err != nil {
		errStatusCode(w, http.StatusInternalServerError)
		return
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	errStatusCode(w, http.StatusNotFound)
}

func errStatusCode(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	_, _ = w.Write([]byte(http.StatusText(code)))
}
