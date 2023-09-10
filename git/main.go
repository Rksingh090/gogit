package git

import (
	"fmt"

	"github.com/google/go-github/v55/github"
)

var GitConf *GitClient

type GitClient struct {
	Client *github.Client
}

func InitClient() {

	t, err := GetToken()
	if err != nil {
		fmt.Println(err.Error())
	}

	gc := github.NewClient(nil).WithAuthToken(t)

	GitConf = &GitClient{
		Client: gc,
	}

}
