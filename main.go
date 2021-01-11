package main

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var githubOauthConfig = &oauth2.Config{
	ClientID:     "f24c7a1b39e88457d490",
	ClientSecret: "", //client secret
	Endpoint:     github.Endpoint,
	// don't have to define redirect here as it is already defined during signup process
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/oauth/github", startGithubOauth)
	http.ListenAndServe(":8080", nil)

}

func startGithubOauth(w http.ResponseWriter, r *http.Request) {
	redirectURL := githubOauthConfig.AuthCodeURL("0000")
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<!DOCTYPE html>
	<html lang="en">
	  <head>
		<meta charset="UTF-8" />
		<title>Document</title>
	  </head>
	  <body>
		<form action="/oauth/github" method="POST">
		  <input type="submit" value="Login With Github"/>
		</form>
	  </body>
	</html>
	`)
}
