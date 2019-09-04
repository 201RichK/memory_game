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

	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)


func ActiveMatrice(MatriceId int)[]string{
	tbl := []string{}
	for i:=0; i<len(Matrices[MatriceId]); i++ {
		switch i {
		case 0:
			for x:= 0; x<len(Matrices[MatriceId][0]); x++ {
				switch Matrices[MatriceId][0][x] {
				case 1:
					l := "c1" + strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 1:
			for x:= 0; x<len(Matrices[MatriceId][1]); x++ {
				switch Matrices[MatriceId][1][x] {
				case 1:
					l := "c2"+strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 2:
			for x:= 0; x<len(Matrices[MatriceId][2]); x++ {
				switch Matrices[MatriceId][2][x] {
				case 1:
					l := "c3"+strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 3:
			for x:= 0; x<len(Matrices[MatriceId][3]); x++ {
				switch Matrices[MatriceId][3][x] {
				case 1:
					l := "c4"+strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 4:
			for x:= 0; x<len(Matrices[MatriceId][4]); x++ {
				switch Matrices[MatriceId][4][x] {
				case 1:
					l := "c5" + strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		case 5:
			for x:= 0; x<len(Matrices[MatriceId][5]); x++ {
				switch Matrices[MatriceId][5][x] {
				case 1:
					l := "c6" + strconv.Itoa( x+1)
					tbl = append(tbl, l)
				}
			}
		}
	}
	return tbl
}

func IsSuccess(matriceId, ligneId, colId,  msgType int, conn *websocket.Conn ) bool {
	ligneId = ligneId-1
	colId = colId-1
	if Matrices[matriceId][ligneId][colId] == 1  {
		_ = conn.WriteMessage(msgType, []byte("success"))
		if VerifyFinish(ligneId, colId, ActiveMatrice(matriceId) ){
			fmt.Println("finish")
		}
		return true
	}
	_ = conn.WriteMessage(msgType, []byte("fail"))
	return false
}


var matrice  []string
func VerifyFinish(ligneId, colId int, tbl []string) bool {
	ligneId = ligneId-1
	colId = colId-1

	c := "c"+strconv.Itoa(ligneId) + strconv.Itoa(colId)
	matrice = append(matrice, c)
	fmt.Println(matrice)
	if len(tbl) == len(matrice) {
		matrice = []string{}
		return true
	}

	return false
}