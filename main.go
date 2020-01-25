package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

// command line default arguments
const (
	defaultAlgorithm = SHA1
	defaultPattern   = "./*"
)

// command line parameters
var (
	usedAlgorithm = defaultAlgorithm
)

func init() {
	registerFlag(&usedAlgorithm, "algorithm", "a", "The algorithm to use for file content reduction.")
}

func registerFlag(value flag.Value, long, short, usage string) {
	flag.Var(value, long, usage)
	flag.Var(value, short, usage+" (shorthand)")
}

func main() {
	exitOnError := func(err error) {
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
	}

	// phase 0: get command line arguments
	flag.Parse()
	patterns := flag.Args()
	if len(patterns) == 0 {
		patterns = append(patterns, defaultPattern)
	}

	// phase 1: get list of files to check
	files, err := findFiles(patterns)
	exitOnError(err)

	// phase 2: compute groups of duplicate files
	groups, err := computeGroups(files)
	exitOnError(err)

	// phase 3: present result to user
	err = showGroups(groups)
	exitOnError(err)
}

func findFiles(patterns []string) (Files, error) {
	// use map keys as set
	fileSet := map[string]struct{}{}
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate pattern: %w", err)
		}
		// add all matches to the set of files
		for _, match := range matches {
			fileSet[match] = struct{}{}
		}
	}
	// copy from set to slice
	files := Files{}
	for file := range fileSet {
		files = append(files, file)
	}
	return files, nil
}

func showGroups(groups FileGroups) error {
	for group, files := range groups {
		fmt.Println(group)
		for _, file := range files {
			fmt.Printf("  %v\n", file)
		}
	}
	return nil
}
