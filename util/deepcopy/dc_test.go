// @author: lls
// @date: 2021/7/23
// @desc:
package deepcopy

import (
	"fmt"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type TestInterface interface {
	Hello()
}

type TestData struct {
	id                 *string
	roleId             string
	grantArmbandTime   int64
	lieutenantBelong   string
	status             int
	startTime          time.Time
	finishTime         time.Time
	marryCdAfterCancel time.Time // 寻缘、求亲取消后可再次发起的时间
	consumeType        int
	Search             Search
	appointRole        string
	Spouse             Spouse
	data               map[int]int
	t                  time.Time
}

func (t TestData) DeepCopy() interface{} {
	return &t
}

// 寻缘大厅匹配列表
type Search struct {
	RefreshTime time.Time // 下次可刷新时间
	List        []string  // 寻缘对象列表
	Sweetheart  string    // 心仪对象
}

// 配偶
type Spouse struct {
	RoleId       string // 归属玩家id
	OffId        string // 配偶id
	AddAttr      int64  // 配偶的属性加成
	MarriageTime int64  // 成婚时间
	Search1      Search
	Search2      Search
}

func (t TestData) Hello() {

}

func GetPtrUnExportFiled(s interface{}, filed string) reflect.Value {
	v := reflect.ValueOf(s).Elem().FieldByName(filed)
	// 必须要调用 Elem()
	return reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}

func SetPtrUnExportFiled(s interface{}, filed string, val interface{}) error {
	v := GetPtrUnExportFiled(s, filed)
	rv := reflect.ValueOf(val)
	if v.Kind() != v.Kind() {
		return fmt.Errorf("invalid kind, expected kind: %v, got kind:%v", v.Kind(), rv.Kind())
	}
	v.Set(rv)
	return nil
}

func TestUnexported(t *testing.T) {
	data := newTestData()
	fmt.Println(SetPtrUnExportFiled(data, "roleId", "changed"))
	fmt.Println(data)
}

type inherit struct {
	id string
}

type Outer struct {
	inherit
	oId string
}

type outer struct {
	inherit
	oId string
}

func TestInherit(t *testing.T) {
	id := "123"
	data := outer{
		inherit: inherit{id: id},
		oId:     id,
	}
	fmt.Println(Copy(data))
}

func TestOuter(t *testing.T) {
	id := "123"
	data := Outer{
		inherit: inherit{id: id},
		oId:     id,
	}
	fmt.Println(Copy(data))
}

type world struct {
	*publisher
	id          string
	transports  transports
	settlements map[int]settlement
	allWarriors map[int]int // 当前英雄的职务
}

type transports struct {
	publisher
	transports map[string]transport
	resetTime  int64 // 本轮重置时间
}

type transport struct {
	settlementNo int   // 居民地编号
	star         int   // 星级
	awardNo      int   // 奖励编号
	guardNo      int   // 护卫英雄编号
	guardPower   int64 // 护卫英雄实力
	beLooted     bool  // 是否被掠夺
	received     bool  // 是否领取奖励
	checked      bool  // 是否查看过,没看过需要飘气泡
}

type settlement struct {
	publisher
	no         int
	level      int
	population int
	admins     map[int]adminSeat // 位置编号->管理员
}

type adminSeat struct {
	warrior      int // 当前管理员
	replaceTimes int // 今日更换次数
}

type publisher struct {
	roleId string
}

func TestWorld(t *testing.T) {
	data := &world{
		publisher: &publisher{roleId: "123"},
		id:        "123",
		transports: transports{
			publisher: publisher{roleId: "123"},
			transports: map[string]transport{
				"123": {
					settlementNo: 1,
					star:         123,
					awardNo:      123,
					guardNo:      123,
					guardPower:   123,
					beLooted:     false,
					received:     false,
					checked:      false,
				},
			},
			resetTime: 0,
		},
		settlements: map[int]settlement{
			1: {
				publisher:  publisher{},
				no:         0,
				level:      0,
				population: 0,
				admins:     map[int]adminSeat{},
			},
		},
		allWarriors: nil,
	}
	fmt.Printf("%p\n", data.publisher)
	fmt.Printf("%+v\n", Copy(data))
}

func TestSlice(t *testing.T) {
	data := []*world{
		{id: "123"},
		{id: "456"},
	}
	cpy := Copy(data)
	fmt.Printf("%+v;%+v\n", data, cpy)
	cpy.([]*world)[0].id = "111"
	fmt.Printf("%v\n%v\n", data[0], cpy.([]*world)[0])
}

func newTestData() TestInterface {
	id := "id"
	return &TestData{
		id:     &id,
		roleId: "roleId",
		data:   map[int]int{},
		t:      time.Now(),
	}
}

func BenchmarkDeepCopy2(b *testing.B) {
	data := &Spouse{}
	for i := 0; i < b.N; i++ {
		cpy := Copy(data).(*Spouse)
		_ = cpy
	}
}

func BenchmarkDeepCopy3(b *testing.B) {
	data := newTestData()
	for i := 0; i < b.N; i++ {
		cpy := Copy(data).(TestInterface)
		_ = cpy
	}
}
