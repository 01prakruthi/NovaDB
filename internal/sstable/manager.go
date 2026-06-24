package sstable

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ListSSTables(dataDir string) ([]string, error) {

	var files []string

	err := filepath.Walk(dataDir,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			if !info.IsDir() &&
				strings.HasPrefix(info.Name(), "sstable_") {

				files = append(files, path)
			}

			return nil
		})

	if err != nil {
		return nil, err
	}

	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	return files, nil
}
func SearchAll(dataDir string, key string) (string, bool, error) {

	files, err := ListSSTables(dataDir)

	if err != nil {
		return "", false, err
	}

	for _, file := range files {

		value, found, err := Search(file, key)

		if err != nil {
			return "", false, err
		}

		if found {
			return value, true, nil
		}
	}

	return "", false, nil
}
