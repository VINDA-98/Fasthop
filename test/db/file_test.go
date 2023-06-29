package db

import (
	"fmt"
	"os"
	"testing"
)

// @Title  db
// @Description  MyGO
// @Author  WeiDa  2023/6/29 14:40
// @Update  WeiDa  2023/6/29 14:40
var filepath = "./db.txt"

func TestFile(t *testing.T) {
	if len(os.Args) >= 3 {
		appendContent(os.Args[1], os.Args[2])
	} else {
		fmt.Println(outputContent())
	}
}

// 追加内容到文件
func appendContent(content string, name string) {
	createFile()
	appendContentToFile("增加" + name + "的数据内容：" + content)
}

// 输出文件现有内容
func outputContent() string {
	fileData, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(fileData)
}

// 创建文件，若文件存在则不创建
func createFile() {
	if fileInfo, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) || fileInfo.IsDir() {
			// 文件不存在，或文件是目录，则创建文件
			file, err := os.Create(filepath)
			if err != nil {
				fmt.Println("创建文件失败，错误信息：", err)
				return
			}
			// 关闭文件操作
			file.Close()
		} else {
			// 其它错误
			panic(err)
		}
	}
}

// 向文件追加内容，每次追加以换行间隔
func appendContentToFile(content string) {
	file, err := os.OpenFile(filepath, os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.Write([]byte(content + "\n"))
	if err != nil {
		fmt.Println("写文件操作出现错误，异常信息：", err)
		return
	} else {
		fmt.Println("数据成功录入！")
	}
}
