package githubDocsHandle

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"super-tool-service/mdToHtml"
	"sync"
)

func Entry()  {
	dir, _ := os.Getwd()
	var wg sync.WaitGroup
	filepath.WalkDir("./assets/learn-and-share", func(path string, d fs.DirEntry, err error) error {
		if d.Name() == "readme.md" {
			wg.Add(1)
			absUrl := dir + "/" + path
			go TranslateHtml(absUrl, &wg)
		}
		return nil
	})
	wg.Wait()
}

func TranslateHtml(path string, wg *sync.WaitGroup)  {
	defer wg.Done()
	data, err := readFileContent(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	strHtml, _ := mdToHtml.MdToHtml(data)
	f, err := os.Create(createHtmlPath(path))
	if err != nil {
		return
	}
	size, err := fmt.Fprint(f, strHtml)
	if err != nil {
		fmt.Println(err, size)
		return
	}
}

func createHtmlPath(path string) string {
	return strings.Replace(path, ".md", ".html", 2)
}
func readFileContent(path string) ([]byte, error)  {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}