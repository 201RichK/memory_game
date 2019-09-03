package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
)

var (
	Matrices = [...][6][6]int{
		{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 1, 0, 0, 0},
			{0, 0, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
		},
		{
			{0, 0, 0, 1, 0, 0},
			{1, 1, 1, 1, 0, 1},
			{0, 0, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 1, 0},
		},
		{
			{0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 1, 1},
			{1, 0, 0, 0, 0, 0},
			{0, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 0},
			{0, 1, 1, 1, 1, 0},
		},
		{
			{0, 0, 1, 1, 0, 1},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{1, 1, 0, 1, 0, 1},
			{0, 0, 0, 1, 0, 0},
			{0, 1, 1, 1, 0, 0},
		},
	}
	TPL *template.Template
)
func init() {
	TPL = template.Must(template.New("").Funcs(fm).ParseFiles("templates/index.html"))
}

func toView(i int) int {
	return i+1
}

func activateMatrice(i int) string {
	if i == 1 {
		return "active"
	}
	return ""
}

var fm = template.FuncMap{
	"T" : toView,
	"Active" : activateMatrice,
}

func main() {
	r := httprouter.New()
	r.ServeFiles("/static/*filepath", http.Dir("./public"))

	r.GET("/memory", memory)

	fmt.Println("run at http://localhost:3000/memory")
	log.Fatal(http.ListenAndServe(":3000", r))
}

func memory(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Header().Set("Content-Type", "text/html")
	_ = TPL.ExecuteTemplate(w, "index.html",  Matrices[0])
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