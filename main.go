package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Coordinate of chess piece
type Coordinate struct {
	row    int
	column int
}

//PieceMoveFunc function to calculate move
type PieceMoveFunc func(int, Coordinate) Coordinate

//PieceMoveMap to map the function
type PieceMoveMap map[string]PieceMoveFunc

//ChessRules Type to read the chess rule json
type ChessRules map[string]Rule

//Rule Type to read the chess rule json
type Rule struct {
	PieceActions interface{}
	Steps        int `json:"steps"`
}

func main() {
	result := getChessMoves("queen", "D5")
	fmt.Println(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getChessMoves(chessPiece string, location string) []string {
	m := make(PieceMoveMap)
	m["HL"] = moveHorizontalLeft
	m["HR"] = moveHorizontalRight
	m["VT"] = moveVerticalTop
	m["VB"] = moveVerticalBottom
	m["UL"] = moveUpLeft
	m["UR"] = moveUpRight
	m["DL"] = moveDownLeft
	m["DR"] = moveDownRight

	var moves []string
	rule := readBundleConfig(chessPiece)
	currentCoordinate := converStringtoCoordinates(location)
	if chessPiece != "horse" {
		for _, action := range rule.PieceActions.([]interface{}) {
			if rule.Steps != 1 {
				for i := 1; i <= rule.Steps; i++ {
					moves = addValidCoordinateAfterMove(moves, currentCoordinate, m[action.(string)], i)
				}
			} else {
				moves = addValidCoordinateAfterMove(moves, currentCoordinate, m[action.(string)], rule.Steps)
			}
		}
	} else {
		for k, v := range rule.PieceActions.(map[string]interface{}) {
			newCoordinate := m[k](rule.Steps, currentCoordinate)
			for _, innerAction := range v.([]interface{}) {
				moves = addValidCoordinateAfterMove(moves, newCoordinate, m[innerAction.(string)], rule.Steps)
			}
		}
	}

	return moves
}

func readBundleConfig(chessPiece string) Rule {
	var chessRules ChessRules
	dat := readJSON("chessRule.json")
	json.Unmarshal(dat, &chessRules)
	return chessRules[chessPiece]
}

func readJSON(file string) []byte {
	jsonFile, err := os.Open("RuleConfig/" + file)
	dat, err := ioutil.ReadAll(jsonFile)
	check(err)
	return dat
}

func getBoardMatrix() [][]string {
	var boardMatrix [][]string
	dat := readJSON("boardMatrix.json")
	json.Unmarshal(dat, &boardMatrix)
	return boardMatrix
}

func converStringtoCoordinates(str string) Coordinate {
	boardMatrix := getBoardMatrix()
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if str == boardMatrix[row][column] {
				return Coordinate{row: row, column: column}
			}
		}
	}
	return Coordinate{row: -1, column: -1}
}

func addValidCoordinateAfterMove(moves []string, c Coordinate, f PieceMoveFunc, steps int) []string {
	coordinateAfterMove := f(steps, c)
	if isCoordinateValid(coordinateAfterMove) {
		str := converCoordinatesToString(coordinateAfterMove)
		moves = append(moves, str)
	}
	return moves
}

func isCoordinateValid(c Coordinate) bool {
	if c.row > -1 && c.row < 8 && c.column > -1 && c.column < 8 {
		return true
	}
	return false
}

func converCoordinatesToString(coordinate Coordinate) string {
	boardMatrix := getBoardMatrix()
	return boardMatrix[coordinate.row][coordinate.column]
}

func moveHorizontalLeft(steps int, currentCoordinate Coordinate) Coordinate {
	c := currentCoordinate.column - steps
	return Coordinate{row: currentCoordinate.row, column: c}
}
func moveHorizontalRight(steps int, currentCoordinate Coordinate) Coordinate {
	c := currentCoordinate.column + steps
	return Coordinate{row: currentCoordinate.row, column: c}
}
func moveVerticalTop(steps int, currentCoordinate Coordinate) Coordinate {
	r := currentCoordinate.row + steps
	return Coordinate{row: r, column: currentCoordinate.column}
}
func moveVerticalBottom(steps int, currentCoordinate Coordinate) Coordinate {
	r := currentCoordinate.row - steps
	return Coordinate{row: r, column: currentCoordinate.column}
}
func moveUpLeft(steps int, currentCoordinate Coordinate) Coordinate {
	c := currentCoordinate.column - steps
	r := currentCoordinate.row + steps
	return Coordinate{row: r, column: c}
}
func moveUpRight(steps int, currentCoordinate Coordinate) Coordinate {
	r := currentCoordinate.row + steps
	c := currentCoordinate.column + steps
	return Coordinate{row: r, column: c}
}
func moveDownLeft(steps int, currentCoordinate Coordinate) Coordinate {
	r := currentCoordinate.row - steps
	c := currentCoordinate.column - steps
	return Coordinate{row: r, column: c}
}
func moveDownRight(steps int, currentCoordinate Coordinate) Coordinate {
	c := currentCoordinate.column + steps
	r := currentCoordinate.row - steps
	return Coordinate{row: r, column: c}
}
