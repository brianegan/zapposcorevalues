package main

import(
"fmt"
"log"
"os"
"net/http"
"encoding/json"
"io/ioutil"
"github.com/bmizerany/pat"
"github.com/realchaseadams/CoreValue/bindata"
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
  mux.Get("/CoreValue", http.HandlerFunc(rootHandler))
  mux.Get("/CoreValue/:id", http.HandlerFunc(valueHandler))

  http.Handle("/", mux)

  log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
  fmt.Println("Server started at http://localhost:8080")

}

func loadCoreValues() {
  doc, err := ioutil.ReadFile("./CoreValues.json")

  if err != nil {
    log.Fatal(err)
  }

  err = json.Unmarshal(doc, &values)

  return
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
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