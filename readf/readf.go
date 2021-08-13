package readf

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// each write/read is system call, make as little as possible such system calls - use buffers

// io. - Reader, Writer (no buffers around - not always efficient, system calls on each Read(), Write())
// bufio. - wrap existing Reader, Writer with Buffer - makes less system calls, memory and cpu efficient
// ioutil - for small and simple task, not efficient at all for long-lived apps and big files
// bytes. - create new buffered Reader, Writer, ....

// readers
// every Read() is system call on disk
// read in memory as little as possible, otherwise system might kill our app if we use too much memory
// use buffers, buffers allow us to read chunks of data only
//BenchmarkReadAll
//BenchmarkReadAll-8     	     153	   8275323 ns/op
//BenchmarkReadChunk
//BenchmarkReadChunk-8   	       3	 485829355 ns/op
//BenchmarkReadBuff
//BenchmarkReadBuff-8    	     226	   5287816 ns/op
//BenchmarkReadBuffR
//BenchmarkReadBuffR-8   	     127	   9181988 ns/op

// writers
// every Write() is system call on disk
// use buffers, buffers allow us to make less write system calls
// system call write is called only when buffer is full
// buff := bytes.NewBuffer(make([]byte, 16))
// buff.WriteTo(io.Writer)

// preferred writer usage
//var w io.WriteCloser
//// initialise writer
//defer w.Close()
//b := bufio.NewWriter(w)
//defer b.Flush() // ensure buffer leftover is written (flushed) as well
//// write operations

func ReadAll() {
	f, err := os.Open("../fff/mf.txt")
	if err != nil {
		logrus.Fatal("failed to open file: ", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Warn("failed to close file: ", err)
		}
	}()

	_, err = io.ReadAll(f)
	if err != nil {
		logrus.Fatal("failed readall: ", err)
	}
}

func ReadChunk() {
	f, err := os.Open("../fff/mf.txt")
	if err != nil {
		logrus.Fatal("failed to open file: ", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Warn("failed to close file: ", err)
		}
	}()

	b := make([]byte, 16)

	for {
		_, err = f.Read(b)

		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			logrus.Warn("unexpected error: ", err)
			break
		}
	}
}

func ReadBuff() {
	f, err := os.Open("../fff/mf.txt")
	if err != nil {
		logrus.Fatal("failed to open file: ", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Warn("failed to close file: ", err)
		}
	}()

	buff := bytes.NewBuffer(make([]byte, 16))

	for n := int64(0); err == nil; {
		buff.Reset()

		n, err = buff.ReadFrom(f)

		if err == io.EOF || n == 0 {
			break
		}

		if err != nil && err != io.EOF {
			logrus.Warn("unexpected error: ", err)
			break
		}
	}
}

func ReadBuffR() {
	f, err := os.Open("../fff/mf.txt")
	if err != nil {
		logrus.Fatal("failed to open file: ", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			logrus.Warn("failed to close file: ", err)
		}
	}()

	r := bufio.NewReader(f) // internally calls NewReaderSize(f, defaultSize)
	b := make([]byte, 16)

	for n := 0; err == nil; {
		n, err = r.Read(b)

		if err == io.EOF || n == 0 {
			break
		}

		if err != nil && err != io.EOF {
			logrus.Warn("unexpected error: ", err)
			break
		}
	}
}

func PipeRW() {
	pr, pw := io.Pipe()
	go func(w io.WriteCloser) {
		for _, s := range []string{"a string", "another string",
			"last one"} {
			fmt.Printf("-> writing %q\n", s)
			fmt.Fprint(w, s)
		}
		w.Close()
	}(pw)

	var err error
	for n, b := 0, make([]byte, 100); err == nil; {
		fmt.Println("<- waiting...")
		n, err = pr.Read(b)
		if err == nil {
			fmt.Printf("<- received %q\n", string(b[:n]))
		}
	}

	if err != io.EOF {
		fmt.Println("error:", err)
	}
}

func CopyFile(from, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
