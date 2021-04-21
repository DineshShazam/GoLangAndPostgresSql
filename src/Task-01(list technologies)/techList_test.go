package main

import (
	"testing"
	"time"
)

func checkLengthOfList(t *testing.T) {
	d := newStack()

	if len(d) != 15 {
		t.Errorf("Technologies List Missing %v", len(d))
	}
}

func ReadWriteFile(t *testing.T) {
	// creating tech list
	techL := newStack()

	// writing it to a file
	currentTime := time.Now()
	techL.writeFile("_testingFileIO_" + currentTime.Format(`2006-01-02`))

	// reading the file
	fromFile := readFile("_testingFileIO_" + currentTime.Format(`2006-01-02`))

	// print it
	fromFile.print()

}
