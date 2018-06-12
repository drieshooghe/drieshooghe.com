package main

import (
	"fmt"
	"testing"

	"github.com/Flaque/filet"
)

var g = &github{}

func TestTokenGetter(t *testing.T) {

	// Keys to test for
	keys := map[string]map[string]string{
		"github": {
			"GITHUB_OAUTH_TOKEN": "ABCD",
		},
	}
	var envContent string
	for _, svc := range keys {
		for k, v := range svc {
			line := fmt.Sprintf("%s%s%s%s", k, "=", v, "\n")
			envContent = envContent + line
		}
	}

	// Mock a .env file and assign it an access token
	defer filet.CleanUp(t)
	filet.File(t, ".env.test", envContent)

	// Test github keys fetching
	for name, token := range keys["github"] {
		if token != g.tokens()[name] {
			t.Errorf("Token for %s was incorrect, got: %s, want: %s", name, g.tokens()[name], token)
		}
	}
}
