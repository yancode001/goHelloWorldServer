package main

import "testing"

func TestGreetingSpecificJohn(t *testing.T) {
	greeting := CreateGreeting("Yan")
	if greeting != "Hello, Yan\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Yan\n")
	}
}

func TestGreetingSpecificDemo(t *testing.T) {
	greeting := CreateGreeting("Demo")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}

func TestShowFailure(t *testing.T) {
	greeting := CreateGreeting("Demo1")
	if greeting != "Hello, Demo\n" {
		t.Errorf("Intentional failure. got: %s, want: %s.", greeting, "Hello, Demo\n")
	}
}



func TestGreetingDefault(t *testing.T) {
	greeting := CreateGreeting("")
	if greeting != "Hello, Guest\n" {
		t.Errorf("Greeting was incorrect, got: %s, want: %s.", greeting, "Hello, Guest\n")
	}
}
 


