package mdtrans

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

//PandocTrans Pandoc转换
type PandocTrans struct {
	PandocPath string `default:"pandoc.exe"`
	DistType   string `default:"html"`
}

//MarkDownTrans md转html
func (trans PandocTrans) MarkDownTrans(transInfo TransInfo) []byte {
	fmt.Printf("transform markdown file %s\n", transInfo.SrcPath)
	distPath := os.TempDir() + string(filepath.Separator) + transInfo.DistName
	cmd := exec.Command(trans.PandocPath, transInfo.SrcPath, "-o", distPath)
	cmd.Run()

	file, _ := os.Open(distPath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	writer := bufio.NewWriter(buf)

	r, _ := regexp.Compile("id=\"([^\"]+)\"")
	r2, _ := regexp.Compile("\\.")

	for scanner.Scan() {
		t := scanner.Text()
		if strings.HasPrefix(t, "<h3") {
			sub := r.FindStringSubmatch(t)
			if strings.Index(sub[1], ".") != -1 {
				id := r2.ReplaceAllString(sub[1], "")
				t = strings.Replace(t, sub[1], id, 1)
			}
		}
		writer.WriteString(t)
		writer.WriteRune('\n')
	}

	return buf.Bytes()
}

//Save 保存html文件
func (trans PandocTrans) Save(transInfo TransInfo) {
	if len(transInfo.DistContent) > 1 {
		body := string(transInfo.DistContent)
		t, _ := template.New(transInfo.DistName).Parse(PageTpl)
		buff := new(bytes.Buffer)
		t.Execute(buff, PageTplInfo{Title: transInfo.DistName, Body: body})
		write(buff.Bytes(), transInfo.DistPath)
	}
}
