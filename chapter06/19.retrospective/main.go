package main

import (
	"fmt"
	"time"
)

type Downloader struct {
	fileNameCh chan string
	finishedCh chan string
}

func main() {
	fileNameCh := make(chan string, 10)
	finishedCh := make(chan string)
	workerCounter := 10
	for i := 0; i < workerCounter; i++ {
		go func() {
			(&Downloader{
				//name : string
				fileNameCh: fileNameCh,
				finishedCh: finishedCh,
			}).Download()

		}()
	}

	fileNumber := 100
	fileNames := make([]string, 0, fileNumber)
	finishedFileCount := 0
	go func() {
		for finishedFile := range finishedCh {
			fmt.Println("文件:", finishedFile, "处理完毕")
			finishedFileCount++
			if finishedFileCount == fileNumber {
				close(finishedCh)
			}
		}
	}()

	for i := 0; i < fileNumber; i++ {
		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
	}
	for _, fileItem := range fileNames {
		fileNameCh <- fileItem
	}
	close(fileNameCh)

	fmt.Println("所有文件处理完毕，结束")

}

func (d *Downloader) Download() {
	for fileName := range d.fileNameCh {
		fmt.Println("开始下载文件:", fileName)
		time.Sleep(1 * time.Second)
		fmt.Println("开始处理文件:", fileName)
		time.Sleep(1 * time.Millisecond)
		fmt.Println("保存文件:", fileName)
		d.finishedCh <- fileName
	}
}
