package main

import (
	"fmt"
	"io"
	"os"
)

// Group identifies a group of equal files.
type Group struct {
	size FileSize // in bytes
	hash string   // optional
}

func (g Group) String() string {
	s := g.size.String()
	if len(g.hash) != 0 {
		s += " (" + g.hash + ")"
	}
	return s
}

// Files is a list of file names.
type Files []string

// FileGroups maps equal files into a group.
type FileGroups map[Group]Files

func computeGroups(files Files) (FileGroups, error) {
	groups, err := groupBySize(files)
	if err != nil {
		return nil, err
	}
	err = splitGroupsByHash(groups)
	if err != nil {
		return nil, err
	}
	return groups, err
}

func groupBySize(files Files) (FileGroups, error) {
	groups := FileGroups{}
	for _, file := range files {
		fi, err := os.Stat(file)
		if err != nil {
			return nil, fmt.Errorf("failed to get file size: %w", err)
		}
		if fi.IsDir() {
			continue
		}
		group := Group{size: FileSize(fi.Size())}
		groups[group] = append(groups[group], file)
	}
	return groups, nil
}

func splitGroupsByHash(groups FileGroups) error {
	for group, files := range groups {
		if len(files) <= 1 {
			continue
		}
		// this group contains more than one entry

		// delete old group
		delete(groups, group)

		// calculate hash for each file in group
		for _, file := range files {
			hash, err := computeHash(file)
			if err != nil {
				return err
			}
			// the new group has the same size but with a hash
			newGroup := Group{group.size, hash}
			// add file to groups map
			groups[newGroup] = append(groups[newGroup], file)
		}
	}
	return nil
}

func computeHash(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hash := usedAlgorithm.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", fmt.Errorf("failed to compute hash: %w", err)
	}

	digest := hash.Sum(nil)
	digestStr := fmt.Sprintf("%x", digest)
	return digestStr, nil
}
