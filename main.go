package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"top-blockchain-projects/logic"
	"top-blockchain-projects/models"
)

var (
	repos []models.Repo
)

func main() {
	accessToken := logic.GetAccessToken()
	byteContents, err := ioutil.ReadFile("projects.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(byteContents), "\n")
	for _, url := range lines {
		if strings.HasPrefix(url, "https://github.com/") {
			var repo models.Repo
			var commit models.HeadCommit

			repoAPI := fmt.Sprintf("https://api.github.com/repos/%s", strings.TrimFunc(url[19:], logic.TrimSpaceAndSlash))

			resp, err := logic.GetResponse(repoAPI, accessToken)

			if err != nil {
				log.Fatal(err)
			}
			if resp.StatusCode != 200 {
				log.Fatal(resp.Status)
			}

			decoder := json.NewDecoder(resp.Body)
			if err = decoder.Decode(&repo); err != nil {
				log.Fatal(err)
			}

			commitAPI := fmt.Sprintf("https://api.github.com/repos/%s/commits/%s", strings.TrimFunc(url[19:], logic.TrimSpaceAndSlash), repo.DefaultBranch)
			resp, err = logic.GetResponse(commitAPI, accessToken)

			if err != nil {
				log.Fatal(err)
			}
			if resp.StatusCode != 200 {
				log.Fatal(resp.Status)
			}

			decoder = json.NewDecoder(resp.Body)
			if err = decoder.Decode(&commit); err != nil {
				log.Fatal(err)
			}

			repo.LastCommitDate = commit.Commit.Committer.Date
			repos = append(repos, repo)

			fmt.Printf("Repository: %v\n", repo)
			fmt.Printf("Head Commit: %v\n", commit)
		}
	}

	logic.GenerateRank(repos)
}
