package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	flag "github.com/spf13/pflag"

	"github.com/mrxinu/gosolar"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
)

const version = "0.2.0"

func main() {
	printVersion := flag.BoolP("version", "v", false, "print application version")
	hostname := flag.StringP("hostname", "h", "localhost", "SolarWinds hostname")
	username := flag.StringP("username", "u", "admin", "SolarWinds username")
	password := flag.StringP("password", "p", "", "SolarWinds password")
	search := flag.StringP("search", "s", "", "search string")
	verbs := flag.Bool("verbs", false, "include verbs")
	flag.Parse()

	if *printVersion {
		fmt.Fprintln(os.Stderr, "version: "+version)
		os.Exit(1)
	}

	// check hostname and username; password could be blank
	if *hostname == "" || *username == "" {
		fmt.Fprintln(os.Stderr, "Username (-u) and password (-p) are required.")
		os.Exit(1)
	}

	if *search == "" {
		fmt.Fprintln(os.Stderr, "You must provide a search string with the -s option.")
		os.Exit(1)
	}

	client := gosolar.NewClient(*hostname, *username, *password, true)

	if err := printMeta(client, *search); err != nil {
		log.Fatal(err)
	}

	if *verbs {
		if err := printVerbs(client, *search); err != nil {
			log.Fatal(err)
		}
	}
}

func printMeta(client *gosolar.Client, search string) error {
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
		"pattern": "%" + search + "%",
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
		return errors.Wrap(err, "failed to query")
	}

	if err := json.Unmarshal(res, &rows); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
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

	return nil
}

func printVerbs(client *gosolar.Client, search string) error {
	query := `
		SELECT
			EntityName
			,MethodName
		FROM Metadata.Verb
		WHERE EntityName LIKE @pattern
		AND ISNULL(MethodName, '') != ''
	`

	params := map[string]string{
		"pattern": "%" + search + "%",
	}

	var rows []struct {
		EntityName string `json:"EntityName,omitempty"`
		MethodName string `json:"MethodName,omitempty"`
	}

	res, err := client.Query(query, params)
	if err != nil {
		return errors.Wrap(err, "failed to query")
	}

	if err := json.Unmarshal(res, &rows); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"EntityName", "MethodName"})
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	for _, r := range rows {
		tableRow := []string{
			r.EntityName,
			r.MethodName,
		}

		table.Append(tableRow)
	}

	table.Render()

	return nil
}
