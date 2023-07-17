package main

import (
	"archive/tar"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"

	//"path/filepath"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "CLI Application for handling tar archives",
	Long: `A CLI application that handles operations with tar archives. It can: 
1. Open an archive 
2. Search for mets.xml in the archive 
3. Compare diffs in mets.xml files in different archives`,
}

var openCmd = &cobra.Command{
	Use:   "open [path]",
	Short: "Open an archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		file, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "This path does not exist: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		r := tar.NewReader(file)
		for {
			_, err := r.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't read the archive: %v\n", err)
				os.Exit(1)
			}
		}
		fmt.Println("Archive has been successfully opened and read.")
	},
}

var searchCmd = &cobra.Command{
	Use:   "search [archive path]",
	Short: "Search for mets.xml in a tar archive",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		file, err := os.Open(path)
		if err != nil {
			fmt.Printf("Failed to open the file: %s", err)
			return
		}
		defer file.Close()

		r := tar.NewReader(file)
		for {
			header, err := r.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("Failed to read the archive: %s", err)
				return
			}
			if strings.HasSuffix(header.Name, "mets.xml") {
				fmt.Println("mets.xml is found in the archive.")
				return
			}
		}
		fmt.Println("mets.xml is not found in the archive.")
	},
}

var compareCmd = &cobra.Command{
	Use:   "compare [path1] [path2]",
	Short: "Compare two mets.xml files",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		xml1, err := readXML(args[0])
		if err != nil {
			fmt.Printf("Failed to read the first file: %s", err)
			return
		}

		xml2, err := readXML(args[1])
		if err != nil {
			fmt.Printf("Failed to read the second file: %s", err)
			return
		}

		if bytes.Equal(xml1, xml2) {
			fmt.Println("The files are identical.")
		} else {
			fmt.Println("The files are different.")
		}
	},
}

type XMLData map[string]interface{}

func readXML(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var xmlData XMLData
	err = xml.Unmarshal(data, &xmlData)
	if err != nil {
		return nil, err
	}

	canonicalXML, err := xml.Marshal(xmlData)
	if err != nil {
		return nil, err
	}

	return canonicalXML, nil
}

func main() {
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(compareCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %v\n", err)
		os.Exit(1)
	}
}
