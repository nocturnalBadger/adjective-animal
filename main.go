package main

import (
    "encoding/json"
    "io/ioutil"
    "math/rand"
    "time"
    "fmt"
    "net/http"
    "strings"
    "log"
)

var adjatives []string
var animals   []string

func main() {
    rand.Seed(time.Now().Unix())

    http.HandleFunc("/", GetAdjativeAnimal)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func GetAdjativeAnimal(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s %s\n", GetAdjative(), GetAnimal())
}

func loadAdjatives() {
    data, err := ioutil.ReadFile("adjatives.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(data, &adjatives)
    if err != nil {
        panic(err)
    }
}

func loadAnimals() {
    data, err := ioutil.ReadFile("animals.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(data, &animals)
    if err != nil {
        panic(err)
    }
}

func GetAdjative() string {
    if len(adjatives) == 0 {
        loadAdjatives()
    }

    return strings.ToLower(adjatives[rand.Intn(len(adjatives))])
}

func GetAnimal() string {
    if len(animals) == 0 {
        loadAnimals()
    }

    return strings.ToLower(animals[rand.Intn(len(animals))])
}
