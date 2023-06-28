package test

import (
	"fmt"
	"testing"
	"time"
)

// @Title  test
// @Description  MyGO
// @Author  WeiDa  2023/6/27 16:54
// @Update  WeiDa  2023/6/27 16:54

var percent = 0

func TestDefer(t *testing.T) {

	var keepChecking = true
	// 开启下载
	fmt.Println("开始下载任务！")
	go download("/download", func() {
		keepChecking = false
		fmt.Println("下载完成！")
	}, func(currentPercent int) {
		keepChecking = false
		fmt.Println("下载取消！", currentPercent)
	}, func(currentPercent int) {
		keepChecking = false
		fmt.Println("下载失败！", currentPercent)
	})

	//开启检查下载进度
	for {
		if keepChecking {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("当前进度：", getPercent()*10, "%")
		} else {
			break
		}
	}
}

// 获取进度
func getPercent() int {
	return percent
}

// 下载
func download(url string, downloadSuccess func(), downloadCancelled func(int), downloadFailed func(int)) {
	for {
		time.Sleep(1 * time.Second)
		percent++
		if percent == 30 {
			downloadFailed(percent)
			break
		}
		if percent == 10 {
			downloadCancelled(percent)
			break
		}
		if percent >= 100 {
			downloadSuccess()
			break
		}
	}

}
