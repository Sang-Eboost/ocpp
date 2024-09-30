package main

import (
	"fmt"
	"time"
)

func main() {
	// Giá trị timestamp tính bằng milliseconds
	timestamp := int64(1727222400000)

	// Chuyển đổi sang thời gian (Time) bằng cách chia cho 1000 để chuyển từ milliseconds sang seconds
	t := time.Unix(0, timestamp*int64(time.Millisecond))

	// In ra thời gian
	fmt.Println("Time:", t)
}
