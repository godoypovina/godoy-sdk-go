package godoy_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	godoy "github.com/godoypovina/godoy-sdk-go"
)

const (
	TestEmail    string = "javierf"
	TestPassword string = "Password1"
)

var g godoy.Godoy

func TestMain(m *testing.M) {
	fmt.Println("Godoy SDK GO - Test : Main")

	g = godoy.NewGodoy(TestEmail, TestPassword, true)

	_, err := g.GetAccessToken()
	if err != nil {
		os.Exit(1)
	}

	//Run the other tests
	os.Exit(m.Run())
}

func TestGodoyInstance(t *testing.T) {
	g = godoy.NewGodoy(TestEmail, TestPassword, true)
	if g.Token != "" {
		t.Errorf("Expected AccessToken to be empty and got %s", g.Token)
	}
}

func TestAccessToken(t *testing.T) {
	at, err := g.GetAccessToken()
	if err != nil {
		t.Fatalf("Error requesting an Access Token: %v", err)
	}
	if at == "" {
		t.Errorf("Expected AccessToken to contain a value and is empty")
	}
}

func TestGetAllArticles(t *testing.T) {
	articles, err := g.GetAllArticles(url.Values{
		"page":  {"1"},
		"limit": {"5"},
	})

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	fmt.Printf("%+v", articles)
}
