package gauth_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/onee-only/gauth-go"
)

func TestGAuth(t *testing.T) {

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	redirectURI := os.Getenv("REDIRECT_URI")

	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	client := gauth.NewDefaultClient(gauth.ClientOpts{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
	})

	code, err := client.IssueCode(email, password)
	if err != nil {
		t.Fatal(err)
	}

	access, refresh, err := client.IssueToken(code)
	if err != nil {
		t.Fatal(err)
	}

	userInfo, err := client.GetUserInfo(access)
	if err != nil {
		t.Fatal(err)
	}

	access, _, err = client.ReIssueToken(refresh)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(access)

	newUserInfo, err := client.GetUserInfo(access)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(userInfo, newUserInfo) {
		t.Fatalf("userInfo deosn't match: %v %v", userInfo, newUserInfo)
	}
}
