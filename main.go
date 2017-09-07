package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/zabawaba99/firego"
)

func main() {
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

		response, err := http.Get(url + "?width=250")
		if err != nil {
			log.Print(err)
			return
		}
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Print(err)
			return
		}

		img := base64.StdEncoding.EncodeToString(data)
		_, err = f.Push(img)
		if err != nil {
			log.Print(err)
			return
		}
	})

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(r)

	port := os.Getenv("PORT")
	addr := ":" + port
	http.ListenAndServe(addr, n)
}
