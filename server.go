package main

import (
	"fmt"
	"github.com/201RichK/memory_game/variable"
	"html/template"
	"log"
	"net/http"
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

	fmt.Println("run at http://localhost:9000/memory")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func memory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	_ = variable.TPL.ExecuteTemplate(w, "index.html",  variable.Matrices[0])
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
		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		mymsg := []byte("je suis la")
		if err= conn.WriteMessage(msgType, mymsg); err != nil {
			return
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