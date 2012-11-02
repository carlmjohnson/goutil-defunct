package file

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Lines(filename string) (<-chan string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	rdr := bufio.NewReader(file)

	ch := make(chan string)

	go readlines(ch, rdr)

	return ch, nil
}

func readlines(ch chan<- string, rdr *bufio.Reader) {
	for {
		switch line, err := rdr.ReadString('\n'); err {

		// If the read succeeded
		case nil:
			ch <- line

		// The `EOF` error is expected when we reach the
		// end of input, so exit gracefully in that case.
		case io.EOF:
			close(ch)
			return

		// Otherwise there's a problem; print the
		// error and exit with non-zero status.
		default:
			log.Fatal(err)
		}
	}
}
