package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/m0n0x41d/claude-code-fpf/internal"
	_ "modernc.org/sqlite"
)

var version = "dev"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "search":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: fpf-rag search <query> [--limit N] [--full]\n")
			os.Exit(1)
		}
		cmdSearch(os.Args[2:])
	case "section":
		if len(os.Args) < 3 {
			fmt.Fprintf(os.Stderr, "Usage: fpf-rag section <heading>\n")
			os.Exit(1)
		}
		cmdSection(strings.Join(os.Args[2:], " "))
	case "info":
		cmdInfo()
	default:
		// Treat as search query
		cmdSearch(os.Args[1:])
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `fpf-rag — search the FPF (First Principles Framework) specification

Commands:
  search <query> [--limit N] [--full]   Search FPF spec (default: snippets, --full for complete sections)
  section <heading>                       Get full content of a section by heading
  info                                    Show version and FPF commit info

Examples:
  fpf-rag search "WLNK weak link"
  fpf-rag search "ADI cycle" --limit 5
  fpf-rag search "characterization" --full
  fpf-rag section "3.1. WLNK"
  fpf-rag info
`)
}

func openDB() (*sql.DB, func(), error) {
	// Write embedded DB to temp file
	tmpDir, err := os.MkdirTemp("", "fpf-rag-*")
	if err != nil {
		return nil, nil, fmt.Errorf("create temp dir: %w", err)
	}

	dbPath := filepath.Join(tmpDir, "fpf.db")
	if err := os.WriteFile(dbPath, embeddedDB, 0644); err != nil {
		os.RemoveAll(tmpDir)
		return nil, nil, fmt.Errorf("write temp db: %w", err)
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		os.RemoveAll(tmpDir)
		return nil, nil, fmt.Errorf("open db: %w", err)
	}

	cleanup := func() {
		db.Close()
		os.RemoveAll(tmpDir)
	}
	return db, cleanup, nil
}

func cmdSearch(args []string) {
	limit := 10
	full := false
	var queryParts []string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--limit":
			if i+1 < len(args) {
				i++
				n, err := strconv.Atoi(args[i])
				if err == nil {
					limit = n
				}
			}
		case "--full":
			full = true
		default:
			queryParts = append(queryParts, args[i])
		}
	}

	query := strings.Join(queryParts, " ")
	if query == "" {
		fmt.Fprintf(os.Stderr, "Error: empty query\n")
		os.Exit(1)
	}

	db, cleanup, err := openDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	results, err := internal.Search(db, query, limit)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Search error: %v\n", err)
		os.Exit(1)
	}

	if len(results) == 0 {
		fmt.Println("No results found.")
		return
	}

	for i, r := range results {
		fmt.Printf("### %d. %s\n\n", i+1, r.Heading)
		if full {
			body, err := internal.GetFullSection(db, r.Heading)
			if err == nil {
				fmt.Println(body)
			} else {
				fmt.Println(r.Snippet)
			}
		} else {
			fmt.Println(r.Snippet)
		}
		fmt.Println()
	}
}

func cmdSection(heading string) {
	db, cleanup, err := openDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	body, err := internal.GetFullSection(db, heading)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Section not found: %s\n", heading)
		os.Exit(1)
	}

	fmt.Printf("## %s\n\n%s\n", heading, body)
}

func cmdInfo() {
	db, cleanup, err := openDB()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer cleanup()

	fmt.Printf("fpf-rag version: %s\n", version)

	commit, err := internal.GetMeta(db, "fpf_commit")
	if err == nil {
		fmt.Printf("FPF upstream commit: %s\n", commit)
		fmt.Printf("FPF source: https://github.com/ailev/FPF/commit/%s\n", commit)
	}
}
