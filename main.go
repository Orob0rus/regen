package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func main() {
	outputFilePath := flag.String("output_file", "README.md", "output path of the markdown file")
	flag.StringVar(outputFilePath, "o", "README.md", "output path of the markdown file")
	flag.Parse()
	if err := generateDoc(outputFilePath); err != nil {
		log.Errorf("[%s] occurred while attempting the markdown file generation!", err.Error())
		os.Exit(1)
	} else {
		log.Infof("Successfully generated the markdown file. Keep coding and documenting!")
	}
}

func generateDoc(outputFP *string) error {
	tmpFile, err := os.Create("/tmp/regen_doc.txt")
	if err != nil {
		return fmt.Errorf("[%s] occurred while creating intermediate file", err.Error())
	}
	defer tmpFile.Close()

	docCmd := exec.Command("/usr/bin/go")
	docCmd.Args = append(docCmd.Args, "doc")
	docCmd.Stdout = tmpFile
	if err := docCmd.Run(); err != nil {
		return fmt.Errorf("[%s] occurred while populating intermediate file", err.Error())
	}
	return nil
}
