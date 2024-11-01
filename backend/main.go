package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitHubStats struct {
	Username     string `json:"username"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
	TotalStars   int    `json:"total_stars"`
	Repositories []struct {
		Name  string `json:"name"`
		Stars int    `json:"stars"`
		Forks int    `json:"forks"`
	} `json:"repositories"`
}

func getGitHubStats(username string) (GitHubStats, error) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// }

	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	user, _, err := client.Users.Get(ctx, username)
	if err != nil {
		return GitHubStats{}, err
	}

	repos, _, err := client.Repositories.List(ctx, username, nil)
	if err != nil {
		return GitHubStats{}, err
	}

	stats := GitHubStats{
		Username:  *user.Login,
		Followers: *user.Followers,
		Following: *user.Following,
	}

	for _, repo := range repos {
		stats.TotalStars += *repo.StargazersCount
		stats.Repositories = append(stats.Repositories, struct {
			Name  string `json:"name"`
			Stars int    `json:"stars"`
			Forks int    `json:"forks"`
		}{
			Name:  *repo.Name,
			Stars: *repo.StargazersCount,
			Forks: *repo.ForksCount,
		})
	}

	//@todo: add more options like languages, etc...

	return stats, nil
}

func githubStatsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	stats, err := getGitHubStats(username)
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}

	http.HandleFunc("/api/github-stats", githubStatsHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
