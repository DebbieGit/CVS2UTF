// main.go
package main

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"os"
)

//Deze werkt goed
var end string

//Reading files requires checking most calls for errors. This helper will streamline our error checks below.R
func check(e error) {
	if e != nil {
		panic(e)
	}
	if e == io.EOF {
		fmt.Println("EOF reached")
	}
}

func main() {
	var filein string
	var fileout string
	if len(os.Args) > 2 {
		filein = os.Args[1]
		fileout = os.Args[2]
		importWinFile(filein, fileout)
	} else {
		fmt.Println("To transform CVS to UTF8")
		fmt.Println("Use MakeUTF8.exe <file.cvs> <file.utf>")
		fmt.Println("Uses current path")
	}
}

func writeBytes(f *os.File, theBytes []byte) {
	//You can Write byte slices as you’d expect.
	_, err := f.Write(theBytes)
	check(err)
}

func importWinFile(file1 string, file2 string) {
	end = "\n"
	mb, err := ioutil.ReadFile(".\\" + file1)
	check(err)
	n := len(mb)

	//For more granular writes, open a file for writing.
	f1, err := os.Create(file2)
	check(err)

	//It’s idiomatic to defer a Close immediately after opening a file.
	defer f1.Close()
	enc := charmap.Windows1252
	//Nu krijg je een transformer
	trf := enc.NewDecoder()
	pb, i1, err := transform.Bytes(trf, mb)
	check(err)
	if i1 == n {
		fmt.Println("Alles ok")
	} else {
		fmt.Println("FOUT: v% - v%", n, i1)
	}
	writeBytes(f1, pb)
}
