package internal

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// SearchResult represents a single search hit.
type SearchResult struct {
	Heading string
	Snippet string
	Rank    float64
}

// BuildIndex creates an FTS5-indexed SQLite database from chunks.
func BuildIndex(dbPath string, chunks []Chunk) error {
	os.Remove(dbPath) // fresh build

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}
	defer db.Close()

	stmts := []string{
		`CREATE VIRTUAL TABLE fpf_fts USING fts5(heading, body, tokenize='porter unicode61')`,
		`CREATE TABLE meta (key TEXT PRIMARY KEY, value TEXT)`,
	}
	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return fmt.Errorf("exec %q: %w", s, err)
		}
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	ins, err := tx.Prepare(`INSERT INTO fpf_fts (heading, body) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer ins.Close()

	for _, c := range chunks {
		if _, err := ins.Exec(c.Heading, c.Body); err != nil {
			return fmt.Errorf("insert chunk %d: %w", c.ID, err)
		}
	}
	return tx.Commit()
}

// SetMeta writes a key-value pair to the meta table.
func SetMeta(dbPath, key, value string) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(`INSERT OR REPLACE INTO meta (key, value) VALUES (?, ?)`, key, value)
	return err
}

// Search queries the FTS5 index and returns matching sections.
func Search(db *sql.DB, query string, limit int) ([]SearchResult, error) {
	if limit <= 0 {
		limit = 10
	}

	// Build FTS5 query: quote each term for prefix matching
	terms := strings.Fields(query)
	var ftsTerms []string
	for _, t := range terms {
		// Escape quotes in term
		t = strings.ReplaceAll(t, `"`, `""`)
		ftsTerms = append(ftsTerms, fmt.Sprintf(`"%s"*`, t))
	}
	ftsQuery := strings.Join(ftsTerms, " OR ")

	rows, err := db.Query(`
		SELECT heading, snippet(fpf_fts, 1, '>>>', '<<<', '...', 64), rank
		FROM fpf_fts
		WHERE fpf_fts MATCH ?
		ORDER BY rank
		LIMIT ?
	`, ftsQuery, limit)
	if err != nil {
		return nil, fmt.Errorf("search query: %w", err)
	}
	defer rows.Close()

	var results []SearchResult
	for rows.Next() {
		var r SearchResult
		if err := rows.Scan(&r.Heading, &r.Snippet, &r.Rank); err != nil {
			return nil, err
		}
		results = append(results, r)
	}
	return results, rows.Err()
}

// GetFullSection returns the complete body of a section by heading.
func GetFullSection(db *sql.DB, heading string) (string, error) {
	var body string
	err := db.QueryRow(`SELECT body FROM fpf_fts WHERE heading = ?`, heading).Scan(&body)
	if err != nil {
		return "", err
	}
	return body, nil
}

// GetMeta reads a value from the meta table.
func GetMeta(db *sql.DB, key string) (string, error) {
	var val string
	err := db.QueryRow(`SELECT value FROM meta WHERE key = ?`, key).Scan(&val)
	if err != nil {
		return "", err
	}
	return val, nil
}
