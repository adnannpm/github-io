package api

import (
	"encoding/json"
	"fmt"
	"github-io/internal/config"
	"github-io/internal/constant"
	"github-io/internal/helper"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

type User struct {
	Login       string `json:"login"`
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

type Repo struct {
	name		string `json:"name"`
	url			string `json:"url"`
}

func Username(name string) {
	endpoint := fmt.Sprintf("%s/users/%s", constant.URL_API_GITHUB, name)
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	fmt.Print(formatUserTable(user))
}

func RepoList(username string) {
	if username != "" {
		endpoint := fmt.Sprintf("%s/users/%s/repos", constant.URL_API_GITHUB, username)
		resp, err := http.Get(endpoint)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()

		var repos []Repo
		err = json.NewDecoder(resp.Body).Decode(&repos)
		if err != nil {
			fmt.Println("Decode error:", err)
			return
		}

		fmt.Print(len(repos))
		return
	}

	endpoint := fmt.Sprintf("%s/user/repos", constant.URL_API_GITHUB)

	token, _ := config.GetToken()
	
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Authorization", "Bearer " + token)
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP error:", err)
		return
	}
	defer resp.Body.Close()

	var repos []Repo
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		fmt.Println("Decode error:", err)
		return
	}

	fmt.Print(len(repos))
}

func formatUserTable(user User) string {
	rows := [][2]string{
		{"FIELD", "VALUE"},
		{"Username", helper.PrintableValue(user.Login)},
		{"Name", helper.PrintableValue(user.Name)},
		{"Public Repos", strconv.Itoa(user.PublicRepos)},
	}

	fieldWidth, valueWidth := 0, 0
	for _, row := range rows {
		fieldWidth = max(fieldWidth, utf8.RuneCountInString(row[0]))
		valueWidth = max(valueWidth, utf8.RuneCountInString(row[1]))
	}

	var table strings.Builder
	writeBorder := func(left, middle, right string) {
		fmt.Fprintf(
			&table,
			"%s%s%s%s%s\n",
			left,
			strings.Repeat("─", fieldWidth+2),
			middle,
			strings.Repeat("─", valueWidth+2),
			right,
		)
	}
	writeRow := func(row [2]string) {
		fmt.Fprintf(&table, "│ %-*s │ %-*s │\n", fieldWidth, row[0], valueWidth, row[1])
	}

	writeBorder("┌", "┬", "┐")
	writeRow(rows[0])
	writeBorder("├", "┼", "┤")
	for _, row := range rows[1:] {
		writeRow(row)
	}
	writeBorder("└", "┴", "┘")

	return table.String()
}


