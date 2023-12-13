package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

const port = ":9090"

func main() {
	fmt.Println("Hello, World!")

	mux := http.NewServeMux()

	jobsHandler := &JobsHandler{
		AllJobs: []Job{
			{Title: "Ninja UX Desginer", Id: "5", Details: "Lorem"},
			{Title: "Ninja Web Developer", Id: "6", Details: "Lorem"},
			{Title: "Ninja Vue Developer", Id: "7", Details: "Lorem"},
		},
	}

	mux.Handle("/", jobsHandler)

	fmt.Println("Server running on port ", port)

	s := http.Server{Addr: port,
		Handler: mux,
	}

	s.ListenAndServe()
}

type JobsHandler struct {
	AllJobs []Job
}

type JobsResponseModel struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	Title   string `json:"title"`
	Id      string `json:"id"`
	Details string `json:"details"`
}

func (h *JobsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/":
		GetJobs(w, r, h)
	case r.Method == http.MethodGet && regexp.MustCompile(`^/\d+$`).MatchString(r.URL.Path):
		GetSingleJob(w, r, h)
	default:
		w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte("Not Found"))
		fmt.Fprint(w, "Not Found")
	}
}

func GetJobs(w http.ResponseWriter, r *http.Request, h *JobsHandler) {

	jsonResponseModel := JobsResponseModel{
		Jobs: h.AllJobs,
	}

	jsonResponse, err := json.Marshal(&jsonResponseModel)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func GetSingleJob(w http.ResponseWriter, r *http.Request, h *JobsHandler) {

	jobID := r.URL.Path[1:]

	if jobID == "" {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Print(jobID)

	var job Job

	for _, j := range h.AllJobs {
		if j.Id == jobID {
			job = j
		}
	}

	jobResponse, err := json.Marshal(&job)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.WriteHeader(http.StatusOK)
	w.Write(jobResponse)

}
