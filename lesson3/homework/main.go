package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Options struct {
	From      string
	To        string
	Offset    int
	Limit     int
	BlockSize int
	Conv      string
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.From, "from", "", "file to read. by default - stdin")
	flag.StringVar(&opts.To, "to", "", "file to write. by default - stdout")
	flag.IntVar(&opts.Offset, "offset", 0, "bytes count in one block")
	flag.IntVar(&opts.Limit, "limit", 1e9, "max bytes to read (default: all)")
	flag.IntVar(&opts.BlockSize, "block-size", 1024, "read/write block size in bytes")
	flag.StringVar(&opts.Conv, "conv", "", "transformations: upper_case, lower_case, trim_spaces")
	flag.Parse()

	return &opts, nil
}

func ReadChunk(reader io.Reader, blockSize int) (string, error) {
	buf := make([]byte, blockSize)
	n, err := reader.Read(buf)
	if err != nil {
		if err == io.EOF {
			return "", err
		}
		return "", fmt.Errorf("cannot read input file: %w", err)
	}
	return string(buf[:n]), nil
}

type Formatter struct {
	writer io.Writer
	up     bool
	low    bool
	trim   bool
	//tbegin, tend   bool
	//s_begin, s_end string
}

func NewFormatter(writer io.Writer, conv string) (*Formatter, error) {
	f := &Formatter{writer: writer, up: false, low: false, trim: false}
	args := strings.Split(conv, ",")
	for _, arg := range args {
		if arg == "upper_case" {
			f.up = true
		} else if arg == "lower_case" {
			f.low = true
		} else if arg == "trim_spaces" {
			f.trim = true
			//f.tbegin = true
			//f.tend = true
		}
	}
	if f.up && f.low {
		return nil, fmt.Errorf("cannot use 'lower_case' and 'upper_case' at the same time")
	}
	return f, nil
}

func (out *Formatter) WriteChunk(s string) error {
	if out.low {
		//strings.ToLower(s)
		runes := []rune(s)
		for _, r := range runes {
			r = unicode.ToLower(r)
		}
		s = string(runes)
	} else if out.up {
		runes := []rune(s)
		for _, r := range runes {
			r = unicode.ToUpper(r)
		}
		s = string(runes)
	}

	//if out.trim && out.tbegin {
	//	s = strings.TrimLeft(s, " ")
	//	if s != "" {
	//		out.tbegin = false
	//	}
	//}
	//
	//if out.trim && out.tend {
	//	out.s_end += s
	//}

	if out.trim {
		s = strings.TrimSpace(s)
	}

	_, err := io.WriteString(out.writer, s)
	if err != nil {
		return fmt.Errorf("cannot write input file: %w", err)
	}
	return nil
}

func SkipOffset(input io.Reader, offset int) error {
	if file, ok := input.(*os.File); ok && offset > 0 {
		size, err := file.Seek(0, io.SeekEnd)
		if err != nil {
			return fmt.Errorf("ошибка определения размера файла: %v", err)
		}
		if int64(offset) > size {
			return fmt.Errorf("ошибка: offset превышает размер файла")
		}
		_, err = file.Seek(int64(offset), io.SeekStart)
		if err != nil {
			return fmt.Errorf("ошибка позиционирования: %v", err)
		}
		return nil
	}

	// Для не-файловых reader'ов (например, stdin)
	_, err := io.CopyN(io.Discard, input, int64(offset))
	if err != nil && err != io.EOF {
		return fmt.Errorf("ошибка пропуска offset: %v", err)
	}
	return nil
}

//func Check(err error, message string) {
//	if err != nil {
//		fmt.Fprintln(os.Stderr, err)
//		os.Exit(1)
//	}
//}

func main() {
	opts, err := ParseFlags()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)

		os.Exit(1)
	}

	// todo: implement the functional requirements described in read.me
	var input io.Reader
	if opts.From == "" {
		input = os.Stdin
	} else {
		file, err := os.Open(opts.From)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open input file: %v", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	var output *Formatter
	if opts.To == "main.go" {
		fmt.Fprintf(os.Stderr, "cannot open main.go as output file")
		os.Exit(1)
	}
	if opts.To == "" {
		output, err = NewFormatter(os.Stdout, opts.Conv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	} else {
		file, err := os.Create(opts.To)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open output file: %v", err)
			os.Exit(1)
		}
		defer file.Close()
		output, err = NewFormatter(file, opts.Conv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
	}

	//err = SkipOffset(input, opts.Offset)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "%v", err)
	//	os.Exit(1)
	//}
	if s, err := ReadChunk(input, opts.Offset); (err != nil && err != io.EOF) || (len(s) < opts.Offset) {
		if err == io.EOF {
			fmt.Fprintf(os.Stderr, "opts.Offset is more than file size")
		} else {
			fmt.Fprintf(os.Stderr, "%v", err)
		}
		os.Exit(1)
	}

	for opts.Limit > 0 {
		s, err := ReadChunk(input, min(opts.BlockSize, opts.Limit))
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(1)
		}
		opts.Limit -= min(opts.BlockSize, opts.Limit)
		if err := output.WriteChunk(s); err != nil {
			fmt.Fprintf(os.Stderr, "problems with writing in file %v", err)
			os.Exit(1)
		}
	}
}
