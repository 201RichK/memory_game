package main

import (
	"fmt"
	"github.com/201RichK/memory_game/variable"
	"honnef.co/go/js/dom"
	"net/http"
	"strconv"
	"time"
)

var d = dom.GetWindow().Document()

func new() {
	http.Handle("/static", http.FileServer(http.Dir("./public")))
	time.AfterFunc(2*time.Second, func() {
		for t := range variable.Matrices {
			fmt.Println("t", t)
			activeMatrice(t)
			time.Sleep(3*time.Second)
			deactiveMatrice()

			addClicklistener(t)
			time.Sleep(15*time.Second)
			removeClickListener()
			deactiveMatrice()

			time.Sleep(time.Second)
		}
	})
}


func addClicklistener(mId int){
	for i:= 1; i<37;i++ {
		e1 := d.GetElementByID("c"+strconv.Itoa(i))
		e1.AddEventListener("click", false, func (event dom.Event) {
			event.PreventDefault()

			var index, _ = strconv.Atoi(e1.ID()[1:])
			if verifyUserinput(index, mId) {
				e1.Class().SetString("active")
			}else {
				e1.Class().SetString("error")
			}
		})
	}
}

func removeClickListener(){
	for i:= 1; i<37;i++ {
		e1 := d.GetElementByID("c"+strconv.Itoa(i))
		e1.AddEventListener("click", false, func (event dom.Event) {
			event.PreventDefault()
			e1.Class().SetString("")
			//fmt.Println("clicked", e1.ID())
		})
	}
}

func deactiveMatrice(){
	for i:= 1; i<37;i++ {
		e2 := d.GetElementByID("c"+strconv.Itoa(i))
		if e2.Class().Contains("active") {
			e2.Class().Remove("active")
		}else if e2.Class().Contains("error") {
			e2.Class().Remove("error")
		}
	}
}

func activeMatrice(y int){
	for i:=0; i<len(variable.Matrices[y]); i++ {
		switch i {
		case 0:
			for x:= 0; x<len(variable.Matrices[y][0]); x++ {
				switch variable.Matrices[y][0][x] {
				case 1:
					e1 := d.GetElementByID("c"+strconv.Itoa(x+1))
					e1.Class().SetString("active")
				}
			}
		case 1:
			for x:= 0; x<len(variable.Matrices[y][1]); x++ {
				switch variable.Matrices[y][1][x] {
				case 1:
					e1 := d.GetElementByID("c"+strconv.Itoa(x+7))
					e1.Class().SetString("active")
				}
			}
		case 2:
			for x:= 0; x<len(variable.Matrices[y][2]); x++ {
				switch variable.Matrices[y][2][x] {
				case 1:
					e1 := d.GetElementByID("c"+strconv.Itoa(x+13))
					e1.Class().SetString("active")
				}
			}
		case 3:
			for x:= 0; x<len(variable.Matrices[y][3]); x++ {
				switch variable.Matrices[y][3][x] {
				case 1:
					e1 := d.GetElementByID("c"+strconv.Itoa(x+19))
					e1.Class().SetString("active")
				}
			}
		case 4:
			for x:= 0; x<len(variable.Matrices[y][4]); x++ {
				switch variable.Matrices[y][4][x] {
				case 1:
					e1 := d.GetElementByID("c"+strconv.Itoa(x+25))
					e1.Class().SetString("active")
				}
			}
		case 5:
			for x:= 0; x<len(variable.Matrices[y][5]); x++ {
				switch variable.Matrices[y][5][x] {
				case 1:
					e1 := d.GetElementByID("c"+strconv.Itoa(x+31))
					e1.Class().SetString("active")
				}
			}
		}
	}
}

func idInSlice(item int, t []int) bool {
	for _, b := range t  {
		if b == item {
			return true
		}
	}
	return false
}

func verifyUserinput(id, Mid int) bool {
	l1 := []int{1, 2, 3, 4, 5, 6}
	l2 := []int{7, 8, 9, 10, 11, 12}
	l3 := []int{13, 14, 15, 16, 17, 18}
	l4 := []int{19, 20, 21, 22, 23, 24}
	l5 := []int{25, 26, 27, 28, 29, 30}
	l6 := []int{31, 32, 33, 34, 35, 36}

	if idInSlice(id, l1)  {
		if variable.Matrices[Mid][0][id-1] == 1 {
			return true
		}
	}

	if idInSlice(id, l2) {
		if variable.Matrices[Mid][1][id - 7] == 1 {
			return true
		}
	}

	if idInSlice(id, l3) {
		if variable.Matrices[Mid][2][id - 13] == 1 {
			return true
		}
	}

	if idInSlice(id, l4) {
		if variable.Matrices[Mid][3][id-19] == 1 {
			return true
		}
	}

	if idInSlice(id, l5) {
		if variable.Matrices[Mid][4][id-25] == 1 {
			return true
		}
	}
	if idInSlice(id, l6) {
		if variable.Matrices[Mid][5][id-31] == 1 {
			return true
		}
	}
	return false
}