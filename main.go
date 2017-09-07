package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/zabawaba99/firego"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	f := firego.New("https://seanmeme-fdcf7.firebaseio.com/memes", nil)
	r := mux.NewRouter()

	r.Methods("POST").Path("/gen/{type}/{text}").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		memeType := vars["type"]
		text := vars["text"]

		var url string
		switch memeType {
		case "winter":
			// Shamefully using the awesome https://github.com/jacebrowning/memegen
			url = "https://memegen.link/winter/brace_yourselves/" + text + ".jpg"
		case "mordor":
			url = "https://memegen.link/mordor/one_does_not_simply/" + text + ".jpg"
		default:
			fmt.Fprint(w, "not found")
			return
		}

		response, err := http.Get(url + "?width=400")
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()

		//open a file for writing
		filename := randStringRunes(5) + ".jpg"
		file, err := os.Create("./static/memes/" + filename)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		_, err = f.Push(filename)
		if err != nil {
			log.Fatal(err)
		}
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	port := os.Getenv("PORT")
	addr := ":" + port
	http.ListenAndServe(addr, n)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
