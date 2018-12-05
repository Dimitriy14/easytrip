package baseController

import (
	"bufio"
	"os"
	"strings"

	"github.com/astaxie/beego"
)

func New() *BaseController {
	return &BaseController{}
}

type BaseController struct {
	beego.Controller
	Path string
	Lang string
}

func (b *BaseController) Set() {
}

func (r *BaseController) readFile() ([]string, error) {
	file, err := os.Open(r.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var lines []string
	in := bufio.NewScanner(file)
	for in.Scan() {
		lines = append(lines, in.Text())
	}
	return lines, in.Err()
}

func (r *BaseController) funcMakeMapTrans() (map[string]string, error) {
	strngs, err := r.readFile()
	if err != nil {
		return nil, err
	}
	tr := make(map[string]string)

	for _, str := range strngs {
		tmp := strings.Split(str, "=")
		tr[tmp[0]] = tmp[1]
	}
	return tr, nil
}

func (r *BaseController) Trans(s string) string {
	tmp, _ := r.funcMakeMapTrans()
	if _, ok := tmp[s]; ok {
		return tmp[s]
	}
	return s
}
