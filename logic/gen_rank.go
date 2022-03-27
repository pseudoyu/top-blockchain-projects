package logic

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"
	"top-blockchain-projects/models"
)

const (
	HEADER = `# Top Blockchain Projects

Top github blockchain projects by number of stars.

| Project Name | Stars | Forks | Open Issues | Description | Last Commit |
| ------------ | ----- | ----- | ----------- | ----------- | ----------- |
`
	FOOTER = "\n*Last Update Time: %v*"
)

func GenerateRank(repos []models.Repo) {

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Stars > repos[j].Stars
	})

	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(readme *os.File) {
		err := readme.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(readme)

	_, err = readme.WriteString(HEADER)
	if err != nil {
		return
	}

	for _, repo := range repos {
		_, err := readme.WriteString(fmt.Sprintf("| [%s](%s) | %d | %d | %d | %s | %v |\n", repo.Name, repo.URL, repo.Stars, repo.Forks, repo.Issues, repo.Description, repo.LastCommitDate.Format("2006-01-02 15:04:05")))
		if err != nil {
			return
		}
	}

	_, err = readme.WriteString(fmt.Sprintf(FOOTER, time.Now().Format(time.RFC3339)))
	if err != nil {
		return
	}
}
