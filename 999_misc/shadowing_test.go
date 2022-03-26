package main

import "testing"

func TestShadowing_DeclareAndAssign_InsideBlock(t *testing.T) {
	myMap := make(map[string]string)
	myMap["key"] = "mapValue"

	var value string
	value = "initialValue"

	{
		t.Logf("inside block, value is %v before lookup", value)
		value, found := myMap["key"] // declare and assign (need new var 'found' to use :=)
		t.Logf("inside block, value is %v and found is %v", value, found)
	}

	if value != "mapValue" { // This will be true. Fail on purpose so that logs always show.
		// value inside of the block is different than value outside the block!
		t.Errorf("Expected for value to be 'mapValue' but it is %v", value)
	}
}

func TestShadowing_DeclareAndAssign_NoBlock(t *testing.T) {
	myMap := make(map[string]string)
	myMap["key"] = "mapValue"

	var value string
	value = "initialValue"

	//{ no block in this test
	value, found := myMap["key"] // declare and assign (need new var 'found' to use :=)
	t.Logf("no block, value is %v and found is %v right after lookup", value, found)
	//}

	if value != "mapValue" {
		t.Errorf("Expected for value to be 'mapValue' but it is %v", value)
	}
}

func TestShadowing_AssignOnly_InBlock(t *testing.T) {
	myMap := make(map[string]string)
	myMap["key"] = "mapValue"

	var value string
	value = "initialValue"

	{
		value = myMap["key"] // assignment only, here
		t.Logf("inside block, value is %v", value)
	}

	if value != "mapValue" {
		t.Errorf("Expected for value to be 'mapValue' but it is %v", value)
	}
}
