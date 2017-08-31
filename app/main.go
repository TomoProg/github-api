package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	// アクセストークン読み込み
	access_token, err := ioutil.ReadFile("config/access_token")
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}

	// 認証
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: string(access_token)},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Println("%v", err)
		os.Exit(1)
	}

	fmt.Println("%v", repos)
}
