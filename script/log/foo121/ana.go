package foo121

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type blockT struct {
	time      int
	goroutine int
	content   []string
}

func LogAna() {
	filePath := "./202504031716.log"

	// 打开日志文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// 创建带缓冲的Scanner逐行读取
	scanner := bufio.NewScanner(file)

	blockList := make([]blockT, 0)
	block := blockT{}
	read := true
	//containLock := false
	repeat := map[string]bool{
		//"logic/player.go:1195":         false,
		//"game/handler/req_login.go:72": false,
		"gorilla/websocket/conn_read.go:12": true,
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			//if read && containLock && len(block.content) > 0 {
			if read && len(block.content) > 0 {
				blockList = append(blockList, block)
			}
			read = true
			//containLock = false
			block = blockT{}
			continue
		}
		if !read {
			continue
		}
		block.content = append(block.content, line)
		//if strings.Contains(line, "lock") {
		//	containLock = true
		//}

		if strings.Contains(line, "minutes]:") {
			minutes := strings.TrimSuffix(line, "minutes]:")
			minutes = strings.TrimSpace(minutes[len(minutes)-4 : len(minutes)-1])
			//fmt.Println(minutes)
			v, _ := strconv.Atoi(minutes)
			block.time = v
			if v < 250 {
				read = false
				block = blockT{}
				continue
			}
		}

		if len(line) > len("goroutine ") && line[:len("goroutine ")] == "goroutine " {
			goroutineLine := strings.TrimPrefix(line, "goroutine ")
			goroutine := strings.Split(goroutineLine, " ")[0]
			v, _ := strconv.Atoi(goroutine)
			fmt.Println(v)
			block.goroutine = v
		}
	}

	//sort.Slice(blockList, func(i, j int) bool {
	//	return blockList[i].time > blockList[j].time
	//})

	//sort.Slice(blockList, func(i, j int) bool {
	//	return len(blockList[i].content) > len(blockList[j].content)
	//})

	sort.Slice(blockList, func(i, j int) bool {
		return blockList[i].goroutine < blockList[j].goroutine
	})

	for i := 0; i < len(blockList); i++ {
		block := blockList[i]
		for _, line := range block.content {
			for key, value := range repeat {
				if strings.Contains(line, key) {
					if value {
						blockList = append(blockList[0:i], blockList[i+1:]...)
						i--
						continue
					}
					repeat[key] = true
				}
			}
		}
	}

	outputFile, err := os.Create("./block_analysis.txt")
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush() // 确保所有缓冲数据写入磁盘

	// 将blockList写入文件
	for _, block := range blockList {
		for _, line := range block.content {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				fmt.Printf("Error writing to file: %v\n", err)
				return
			}
		}
		_, err := writer.WriteString("\n") // 块之间加空行
		if err != nil {
			fmt.Printf("Error writing separator: %v\n", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
