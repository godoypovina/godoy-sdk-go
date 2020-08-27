package godoy_test

import (
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/godoypovina/godoy-sdk-go"
)

const (
	TestEmail    string = ""
	TestPassword string = ""
)

var g *godoy.Godoy

func TestMain(m *testing.M) {
	fmt.Println("Godoy SDK GO - Test : Main")

	g = godoy.NewGodoy(TestEmail, TestPassword, true)

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
	at, err := g.GetAccessToken()
	if err != nil {
		t.Fatalf("Error requesting an Access Token: %v", err)
	}
	if at == "" {
		t.Errorf("Expected AccessToken to contain a value and is empty")
	}
	articles, err := g.GetAllArticles(url.Values{
		"page":  {"1"},
		"limit": {"100"},
	})

	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	fmt.Printf("%+v", articles)
}
