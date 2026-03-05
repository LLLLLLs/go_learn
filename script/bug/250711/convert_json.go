package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputPath := "input.json"
	outputPath := "output.json"

	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var objects []string
	var builder strings.Builder
	inObject := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		if line == "{" {
			inObject = true
			builder.Reset()
		}
		if inObject {
			builder.WriteString(line)
			builder.WriteString("\n")
		}
		if line == "}" && inObject {
			inObject = false
			objects = append(objects, builder.String())
		}
	}

	// 替换 Int32("数字") 为数字
	int32Re := regexp.MustCompile(`Int32\("([0-9]+)"\)`)

	var output []map[string]interface{}
	for _, rec := range objects {
		rec = int32Re.ReplaceAllString(rec, "$1")
		rec = regexp.MustCompile(`,\s*([}\]])`).ReplaceAllString(rec, "$1")
		var obj map[string]interface{}
		err := json.Unmarshal([]byte(rec), &obj)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse record: %v\n%s\n", err, rec)
			continue
		}
		if v, ok := obj["createtime"]; ok {
			switch vv := v.(type) {
			case float64:
				t := time.Unix(int64(vv), 0).In(time.FixedZone("UTC+8", 8*3600))
				obj["createtime"] = t.Format("2006-01-02 15:04:05")
			case string:
				if unix, err := strconv.ParseInt(vv, 10, 64); err == nil {
					t := time.Unix(unix, 0).In(time.FixedZone("UTC+8", 8*3600))
					obj["createtime"] = t.Format("2006-01-02 15:04:05")
				}
			}
		}
		output = append(output, obj)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	var sumDead int64
	for _, obj := range output {
		line, err := json.Marshal(obj)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to marshal object: %v\n", err)
			continue
		}
		outFile.Write(line)
		outFile.Write([]byte("\n"))

		// 统计 dead
		if atk, ok := obj["attackinfo"].(map[string]interface{}); ok {
			if deadVal, ok := atk["dead"]; ok {
				switch v := deadVal.(type) {
				case float64:
					sumDead += int64(v)
				case string:
					if n, err := strconv.ParseInt(v, 10, 64); err == nil {
						sumDead += n
					}
				}
			}
		}
	}
	// 将统计结果写入文件
	outFile.Write([]byte(fmt.Sprintf("dead总和: %d\n", sumDead)))
	fmt.Println("Converted JSON written to", outputPath)
}
