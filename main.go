package main

import (
	"io"
	"os"
	"path/filepath"

	"github.com/bjeight/fastats/fasta"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
	defer f.Close()

	var r *fasta.Reader
	switch filepath.Ext(os.Args[1]) {
	case ".gz", ".bgz":
		r = fasta.NewZReader(f)
	default:
		r = fasta.NewReader(f)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(1)
		}

		os.Stdout.WriteString(">" + record.Description + "\n")
		os.Stdout.Write(record.Seq)
		os.Stdout.WriteString("\n")
	}
}
