// Time        : 2019/09/02
// Description :

package mongotest

import (
	"fmt"
	model2 "golearn/sundry/mongo-test/model"
	mongodb "golearn/sundry/mongo-test/mongo-db"
	"testing"
)

func TestMain(m *testing.M) {
	mongodb.InitClient("mongodb://localhost")
	client = mongodb.GetClient()
	m.Run()
}

func TestInsertStudent(t *testing.T) {
	insertStudent(model2.StudentValue{
		Id:            "1234",
		Name:          "li si",
		BeautyNo:      1,
		Sex:           2,
		Talent:        3,
		Power:         999,
		Prof:          2,
		Status:        2,
		Exp:           10,
		RecoverRemain: 1673,
	})
}

func TestQueryStudent(t *testing.T) {
	student := queryStudent("1234")
	fmt.Printf("%+v\n", student)
}

func TestInsertRole(t *testing.T) {
	insertRole("role1", 4)
}

func TestQueryRole(t *testing.T) {
	role := queryRole("role")
	fmt.Printf("%+v\n", role)
}

func BenchmarkQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryRole("role")
	}
}

func TestQueryAll(t *testing.T) {
	queryAll()
}

func BenchmarkQueryAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		queryAll()
	}
}

func TestInsertTest(t *testing.T) {
	insertTest()
}

func TestInsertPhase(t *testing.T) {
	insertPhase()
}
