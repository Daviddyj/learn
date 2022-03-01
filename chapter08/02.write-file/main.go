package main

import (
	"log"
	"os"
)

func main() {
	filePath := "C:\\Users\\Jack\\Desktop\\Yijie.txt"
	writeFile(filePath) //ctrl+alt+m 将方法提取出来
	//writeFileWithAppend(filePath)

}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("this is first line---"))
	_, err = file.Write([]byte("this is first line\n"))
	_, err = file.WriteString("第二行内容5156156156\n")
	file.Sync() //告诉操作系统强制把内容写在磁盘上
	_, err = file.WriteAt([]byte("Changed456456456456"), 0)
	_, err = file.WriteString("小傻子嘤嘤嘤\n")
}
func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte("this is first line---"))
	_, err = file.Write([]byte("this is first line\n"))
	_, err = file.WriteString("第二行内容5156156156\n")
	file.Sync() //告诉操作系统强制把内容写在磁盘上
	_, err = file.WriteAt([]byte("DaiyijieChanged"), 0)
	if err != nil {
		log.Fatal(err)
	}
}
