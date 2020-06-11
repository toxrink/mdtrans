package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	. "mdtrans/mdtrans"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	mode         = flag.String("mode", "pandoc", "转换工具或接口类型")
	markdownPath = flag.String("path", "", "markdown文件路径或目录,默认当前目录下的md文件")
	storePath    = flag.String("store", "", "文件转换后的保存目录,默认和源文件同一路径")
	target       = flag.String("target", "html", "要转换的目标文件格式[html]")
	distName     = flag.String("name", "", "输出的文件名")
	pandoc       = flag.String("pandoc", "pandoc.exe", "pandoc文件路径")
)

func main() {
	flag.Parse()

	if "" == *markdownPath {
		currentPath, _ := filepath.Abs(".")
		filepath.Walk(currentPath, func(path string, info os.FileInfo, err error) error {
			if isMdFile(info.Name()) {
				dotrans(buildTransInfo(path, false))
			}
			return nil
		})
		for _, a := range os.Args {
			if isMdFile(a) {
				mdPath, _ := filepath.Abs(a)
				dotrans(buildTransInfo(mdPath, false))
			}
		}
	} else {
		_, err := os.Stat(*markdownPath)
		if nil != err {
			fmt.Println(E4, err)
			shutdown()
		}
		if isMdFile(*markdownPath) {
			dotrans(buildTransInfo(*markdownPath, true))
		} else {
			filepath.Walk(*markdownPath, func(path string, info os.FileInfo, err error) error {
				if isMdFile(info.Name()) {
					dotrans(buildTransInfo(path, false))
				}
				return nil
			})
		}
	}
}

func isMdFile(name string) bool {
	return strings.HasSuffix(name, ".md") || strings.HasSuffix(name, ".markdown")
}

func buildTransInfo(md string, useDistName bool) TransInfo {
	ti := TransInfo{}
	ti.TimeStamp = getTimestamp()

	fi, _ := os.Stat(md)
	ti.SrcName = fi.Name()
	ti.SrcPath = md
	ti.SrcDir = filepath.Dir(md)
	ti.SrcContent, _ = ioutil.ReadFile(md)

	if "" == *storePath {
		ti.DistDir = ti.SrcDir
	} else {
		ti.DistDir = *storePath
	}
	ti.DistType = *target

	if useDistName && "" != *distName {
		ti.DistName = *distName
	} else {
		index := strings.Index(ti.SrcName, ".")
		ti.DistName = ti.SrcName[:index] + "." + ti.DistType
	}

	ti.DistPath = filepath.Join(ti.DistDir, ti.DistName)

	return ti
}

func dotrans(transinfo TransInfo) {
	trans := getTransform(transinfo)
	if nil == trans {
		fmt.Println(E7)
		shutdown()
	}
	transinfo.DistContent = trans.MarkDownTrans(transinfo)
	trans.Save(transinfo)
}

func getTransform(transInfo TransInfo) Transform {
	if *mode == "pandoc" {
		if !exist(*pandoc) {
			fmt.Println(E6, *pandoc)
			return nil
		}
		return PandocTrans{PandocPath: *pandoc, DistType: transInfo.DistType}
	}
	switch transInfo.DistType {
	case "html":
		return HTMLTrans{APIPath: "https://api.github.com/markdown/raw"}
	default:
		fmt.Println(E5, transInfo.DistType)
		shutdown()
		return nil
	}
}

//shutdown 强制关闭程序
func shutdown() {
	os.Exit(-1)
}

//exist 文件是否存在
func exist(p string) bool {
	_, e := os.Stat(p)
	if nil == e {
		return true
	}
	return false
}

func getTimestamp() string {
	return time.Now().Format("20060102150405")
}
