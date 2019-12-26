package main


import (
	"testing"
	// "fmt"
)


func TestContainsCapital(t *testing.T){
	str1:="aoeurperueBc"
	str2:="bc"

	if Contains(str1,str2) == true {
		t.Errorf("error")
	}
}


func TestContainsBeginning(t *testing.T){
	str1:="aboeurperue"
	str2:="ab"

	if Contains(str1,str2) == false {
		t.Errorf("error")
	}
}

func TestContainsEnding(t *testing.T){
	str1:="aoeurperuebc"
	str2:="bc"

	if Contains(str1,str2) == false {
		t.Errorf("error")
	}
}

func TestContainsNotFound(t *testing.T){
	str1:="aoeurperuedf"
	str2:="bc"

	if Contains(str1,str2) == true {
		t.Errorf("error")
	}
}

func TestContainsinSymbol(t *testing.T){
	str1:="aoeur bc perued f"
	str2:="bc"

	if Contains(str1,str2) == false {
		t.Errorf("error")
	}
}