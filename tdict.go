package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "reflect"
    "strings"
)

func handleError(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

type Dictionary struct {
    Entries map[string]json.RawMessage
}

type Entry struct {
    Antonyms []string `json:"ANTONYMS"`
    Synonyms []string `json:"SYNONYMS"`
    Meanings map[string]interface{} `json:"MEANINGS"`
}

func meaningPrinter(t interface{}) {
    s := reflect.ValueOf(t)
    fmt.Println(s.Index(0), "|", s.Index(1))
}

func listPrinter(class string, l []string) {
    if len(l) == 0 {
        return
    }
    fmt.Printf("%s: ", class)
    for i, word := range l {
        if (i + 1) == len(l) {
            fmt.Printf("%s\n", word)
        } else {
            fmt.Printf("%s, ", word)
        }
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage:", os.Args[0], "<word>")
        os.Exit(1)
    }
    arg := strings.ToUpper(os.Args[1])
    datadir := "/usr/local/share/tdict/"
    raw, err := ioutil.ReadFile(datadir + "D" + string(arg[0]) + ".json")
    handleError(err)
    var d map[string]*json.RawMessage
    json.Unmarshal([]byte(raw), &d)

    if _, ok := d[arg]; ok {
        var e Entry
        json.Unmarshal(*d[arg], &e)

        for _, v := range e.Meanings {
            meaningPrinter(v)
        }
        listPrinter("SYNONYMS", e.Synonyms)
        listPrinter("ANTONYMS", e.Antonyms)
    } else {
        fmt.Println(os.Args[1], "not found")
    }
}
