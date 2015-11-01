package main

import (
	"fmt"
)

type TargetFile struct {
	files []string
}

func NewTargetFile(files []string) *TargetFile {
	targetFile := &TargetFile{files: files}
	return targetFile
}

func (target *TargetFile) GetLen() int {
	return len(target.files)
}

func (target *TargetFile) Shift() string {
	val := target.files[0]
	target.files = target.files[1:]
	return val
}

func main() {
	files := []string{"/var/tmp/1.txt", "/var/tmp/2.txt", "/var/tmp/3.txt"}

	target := NewTargetFile(files)

	fmt.Printf("target len = %d\n", target.GetLen())
	fmt.Printf("shift val = %s\n", target.Shift())
	fmt.Printf("target len = %d\n", target.GetLen())
	fmt.Printf("shift val = %s\n", target.Shift())
	fmt.Printf("target len = %d\n", target.GetLen())
	fmt.Printf("shift val = %s\n", target.Shift())
}
