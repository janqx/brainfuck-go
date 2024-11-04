package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/janqx/brainfuck-go"
)

const (
	REPL_PROMPT     = ">> "
	SOURCE_FILE_EXT = ".bf"
)

var (
	flagShowVersion bool
	flagShowHelp    bool
	flagCmd         string
)

func repl(ctx *brainfuck.Context) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(REPL_PROMPT)
		if !scanner.Scan() {
			fmt.Fprintln(os.Stderr, "failed to get input data from console")
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if line == "!exit" {
			break
		}
		err := ctx.Execute(line)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println()
	}
}

func runFile(ctx *brainfuck.Context, filename string) error {
	var err error
	var fullpath string
	if ext := filepath.Ext(filename); ext != SOURCE_FILE_EXT {
		return fmt.Errorf("invalid ext name: %s, except: %s", ext, SOURCE_FILE_EXT)
	}
	if fullpath, err = filepath.Abs(filename); err != nil {
		return err
	}
	var source []byte
	source, err = os.ReadFile(fullpath)
	if err != nil {
		return err
	}
	return ctx.Execute(string(source))
}

func main() {
	var err error

	flag.BoolVar(&flagShowVersion, "version", false, "show version information")
	flag.BoolVar(&flagShowHelp, "help", false, "show help information")
	flag.StringVar(&flagCmd, "c", "", "execute string")
	flag.Parse()

	if flagShowHelp {
		fmt.Printf("Usage: brainfuck-go [file] [options]\nOptions:\n")
		flag.PrintDefaults()
		return
	} else if flagShowVersion {
		fmt.Printf("brainfuck-go v%d.%d.%d\nrepository: %s", brainfuck.VERSION_MAJOR, brainfuck.VERSION_MINOR, brainfuck.VERSION_PATCH, "github.com/janqx/brainfuck-go")
		return
	}

	ctx := brainfuck.NewContext()

	if flagCmd != "" {
		err = ctx.Execute(flagCmd)
	} else {
		filename := flag.Arg(0)
		if filename == "" {
			repl(ctx)
		} else {
			err = runFile(ctx, filename)
		}
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
