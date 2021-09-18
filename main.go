package main

import (
	"os"
)

func main() {
	os.Exit(Run())
}

func Run() int {
	files := []string{"./test0", "./test1", "./test2"}

	ch1 := make(chan []byte, 3)
	ch2 := make(chan []byte, 3)
	ch3 := make(chan []byte, 3)
	chs := []chan []byte{ch1, ch2, ch3}
	// chs := make([]chan []byte, 3)

	fc := func(i int, f *os.File) {
		defer close(chs[i])
		buf := make([]byte, 4096)
		for io := 0; io < 3; io++ {
			n, err := f.Read(buf)
			if err != nil {
				return
			}
			chs[i] <- buf[:n]
		}
	}

	for i, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return 1
		}
		defer f.Close()
		go fc(i, f)
	}

	var data []byte
	for {
		select {
		case data = <-chs[0]:
		case data = <-chs[1]:
		case data = <-chs[2]:
		}

		os.Stdout.Write(data)
	}
}
