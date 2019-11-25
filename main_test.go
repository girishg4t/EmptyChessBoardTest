package main

import (
	"testing"
)

//TestreadBundleConfig to test the  config
// func TestReadbundleconfig(t *testing.T) {
// 	d := readBundleConfig("king")
// 	if (d.PieceActions == ([]struct{})) {
// 		t.Errorf("Expected deck length of 16 , but got %v")
// 	}
// }

func TestGetCoordinates(t *testing.T) {
	r, c := converStringtoCoordinates("D5")
	if r == -1 || c == -1 {
		t.Errorf("Row Column")
	}
}
