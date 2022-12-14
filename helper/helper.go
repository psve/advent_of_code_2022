package helper

import (
	"bufio"
	"os"
)

func ForEachLine(path string, do func(line string) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err := do(scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}
