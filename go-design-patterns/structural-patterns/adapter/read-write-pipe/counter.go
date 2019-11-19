package main

import (
	"io"
	"os"
	"strconv"
)

type Counter struct {
	Writer io.Writer
}

func (c *Counter) Count(n uint64) uint64 {
	if n == 0 {
		//println(strconv.Itoa(0))
		c.Writer.Write([]byte(strconv.Itoa(0) + "\n"))
		return 0
	}

	cur := n
	//println(strconv.FormatUint(cur, 10))
	c.Writer.Write([]byte(strconv.FormatUint(cur, 10) + "\n"))
	return c.Count(n - 1)
}

func main() {

	pipeReader, pipeWriter := io.Pipe()
	defer pipeReader.Close()
	defer pipeWriter.Close()

	c := Counter{Writer: pipeWriter}

	file, err := os.Create("test.txt")
	if err != nil {
		os.Exit(1)
	}

	// The io.TeeReader function helps us to copy the stream of data
	// from a Reader interface to the Writer interface and,
	// it returns a new Reader that you can still use to stream data again to a second writer
	tee := io.TeeReader(pipeReader, file)

	go func() {
		// The io.Copy adapter can be used like TeeReader
		// it takes a reader and writes its contents to a writer
		io.Copy(os.Stdout, tee)
	}()

	c.Count(5)

}