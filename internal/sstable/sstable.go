package sstable

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Flush(data map[string]string, filename string) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	keys := make([]string, 0, len(data))

	for k := range data {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {

		record := fmt.Sprintf(
			"%s,%s\n",
			k,
			data[k],
		)

		_, err := file.WriteString(record)

		if err != nil {
			return err
		}
	}

	return nil
}

func Search(filename string, targetKey string) (string, bool, error) {

	file, err := os.Open(filename)

	if err != nil {
		return "", false, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, ",")

		if len(parts) < 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		if key == targetKey {
			return value, true, nil
		}
	}

	return "", false, nil
}
