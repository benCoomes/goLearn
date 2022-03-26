package main

import "testing"

func TestShadowing_DeclareAndAssign_InsideBlock(t *testing.T) {
	value := "initialValue"

	{
		t.Logf("inside block, value is %v before", value)                                                     // here, value is the variable outside the block.
		value, someOther := "insideBlock", "foo"                                                              // declare and assign (need new var to use :=)
		t.Logf("inside block, value is %v after. (avoid 'unused' error for someOther: %v)", value, someOther) // now, value is the new variable created inside this block!
	}

	if value != "initialValue" {
		// since new value var was created inside the block, outer value var is not changed
		t.Errorf("Expected for value to be 'initialValue' but it is %v", value)
	}
}

func TestShadowing_DeclareAndAssign_NoBlock(t *testing.T) {
	value := "initialValue"

	//{ no block in this test
	t.Logf("no block, value is %v before", value)
	value, someOther := "newValue", "foo" // declare and assign (need new var to use :=)
	t.Logf("no block, value is %v after. (avoid 'unused' error for someOther: %v)", value, someOther)
	//}

	if value != "newValue" {
		t.Errorf("Expected for value to be 'newValue' but it is %v", value)
	}
}

func TestShadowing_AssignOnly_InBlock(t *testing.T) {
	value := "initialValue"

	{
		t.Logf("inside block, value is %v before", value)
		value = "newValue" // assignment only, here
		t.Logf("inside block, value is %v after", value)
	}

	if value != "newValue" {
		t.Errorf("Expected for value to be 'newValue' but it is %v", value)
	}
}
