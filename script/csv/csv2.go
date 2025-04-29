package csv

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func DataConvert2() {
	inputFile := "./data.csv"     // 输入文件路径
	outputFile := "./output2.csv" // 输出文件路径

	// 读取CSV文件
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// 去重和时间转换
	uniqueRows := make(map[string][]string)
	for i, row := range rows[1:] { // 从第二行开始
		region, err := strconv.Atoi(row[1])
		if err != nil || region >= 1500 {
			continue // 跳过region大于等于1500的行
		}

		key := strings.TrimSpace(row[0]) // 去除首尾空格
		row[0] = key
		timeStr := row[len(row)-1] // 最后一列为time列

		// 转换时间为UTC+8
		timestamp, err := strconv.ParseInt(timeStr, 10, 64)
		if err != nil {
			fmt.Printf("Error parsing timestamp in row %d: %v\n", i+2, err)
			continue
		}
		utcTime := time.Unix(timestamp, 0)
		localTime := utcTime.UTC().Add(8 * time.Hour).Format("2006-01-02 15:04:05")

		row[len(row)-1] = localTime // 更新time列

		// 处理 resp 列
		resp := row[2] // resp 列
		if resp != "" {
			// 解析 InfoList
			var infoList []map[string]interface{}
			if err := json.Unmarshal([]byte(resp), &infoList); err != nil {
				fmt.Printf("Error parsing resp in row %d: %v\n", i+2, err)
				continue
			}

			// 过滤 InfoList
			filteredInfoList := make([]map[string]interface{}, 0)
			for _, info := range infoList {
				heroUidList, ok := info["HeroUidList"].([]interface{})
				if !ok {
					continue
				}

				// 检查是否包含 "NPC"
				containsNPC := false
				for _, heroUid := range heroUidList {
					if strings.Contains(heroUid.(string), "NPC") {
						containsNPC = true
						break
					}
				}

				if containsNPC {
					filteredInfoList = append(filteredInfoList, info)
				}
			}

			// 更新 resp 列
			if len(filteredInfoList) > 0 {
				filteredResp, err := json.Marshal(filteredInfoList)
				if err != nil {
					fmt.Printf("Error marshaling filtered resp in row %d: %v\n", i+2, err)
					continue
				}
				row[2] = string(filteredResp)
			} else {
				row[2] = "" // 如果过滤后为空，则清空 resp 列
			}
		}

		// 去重处理
		if _, exists := uniqueRows[key]; !exists {
			uniqueRows[key] = row
		}
	}

	// 写入新的CSV文件
	newFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer newFile.Close()

	writer := csv.NewWriter(newFile)
	defer writer.Flush()

	// 写入表头
	writer.Write(rows[0])

	// 写入去重后的数据
	for _, row := range uniqueRows {
		writer.Write(row)
	}

	fmt.Println("Processing completed.")
}
