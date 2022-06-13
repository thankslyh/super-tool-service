package sliceUpload

type FileSlice struct {
	Filename string `json:"filename"`
	FileHash string `json:"fileHash"`
	Index int `json:"index"`
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

func (f *fileSlices) MergeHash() string  {
	var str string
	for _, fs := range *f {
		str += fs.FileHash
	}
	return str
}