package mdtrans

//TransInfo 转换文件信息
type TransInfo struct {
	SrcPath     string
	SrcName     string
	SrcDir      string
	SrcContent  []byte
	DistPath    string
	DistName    string
	DistDir     string
	DistType    string
	DistContent []byte
	TimeStamp   string
}

//Transform md文件转换
type Transform interface {
	MarkDownTrans(transInfo TransInfo) []byte
	Save(transInfo TransInfo)
}
