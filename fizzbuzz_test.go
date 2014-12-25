package main

import (
	"testing"
)

func TestInput1ShouldBe1(t *testing.T){
	output := FizzBuzz(1)
	if output != "1"{
		t.Error("Expected : 1   Actual : ",output);
	}
}

func TestInput2ShouldBe2(t *testing.T){
	output := FizzBuzz(2)
	if output != "2" {
		t.Error("Expected : 2  Actual : ",output);
	}
}

func TestInput3ShouldBeFizz(t *testing.T){
	output := FizzBuzz(3)
	if output != "Fizz" {
		t.Error("Expected : Fizz  Actual : ",output)
	}
}

func TestInput5ShouldBeBuzz(t *testing.T){
	output := FizzBuzz(5)
	if output != "Buzz" {
		t.Error("Expected : Buzz  Actual : ",output)
	}
}

func TestInput6ShouldBeFizz(t *testing.T){
	output := FizzBuzz(6)
	if output != "Fizz" {
		t.Error("Expected : Fizz  Actual : ",output)
	}
}
 
func TestInput10ShouldBeBuzz(t *testing.T){
	output := FizzBuzz(10)
	if output != "Buzz" {
		t.Error("Expected : Buzz  Actual : ",output)
	}
}

func TestInput15ShouldBeFizzBuzz(t *testing.T){
	output := FizzBuzz(15)
	if output != "FizzBuzz" {
		t.Error("Expected : FizzBuzz  Actual : ",output)
	}
}
