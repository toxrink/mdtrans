package mdtrans

//ErrorCode 错误代码
type ErrorCode struct {
	Code int
	Info string
}

var (
	//E1 错误代码
	E1 = ErrorCode{Code: 1000, Info: "文件写入失败"}
	//E2 错误代码
	E2 = ErrorCode{Code: 1001, Info: "请求markdown转换html接口失败"}
	//E3 错误代码
	E3 = ErrorCode{Code: 1002, Info: "读取markdown转换html接口返回内容失败"}
	//E4 错误代码
	E4 = ErrorCode{Code: 1003, Info: "目标文件或目录不存在"}
	//E5 错误代码
	E5 = ErrorCode{Code: 1004, Info: "未知的目标文件类型"}
	//E6 错误代码
	E6 = ErrorCode{Code: 1005, Info: "pandoc.exe文件路径错误"}
	//E7 错误代码
	E7 = ErrorCode{Code: 1007, Info: "trans类型为nil"}
)
