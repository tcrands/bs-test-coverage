package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Walker struct {
	p  string
	fn filepath.WalkFunc
}

/*
	@function Walk - runs the inbuild file walker
	@return {error}
*/
func (w *Walker) Walk() error {
	return filepath.Walk(w.p, w.fn)
}

/////////////////////
// Walking Functions
/////////////////////

/*
	@function NewWalker - creates a new walker struct
	@param {String} p - the directory path to transverce
	@param {filepath.WalkFunc} fn - the function to run on each directory pass
	@return {Walker} - a new instance of a Walker
*/
func NewWalker(p string, fn filepath.WalkFunc) *Walker {
	return &Walker{
		p:  p,
		fn: fn,
	}
}

/*
	@function walkRootPath - transverses the root path
	@param {String} extention - the file extention to be isolated
	@return {filepath.WalkFunc} - a wrapper for custom logic on the directory transversal
*/
func walkDirectory(extention string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		if strings.Contains(path, extention) {
			if stringInArray(f.Name(), testableFiles) {
				fmt.Println(f.Name())
				file, _ := ioutil.ReadFile(path)
				testableFunctions := getFunctions(string(file))
				updateDatabase((f.Name() + strconv.Itoa(len(testableFunctions))), strconv.Itoa(0))
				markedFile := addMarkers(string(file), testableFunctions)
				fmt.Println(markedFile)
			}
		}
		return nil
	}
}

func addMarkers(file string, testableFunctions [][]string) string {
	for _, function := range testableFunctions {
		fmt.Println(function)
		splitSlice := strings.SplitAfter(file, function[0])
		file = splitSlice[0] + "\n MARKER \n" + splitSlice[1]
	}
	return file
}
