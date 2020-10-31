package counters

import (
	"bufio"
	"bytes"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return len(p), nil
}
