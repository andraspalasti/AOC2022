package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"text/template"
)

//go:embed tmpls/*.go
var files embed.FS

func main() {
	day := flag.Int("day", 1, "the day of the problem")
	flag.Parse()

	if *day < 1 || 25 < *day {
		flag.PrintDefaults()
		log.Fatalf("day flag has to be between 1 and 25, got %d", *day)
	}

	ts, err := template.ParseFS(files, "tmpls/*.go")
	if err != nil {
		log.Fatalf("parsing tmpls directory: %s", err)
	}

	mainFilename := path.Join(dirname(), "..", "..", fmt.Sprintf("day%02d", *day), "main.go")
	testFilename := path.Join(dirname(), "..", "..", fmt.Sprintf("day%02d", *day), "main_test.go")

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	if _, err = os.Stat(mainFilename); err == nil {
		log.Fatalf("file already exists: %s", mainFilename)
	}

	if _, err = os.Stat(testFilename); err == nil {
		log.Fatalf("file already exists: %s", testFilename)
	}

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %s", err)
	}

	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %s", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	ts.ExecuteTemplate(testFile, "main_test.go", nil)
}

func dirname() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get the current filename")
	}
	return filepath.Dir(filename)
}
