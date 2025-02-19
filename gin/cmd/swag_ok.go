//go:build dev

package main

// 该函数仅在go build命令带上-tags dev后生效
func init() { swagOk = true }
