package sliceUpload

import (
	"fmt"
	"os"
)

type FileSlice struct {
	Filename string `json:"filename"`
	FileHash string `json:"fileHash"`
	Index int `json:"index"`
	Data []byte `json:"data"`
	Size int
}

type fileSlices []FileSlice

func NewFileSlices() *fileSlices  {
	return &fileSlices{}
}

func (f *fileSlices) Len() int  {
	return len(*f)
}

func (f *fileSlices) Less(i, j int) bool  {
	return (*f)[i].Index < (*f)[j].Index
}

func (f *fileSlices) Swap(i, j int)  {
	(*f)[i], (*f)[j] = (*f)[j], (*f)[i]
}

func (f *fileSlices) MergeChunk() error {
	var data []byte
	var filename string
	for _, fs := range *f {
		data = append(data, fs.Data...)
		filename = fs.Filename
	}
	fmt.Println(len(data))
	fc, _ := os.Create(assetsPath + filename)
	defer fc.Close()
	_, err := fmt.Fprint(fc, string(data))
	return err
}