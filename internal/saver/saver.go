package saver

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// SaveM3U saves the links in the specified format.
func SaveM3U(links []string, format string) error {
	switch format {
	case "m3u":
		return saveAsM3U(links)
	case "json":
		return saveAsJSON(links)
	case "csv":
		return saveAsCSV(links)
	default:
		return fmt.Errorf("unsupported format: %s. Use 'm3u', 'json', or 'csv'", format)
	}
}

func saveAsM3U(links []string) error {
	file, err := os.Create("links.m3u")
	if err != nil {
		return err
	}
	defer file.Close()

	for _, link := range links {
		file.WriteString(fmt.Sprintf("%s\n", link))
	}

	return nil
}

func saveAsJSON(links []string) error {
	file, err := os.Create("links.json")
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(links)
	if err != nil {
		return err
	}

	file.Write(jsonData)
	return nil
}

func saveAsCSV(links []string) error {
	file, err := os.Create("links.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, link := range links {
		writer.Write([]string{link})
	}

	return nil
}