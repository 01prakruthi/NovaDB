package wal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WAL struct {
	file *os.File
}

func New(path string) (*WAL, error) {

	file, err := os.OpenFile(
		path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)

	if err != nil {
		return nil, err
	}

	return &WAL{
		file: file,
	}, nil
}

func (w *WAL) Log(operation string, key string, value string) error {

	record := fmt.Sprintf(
		"%s,%s,%s\n",
		operation,
		key,
		value,
	)

	_, err := w.file.WriteString(record)

	return err
}

func (w *WAL) Close() error {
	return w.file.Close()
}

func Recover(path string) (map[string]string, error) {

	data := make(map[string]string)

	file, err := os.Open(path)

	if err != nil {
		return data, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, ",")

		if len(parts) < 3 {
			continue
		}

		operation := parts[0]
		key := parts[1]
		value := parts[2]

		switch operation {

		case "PUT":
			data[key] = value

		case "DELETE":
			delete(data, key)
		}
	}

	return data, nil
}
