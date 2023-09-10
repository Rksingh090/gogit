package git

import (
	"context"

	"github.com/google/go-github/v55/github"
)

func generateNewToken(installationId int64) (error, *github.InstallationToken) {
	t, err := GetToken()
	if err != nil {
		return err, nil
	}

	gc := github.NewClient(nil).WithAuthToken(t)

	token, resp, err := gc.Apps.CreateInstallationToken(context.Background(), installationId, &github.InstallationTokenOptions{})
	if err != nil {
		return err, nil
	}

	defer resp.Body.Close()

	return nil, token
}
