package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// GetSnapshot 生成视频缩略图（作为封面）
func GetSnapshot(filePath string, width, height int) {
	// 生成 CMD 命令
	cmd := exec.Command("ffmpeg", "-i", filePath, "-vframes", "1", "-s",
		fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")

	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	if err := cmd.Run(); err != nil {
		fmt.Println("could not generate frame", err)
	}

}
