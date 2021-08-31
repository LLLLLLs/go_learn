// Time        : 2019/08/27
// Description :

package main

import (
	"fmt"
	"golearn/pkgtest/graphql"
	"golearn/util"
	"io/ioutil"
	"net/http"
	"strconv"
)

func printReminder() {
	reminder := `请选择功能：
	1=创建学员 2=培养学员 3=放弃学员 4=学员列表 输入其他数字退出...`
	fmt.Println(reminder)
}

func main() {
	// printReminder()
	// var purpose int
	// _, err := fmt.Scanln(&purpose)
	// util.MustNil(err)
	// for {
	// 	switch purpose {
	// 	case 1:
	// 		create()
	// 	case 2:
	// 		train()
	// 	case 3:
	// 		reject()
	// 	case 4:
	// 		list()
	// 	default:
	// 		return
	// 	}
	// 	printReminder()
	// 	_, err := fmt.Scanln(&purpose)
	// 	util.MustNil(err)
	// }
	http.HandleFunc("/graphql", func(writer http.ResponseWriter, request *http.Request) {
		data, err := ioutil.ReadAll(request.Body)
		util.MustNil(err)
		fmt.Println(string(data))
		result := graphql.ExecQuery(string(data))
		_, err = writer.Write(util.Marshal(result))
		util.MustNil(err)
	})
	util.MustNil(http.ListenAndServe(":8080", nil))
}

var allField = `
		Id
		Name
		BeautyNo
		Sex      
		Talent
		Power
		Prof
		Status
		Exp
		RecoverRemain
`

func create() {
	var name string
	var beauty int
	fmt.Println("请输入学员名称")
	_, err := fmt.Scanln(&name)
	util.MustNil(err)
	fmt.Println("请输入名媛编号")
	_, err = fmt.Scanln(&beauty)
	util.MustNil(err)
	mutation := `mutation {
	create(Name: "` + name + `",BeautyNo: ` + strconv.Itoa(beauty) + `) {` + allField + `	}
}`
	fmt.Println(mutation)
	result := graphql.ExecQuery(mutation)
	fmt.Println(result)
}

func train() {
	var id string
	fmt.Println("请输入学员ID")
	_, err := fmt.Scanln(&id)
	util.MustNil(err)
	mutation := `mutation {
	train(Id: "` + id + `") {` + allField + `	}
}`
	fmt.Println(mutation)
	result := graphql.ExecQuery(mutation)
	fmt.Println(result)
}

func reject() {
	var id string
	fmt.Println("请输入学员ID")
	_, err := fmt.Scanln(&id)
	util.MustNil(err)
	mutation := `mutation {
	reject(Id:"` + id + `")
}`
	fmt.Println(mutation)
	result := graphql.ExecQuery(mutation)
	fmt.Println(result)
}

func list() {
	query := `query {
	list {` + allField + `	}
}`
	fmt.Println(query)
	result := graphql.ExecQuery(query)
	fmt.Println(result)
}
