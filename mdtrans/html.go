package mdtrans

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

//HTMLTrans md转html
type HTMLTrans struct {
	APIPath string `default:"https://api.github.com/markdown/raw"`
}

//MarkDownTrans md转html
func (trans HTMLTrans) MarkDownTrans(transInfo TransInfo) []byte {
	fmt.Printf("transform markdown file %s\n", transInfo.SrcPath)
	file, _ := os.Open(transInfo.SrcPath)
	defer file.Close()
	resp, err := http.Post(trans.APIPath, "text/plain", file)
	if nil != err {
		fmt.Println(E2, err)
		return make([]byte, 0)
	}
	defer resp.Body.Close()

	content, err2 := ioutil.ReadAll(resp.Body)
	if nil != err2 {
		fmt.Println(E3, err2)
		return make([]byte, 0)
	}
	return content
}

//Save 保存html文件
func (trans HTMLTrans) Save(transInfo TransInfo) {
	if len(transInfo.DistContent) > 1 {
		body := string(transInfo.DistContent)
		t, _ := template.New(transInfo.DistName).Parse(PageTpl)
		buff := new(bytes.Buffer)
		t.Execute(buff, PageTplInfo{Title: transInfo.DistName, Body: body})
		write(buff.Bytes(), transInfo.DistPath)
	}
}
