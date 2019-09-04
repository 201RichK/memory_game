package variable

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"strconv"
)

var (
	Matrices = [...][6][6]int{
		{
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0},
			{1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0},
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

	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)


func ActiveMatrice(y int)[]string{
	tbl := []string{}
	for i:=0; i<len(Matrices[y]); i++ {
		switch i {
		case 0:
			for x:= 0; x<len(Matrices[y][0]); x++ {
				switch Matrices[y][0][x] {
				case 1:
					l := "c1" + strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 1:
			for x:= 0; x<len(Matrices[y][1]); x++ {
				switch Matrices[y][1][x] {
				case 1:
					l := "c2"+strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 2:
			for x:= 0; x<len(Matrices[y][2]); x++ {
				switch Matrices[y][2][x] {
				case 1:
					l := "c3"+strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 3:
			for x:= 0; x<len(Matrices[y][3]); x++ {
				switch Matrices[y][3][x] {
				case 1:
					l := "c4"+strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 4:
			for x:= 0; x<len(Matrices[y][4]); x++ {
				switch Matrices[y][4][x] {
				case 1:
					l := "c5" + strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 5:
			for x:= 0; x<len(Matrices[y][5]); x++ {
				switch Matrices[y][5][x] {
				case 1:
					l := "c6" + strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		}
	}
	return tbl
}

func VerifyInput(y, x, i int) bool {
	x = x-1
	i = i-1
	if Matrices[y][x][i] == 1 {
		return true
	}
	return false
}
var matrice  []string
func VerifyFinish(x, i int, tbl []string) bool {
	x = x-1
	i = i-1

	c := "c"+strconv.Itoa(x) + strconv.Itoa(i)
	matrice = append(matrice, c)
	fmt.Println(matrice)
	if len(tbl) == len(matrice) {
		matrice = []string{}
		return true
	}

	return false
}

