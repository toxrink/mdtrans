package mdtrans

import (
	"fmt"
	"io/ioutil"
)

func write(content []byte, dist string) {
	err := ioutil.WriteFile(dist, content, 0666)
	if nil != err {
		fmt.Println(E1, err)
	}
	fmt.Println(dist)
}
