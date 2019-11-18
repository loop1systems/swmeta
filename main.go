package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/mrxinu/gosolar"
	"github.com/olekukonko/tablewriter"
)

const version = "0.1.0"

func main() {
	search := flag.StringP("search", "s", "", "search string")
	printVersion := flag.BoolP("version", "v", false, "print application version")
	flag.Parse()

	if *printVersion {
		fmt.Fprintln(os.Stderr, "version: "+version)
		os.Exit(1)
	}

	if *search == "" {
		fmt.Fprintln(os.Stderr, "You must provide a search string with the -s option.")
		os.Exit(1)
	}

	hostname := "192.168.21.58"
	username := "apiuser"
	password := "apiuser151515"

	client := gosolar.NewClient(hostname, username, password, true)

	query := `
		SELECT
			FullName
			,CanCreate
			,CanRead
			,CanUpdate
			,CanDelete
			,CanInvoke
		FROM Metadata.Entity
		WHERE FullName LIKE @pattern
	`

	params := map[string]string{
		"pattern": "%" + *search + "%",
	}

	var rows []struct {
		FullName  string `json:"FullName,omitempty"`
		CanCreate bool   `json:"CanCreate,omitempty"`
		CanRead   bool   `json:"CanRead,omitempty"`
		CanUpdate bool   `json:"CanUpdate,omitempty"`
		CanDelete bool   `json:"CanDelete,omitempty"`
		CanInvoke bool   `json:"CanInvoke,omitempty"`
	}

	res, err := client.Query(query, params)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(res, &rows); err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "CanCreate", "CanRead", "CanUpdate", "CanDelete", "CanInvoke"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	for _, r := range rows {
		tableRow := []string{
			r.FullName,
			fmt.Sprintf("%t", r.CanCreate),
			fmt.Sprintf("%t", r.CanRead),
			fmt.Sprintf("%t", r.CanUpdate),
			fmt.Sprintf("%t", r.CanDelete),
			fmt.Sprintf("%t", r.CanInvoke),
		}

		table.Append(tableRow)
	}

	table.Render()
}
