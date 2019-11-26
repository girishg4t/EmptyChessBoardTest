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
	result := getChessMoves("king", "D5")
	fmt.Println(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getChessMoves(chessPiece string, location string) []string {
	rule := readBundleConfig(chessPiece)
	m := getPieceMoveMapping()
	currentCoordinate := converStringtoCoordinates(location)
	if chessPiece != "horse" {
		return addValidCoordinateAfterMove(currentCoordinate, rule, rule.PieceActions, m)
	}
	for k, v := range rule.PieceActions.(map[string]interface{}) {
		newCoordinate := m[k](rule.Steps, currentCoordinate)
		return addValidCoordinateAfterMove(newCoordinate, rule, v, m)
	}
	return []string{}
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

func getPieceMoveMapping() PieceMoveMap {
	m := make(PieceMoveMap)
	m["HL"] = moveHorizontalLeft
	m["HR"] = moveHorizontalRight
	m["VT"] = moveVerticalTop
	m["VB"] = moveVerticalBottom
	m["UL"] = moveUpLeft
	m["UR"] = moveUpRight
	m["DL"] = moveDownLeft
	m["DR"] = moveDownRight
	return m
}
func addValidCoordinateAfterMove(c Coordinate, rule Rule, actions interface{}, m PieceMoveMap) []string {
	var moves []string
	for _, action := range actions.([]interface{}) {
		for i := 1; i <= rule.Steps; i++ {
			coordinateAfterMove := m[action.(string)](rule.Steps, c)
			if isCoordinateValid(coordinateAfterMove) {
				str := converCoordinatesToString(coordinateAfterMove)
				moves = append(moves, str)
			}
		}
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
