package sstable

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Compact(dataDir string, outputFile string) error {

	files, err := ListSSTables(dataDir)

	if err != nil {
		return err
	}

	merged := make(map[string]string)

	for i := len(files) - 1; i >= 0; i-- {

		file, err := os.Open(files[i])

		if err != nil {
			return err
		}

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {

			line := scanner.Text()

			parts := strings.Split(line, ",")

			if len(parts) < 2 {
				continue
			}

			key := parts[0]
			value := parts[1]

			merged[key] = value
		}

		file.Close()
	}

	out, err := os.Create(outputFile)

	if err != nil {
		return err
	}

	defer out.Close()

	for key, value := range merged {

		record := fmt.Sprintf("%s,%s\n", key, value)

		_, err := out.WriteString(record)

		if err != nil {
			return err
		}
	}

	return nil
}
