/*
 * @Description:
 * @Author: violeteve
 * @Date: 2020-02-19 16:04:02
 * @LastEditors: violeteve
 * @LastEditTime: 2020-02-19 17:04:29
 * @FilePath: /go-github-actions/main.go
 */
package main

import (
	"fmt"

	"github.com/duangfei/go-github-actions/hello"
)

func main() {
	fmt.Println(hello.Greet())
}
