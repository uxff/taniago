package utils

import "fmt"

func Size4Human(fileSize int64) string {
	if fileSize < 1024 {
		return fmt.Sprintf("%dB", fileSize)
	}
	if fileSize < 1024*1024 {
		return fmt.Sprintf("%.01fkB", float32(fileSize)/1024)
	}
	if fileSize < 1024*1024*1024 {
		return fmt.Sprintf("%.01fMB", float32(fileSize)/1024/1024)
	}
	return fmt.Sprintf("%.02fGB", float32(fileSize)/1024/1024/1024)
}
