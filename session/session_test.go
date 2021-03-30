package session

import (
	"fmt"
	"testing"
)

func TestGetSessionActive(t *testing.T) {
	ret := GetSessionActive([]byte(testOpenviduJson))
	fmt.Println("Number of Session := ", ret)
	if ret != 54 {
		t.Errorf("Error occured: %d", ret)
	}
}

func TestGetConnectionActive(t *testing.T) {
	ret := GetConnectionActive([]byte(testOpenviduJson))
	fmt.Println("Number of Connection := ", ret)
	if ret != 54 {
		t.Errorf("Error occured: %d", ret)
	}
}

func TestGetConnectionInSession(t *testing.T) {
	ret := GetConnectionInSession([]byte(testOpenviduJson), 0)
	fmt.Println("Number of Connection in Session[0] := ", ret)
	if ret != 1 {
		t.Errorf("Error occured: %d", ret)
	}
}
