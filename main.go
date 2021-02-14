package main

import (
	"encoding/json"
	"flag"
	f "fmt"
	"io/ioutil"
	"os"
	// github "github.com/google/go-github/v33/github"
	// "golang.org/x/oauth2"
)

/*
	ConfigCommand is struct of config.json's
		kid element, command's children's type
	ConfigCommand consists of same elements
		as config.json's command's children

	ID - "id" - identifier of command
	Command - "command" - command that
				should be executed on github
	Action - "action" - action to do
				when command is executed
*/
type ConfigCommand struct {
	ID      string `json:"id"`
	Command string `json:"command"`
	Action  string `json:"action"`
}

/*
	ConfigJSON isstructure that reflecs
		root body of config.json
	ConfigJSON consists of same elements
		as config.json

	Name - "name" - name of application
	Commands - "commands" - list of all commands
	Comment - "on_comment" - list of commands
					that can be executed by github comments
	PLEequest - "on_pull_request" - list of
					commands that can be executed by github pull requets
*/
type ConfigJSON struct {
	Name      string          `json:"name"`
	Commands  []ConfigCommand `json:"commands"`
	Comment   []string        `json:"on_comment"`
	PLEequest []string        `json:"on_pull_request"`
}

/*
	ParseJson reads JSON by passed structure
		and stores output in passed variable
	After parsing and storing, defer
		opened json file


	Args:
		filename:  (string) name of json file
		store: structure that json will be parsed by and will
			be returned inpassed variabke
*/
func ParseJSON(filename string, store interface{}) {

	configFile, err := os.Open(filename)

	if err != nil {
		f.Println("Could not open config.yml file")
	}

	configJSONBytes, err := ioutil.ReadAll(configFile)
	if err != nil {
		f.Println("could not convert json to byte")
	}

	json.Unmarshal(configJSONBytes, &store)

	defer configFile.Close()
}

func main() {

	var config ConfigJSON
	ParseJSON("config.json", &config)

	var userGit = flag.String("git-user", "levankhelo", "GitHub Token")
	var tokenGit = flag.String("git-token", "", "GitHub's Personal Access Token with admin:repo_hook and repo permissions")
	var webhookGit = flag.String("webhook", "/events", "GitHub webhook tag")
	var localPort = flag.String("port", "4141", "local port for listener")

	// display arg values
	f.Println("----GIT----", "\n|\tUser:", *userGit, "\n|\tToken:", *tokenGit)
	f.Println("--WEBHOOK--", "\n|\tHook:", *webhookGit, "\n|\tPort:", *localPort)
	f.Println()

	f.Println("Hello")

	// TODO: create structure with WEBHOOK request, Request and response (r, w) tags
	// TODO: in WEBHOOK handler, put all requests in queue and execute commands asyncronously
	// TODO: start ngrok server
	// TODO: start server
	// TODO: use GitHub api to log in to github - use token and username
	// TODO: use ?Git? api to create webhook on repo from current application
	// TODO: parse all commands from Github API and put them in queue
	// TODO: remove finished element from queue

}
