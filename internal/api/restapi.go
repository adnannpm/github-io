package api

import (
	"encoding/json"
	"fmt"
	"github-io/internal/constant"
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

func formatUserTable(user User) string {
	rows := [][2]string{
		{"FIELD", "VALUE"},
		{"Username", printableValue(user.Login)},
		{"Name", printableValue(user.Name)},
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

func printableValue(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return "-"
	}

	return strings.NewReplacer("\r", " ", "\n", " ", "\t", " ").Replace(value)
}
