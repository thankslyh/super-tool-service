package sliceUpload_test

import (
	"log"
	"sort"
	"super-tool-service/sliceUpload"
	"testing"
)

func TestFileSlice(t *testing.T) {
	fs := sliceUpload.NewFileSlices()
	*fs = append(*fs, sliceUpload.FileSlice{
		Index: 5,
		Filename: "第五个字符串",
		FileHash: "xxxxxx5",
	})
	*fs = append(*fs, sliceUpload.FileSlice{
		Index: 4,
		Filename: "第4个字符串",
		FileHash: "xxxxxx4",
	})
	*fs = append(*fs, sliceUpload.FileSlice{
		Index: 2,
		Filename: "第2个字符串",
		FileHash: "xxxxxx2",
	})
	*fs = append(*fs, sliceUpload.FileSlice{
		Index: 1,
		Filename: "第1个字符串",
		FileHash: "xxxxxx1",
	})
	*fs = append(*fs, sliceUpload.FileSlice{
		Index: 3,
		Filename: "第3个字符串",
		FileHash: "xxxxxx3",
	})

	sort.Sort(fs)

	log.Println(fs)
}
