package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadAndLogEnvVarsFromWorkflow(t *testing.T) {
	name := os.Getenv("NAME")
	email := os.Getenv("EMAIL")

	if name == "" || email == "" {
		t.Error("env vars for name and email not loaded")
	}

	fmt.Printf("Name : %s\n", name)
	fmt.Printf("Email : %s\n", email)
}
