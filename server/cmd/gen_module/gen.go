package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/duke-git/lancet/v2/fileutil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TemplateContext struct {
	Name        string // 模块名称首字母大写
	ServiceName string // 服务名称首字母小写
	DbName      string // 数据库名称
	CName       string // 模块中文名
}

func main() {
	name := flag.String("name", "", "模块名称")
	cname := flag.String("cname", "", "模块中文名称")
	flag.Parse()

	if *name == "" || *cname == "" {
		panic("name or cname 不能为空")
	}

	if err := genServer(*name, *cname); err != nil {
		panic(err)
	}
	//genClient(*name, *cname)
}

func genServer(name, cname string) error {
	info := formatModuleContext(name, cname)
	fmt.Printf("info %+v\n", info)

	rootDir, _ := filepath.Abs(path.Join("cmd", "gen_module"))
	// dir := path.Join(rootDir, "output", "server", info.DbName)

	// if fileutil.IsExist(dir) {
	// 	return errors.New("文件夹已存在")
	// } else {
	// 	fileutil.CreateDir(dir)
	// }

	modNames := []string{"api", "service", "model", "permission"}
	for _, modName := range modNames {
		// 生成的文件
		var filePath, _ = filepath.Abs(path.Join("modules", modName, info.DbName+".go"))
		if modName == "model" {
			filePath, _ = filepath.Abs(path.Join("dal", "model", info.DbName+".go"))
		}
		if modName == "permission" {
			filePath, _ = filepath.Abs(path.Join("consts", "permission", info.DbName+".go"))
		}
		if fileutil.IsExist(filePath) {
			return errors.New("文件已存在: " + filePath)
		}
		// 模板文件路径
		tempFilePath, _ := filepath.Abs(path.Join(rootDir, "template", "tpl_server", modName+".tpl"))

		// 加载模板文件
		tempFile, _ := template.ParseFiles(tempFilePath)

		// 写入
		writeFileErr := os.WriteFile(filePath, []byte(""), 0755)
		if writeFileErr != nil {
			return writeFileErr
		}
		file, _ := os.OpenFile(filePath, os.O_RDWR, 0755)
		err := tempFile.Execute(file, info)
		if err != nil {
			return err
		}
	}
	return nil
}

func genClient(name, cname string) {

}

func formatModuleContext(name, cname string) *TemplateContext {
	n := strings.ReplaceAll(name, " ", "")
	cname = strings.ReplaceAll(cname, " ", "")

	return &TemplateContext{
		Name: cases.Title(language.English).String(n),
		// 首字母小写
		ServiceName: n,
		DbName:      strings.ToLowerSpecial(unicode.SpecialCase{}, n),
		CName:       cname,
	}
}
