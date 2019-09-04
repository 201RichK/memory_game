package main

import (
	"fmt"
	"github.com/201RichK/memory_game/variable"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	variable.TPL = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.html"))
}

func toView(i int) int {
	return i+1
}

var fm = template.FuncMap{
	"T" : toView,
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))

	http.HandleFunc("/", redirect)
	http.HandleFunc("/memory", memory)
	http.HandleFunc("/echo",  echo)

	fmt.Println("run at http://localhost:3002/memory")
	log.Fatal(http.ListenAndServe(":3002", nil))
}

func memory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	_ = variable.TPL.ExecuteTemplate(w, "index.html",  variable.Matrices[1])
}

func redirect(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		http.Redirect(w, r, "/memory", http.StatusSeeOther)
	}
	return
}

func echo (w http.ResponseWriter, r *http.Request) {
	conn, _ := variable.Upgrader.Upgrade(w, r, nil)
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}


		if _ , err := strconv.Atoi(string(msg)); err == nil {

			x, _ := strconv.Atoi(string(msg[0]))
			t, _ := strconv.Atoi(string(msg[1]))
			if variable.VerifyInput(1, x, t) {
				_ = conn.WriteMessage(msgType, []byte("success"))
				if variable.VerifyFinish(x, t, variable.ActiveMatrice(1) ) {
					fmt.Println("finish")
				}
			} else {
				_ = conn.WriteMessage(msgType, []byte("fail"))
			}
		}

		if string(msg) == "ok" {
			time.AfterFunc(1*time.Second, func() {

				for index := range variable.Matrices{

					msg := variable.ActiveMatrice(index)
					time.Sleep(15*time.Second)
					nmsg := strings.Join(msg, " ")
					if err = conn.WriteMessage(msgType, []byte(nmsg)); err != nil {
						return
					}
					time.AfterFunc(2*time.Second, func() {
						_ = conn.WriteMessage(msgType, []byte("stop"))
					})

				}

			})
		}
	}
}

/**
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	boom := time.After(20 * time.Second)
	i := 20
	for {
		select {
		case <-tick:
			fmt.Println(i)
		case <-boom:
			fmt.Println("FIN !")
			return
		default:
			i--
			//fmt.Println("    .")
			time.Sleep(1 * time.Second)
		}
	}
}

 */