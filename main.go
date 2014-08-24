package main

import(
"fmt"
"log"
"net/http"
"encoding/json"
// "io/ioutil"
"os"
"github.com/bmizerany/pat"
"github.com/realchaseadams/corevalue/data"
)

type CoreValueCollection struct {
  JobsUrl string `json:"jobsUrl"`
  CoreValues map[string]CoreValue `json:"values"`
}

type CoreValue struct {
  Name string `json:"name"`
  Id string `json:"id"`
  Description string `json:"description"`
  Summary string `json:"summary"`
}

type Response struct {
  StatusCode int `json:"statusCode"`
  Message string `json:"message"`
}

var values CoreValueCollection

func main() {
  loadCoreValues()
  mux := pat.New()
  mux.Get("/", http.HandlerFunc(rootHandler))
  mux.Get("/CoreValue", http.HandlerFunc(allValuesHandler))
  mux.Get("/CoreValue/", http.HandlerFunc(allValuesHandler))
  mux.Get("/CoreValue/:id", http.HandlerFunc(valueHandler))

  http.Handle("/", mux)

  fmt.Println("Server started at http://localhost:" + os.Getenv("PORT"))
  log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), nil))

}

func loadCoreValues() {
  doc, err := data.Asset("CoreValues.json")

  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(doc, &values)

  return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  healthcheck := Response{200, "OKAY"}

  response, err := json.Marshal(healthcheck)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func allValuesHandler(w http.ResponseWriter, r *http.Request) {
  response, err := json.Marshal(values)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func valueHandler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  id := params.Get(":id")
  response, err := json.Marshal(values.CoreValues[id])

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }

  if values.CoreValues[id].Id == "" {
    errorHandler(w, r, http.StatusNotFound)
    return
  }


  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
  w.WriteHeader(status)
  if status == http.StatusNotFound {
    errorMessage := &Response{404, "For non core values, please refer to our non-core value API, coming soon."}
    response, err := json.Marshal(errorMessage)

    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    w.Write(response)
  }
}