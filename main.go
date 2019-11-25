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

//ChessRules Type to read the chess rule json
type ChessRules map[string]Rule

//Rule Type to read the chess rule json
type Rule struct {
	PieceActions []string
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
	var moves []string
	rule := readBundleConfig(chessPiece)
	currentCoordinate := converStringtoCoordinates(location)
	for _, action := range rule.PieceActions {
		coordinateAfterMove := pieceDirection(action, rule.Steps, currentCoordinate)
		str := converCoordinatesToString(coordinateAfterMove)
		moves = append(moves, str)
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

func converCoordinatesToString(coordinate Coordinate) string {
	boardMatrix := getBoardMatrix()
	return boardMatrix[coordinate.row][coordinate.column]
}

func pieceDirection(action string, steps int, currentCoordinate Coordinate) Coordinate {
	switch action {
	case "HL":
		c := currentCoordinate.column - steps
		return Coordinate{row: currentCoordinate.row, column: c}
	case "HR":
		c := currentCoordinate.column + steps
		return Coordinate{row: currentCoordinate.row, column: c}
	case "VT":
		r := currentCoordinate.row + steps
		return Coordinate{row: r, column: currentCoordinate.column}
	case "VB":
		r := currentCoordinate.row - steps
		return Coordinate{row: r, column: currentCoordinate.column}
	case "UL":
		c := currentCoordinate.column - steps
		r := currentCoordinate.row + steps
		return Coordinate{row: r, column: c}
	case "UR":
		r := currentCoordinate.row + steps
		c := currentCoordinate.column + steps
		return Coordinate{row: r, column: c}
	case "DL":
		r := currentCoordinate.row - steps
		c := currentCoordinate.column - steps
		return Coordinate{row: r, column: c}
	case "DR":
		c := currentCoordinate.column + steps
		r := currentCoordinate.row - steps
		return Coordinate{row: r, column: c}
	default:
		return Coordinate{row: -1, column: -1}
	}
}
