package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/csv"
	"os"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    // NOTE: Ordering is important here! This has to be defined before
    // the other `/article` endpoint. 
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	
	fileServer := http.FileServer(http.Dir("./static")) // New code
    http.Handle("/", fileServer) // New code
	http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
	}
	
    // Articles = []Article{
        // Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        // Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    // }
    // handleRequests()
}

type Article struct {
    Id      string `json:"Id"`
    Title   string `json:"Title"`
    Desc    string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article


func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful \n")
    DAAL_BATI_COMBO_1 := r.FormValue("qty_DAAL_BATI_COMBO_1")
    DAAL_BATI_COMBO_2 := r.FormValue("qty_DAAL_BATI_COMBO_2")
	
	empData := []string {DAAL_BATI_COMBO_1, DAAL_BATI_COMBO_2}
	
	if err := addcol("employee.csv", empData); err != nil {
        panic(err)
    }

    fmt.Fprintf(w, "DAAL BATI COMBO 1 = %s\n", DAAL_BATI_COMBO_1)
    fmt.Fprintf(w, "DAAL BATI COMBO 2 = %s\n", DAAL_BATI_COMBO_2)
}

func addcol(fname string, empData []string) error {
    // read the file
    f, err := os.Open(fname)
    if err != nil {
        return err
    }
    r := csv.NewReader(f)
    lines, err := r.ReadAll()
    if err != nil {
        return err
    }
    if err = f.Close(); err != nil {
        return err
    }

    // add column
	lines = append(lines, empData)
    //l := len(lines)
    //if len(column) < l {
    //    l = len(column)
    //}
    //for i := 0; i < l; i++ {
    //    lines[i] = append(lines[i], column[i])
    //}

    // write the file
    f, err = os.Create(fname)
    if err != nil {
        return err
    }
    w := csv.NewWriter(f)
    if err = w.WriteAll(lines); err != nil {
        f.Close()
        return err
    }
    return f.Close()
}

