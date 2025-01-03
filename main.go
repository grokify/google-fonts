package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/grokify/mogo/os/osutil"
	"github.com/grokify/mogo/path/filepathutil"
	"github.com/grokify/mogo/type/stringsutil"
)

func main() {
	if len(os.Args) < 2 {
		slog.Error("enter a directory")
		os.Exit(1)
	}
	dir := os.Args[1]
	entries, err := osutil.ReadDirMore(dir, regexp.MustCompile(`\.ttf$`), false, true, false)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	lines := []string{
		fmt.Sprintf("package %s", dir),
		`import _ "embed"`}
	for _, ttf := range entries {
		anslug := filenameToSlugAlphaNumeric(ttf.Name())
		lines = append(lines,
			fmt.Sprintf("//go:embed %s", ttf.Name()),
			fmt.Sprintf("var font%s []byte", anslug),
			fmt.Sprintf("func %s() []byte { return font%s } ", anslug, anslug),
		)
	}

	body := strings.Join(lines, "\n")

	outpath := filepath.Join(dir, "fontbytes.go")

	err = os.WriteFile(outpath, []byte(body), 0600)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func filenameToSlugAlphaNumeric(s string) string {
	_, filename := filepath.Split(s)
	return stringsutil.ToAlphaNumeric(filepathutil.TrimExt(filename))
}
