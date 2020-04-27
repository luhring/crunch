package crunch

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func InputFile(path string) (*os.File, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// We're getting input from STDIN

		if path != "" {
			//noinspection GoUnhandledErrorResult
			fmt.Fprintf(os.Stderr, "warning: received data on STDIN, so ignoring argument specifying input file (argument value: %s)", path)
		}

		return os.Stdin, nil
	} else {
		if path != "" {
			absolutePath, err := filepath.Abs(path)
			if err != nil {
				return nil, err
			}

			file, err := os.Open(absolutePath)
			if err != nil {
				return nil, err
			}

			return file, nil
		} else {
			return nil, errors.New("unable to find input file, file path was empty")
		}
	}
}
