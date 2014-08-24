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

var values CoreValueCollection

func main() {
  loadCoreValues()
  mux := pat.New()
  mux.Get("/", http.HandlerFunc(rootHandler))
  mux.Get("/CoreValue", http.HandlerFunc(allValuesHandler))
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
}

func allValuesHandler(w http.ResponseWriter, r *http.Request) {
  response, _ := json.Marshal(values)
  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func valueHandler(w http.ResponseWriter, r *http.Request) {
  params := r.URL.Query()
  id := params.Get(":id")
  w.Header().Set("Content-Type", "application/json")
  response, _ := json.Marshal(values.CoreValues[id])
  w.Write(response)
}