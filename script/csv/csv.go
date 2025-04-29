package csv

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func DataConvert() {
	inputFile := "./data.csv"    // 输入文件路径
	outputFile := "./output.csv" // 输出文件路径

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

	// 提取去重后的数据并排序
	sortedRows := make([][]string, 0, len(uniqueRows))
	for _, row := range uniqueRows {
		sortedRows = append(sortedRows, row)
	}

	// 按时间列排序
	sort.Slice(sortedRows, func(i, j int) bool {
		timeI := sortedRows[i][len(sortedRows[i])-1]
		timeJ := sortedRows[j][len(sortedRows[j])-1]
		return timeI < timeJ
	})

	// 写入去重后的数据
	for _, row := range sortedRows {
		writer.Write(row)
	}

	fmt.Println("Processing completed.")
}
