// Time        : 2019/09/05
// Description :

package example

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golearn/sundry/profile"
	"golearn/sundry/profile/errors"
	"golearn/sundry/profile/loader"
	"sync"
	"testing"
)

//func TestMain(m *testing.M) {
//	testLoader := loader.NewTestConfDefiner()
//	testLoader.AddTable([]model.Phase{
//		{Index1: 1, Index2: 1, Index3: 1, Conf: "1"},
//		{Index1: 1, Index2: 1, Index3: 2, Conf: "2"},
//		{Index1: 1, Index2: 2, Index3: 1, Conf: "3"},
//		{Index1: 1, Index2: 2, Index3: 2, Conf: "4"},
//	})
//	testLoader.AddTable()
//	m.Run()
//}

// 自定义配置：
// 	1、可在TestMain中一次性定义所有配置
// 	2、也可在具体测试中再来指定配置
// 	3、以上2种方式只能2选1
func TestGetConfig(t *testing.T) {
	ast := assert.New(t)
	testLoader := loader.NewTestLoader()
	testLoader.AddTable([]Phase{
		{Index1: 1, Index2: 1, Index3: 1, Conf: "1"},
		{Index1: 1, Index2: 1, Index3: 2, Conf: "2"},
		{Index1: 1, Index2: 2, Index3: 1, Conf: "3"},
		{Index1: 1, Index2: 2, Index3: 2, Conf: "4"},
	})
	ctx := profile.ContextWithTestVersion(context.Background())
	center := profile.New(testLoader)
	// 找不到表
	p111, err := center.GetTable(ctx, struct{}{}).Get(1, 1, 1)
	ast.Equal(errors.ErrNoTable, err)
	ast.Nil(p111)
	// 正确获取所有配置
	all, err := center.GetTable(ctx, Phase{}).All()
	ast.Nil(err)
	ast.Equal(4, len(all.([]Phase)))
}

func TestGetConfig1(t *testing.T) {
	ast := assert.New(t)
	testLoader := loader.NewTestLoader()
	testLoader.AddTable([]Phase{
		{Index1: 2, Index2: 1, Index3: 1, Conf: "1"},
		{Index1: 2, Index2: 1, Index3: 2, Conf: "2"},
		{Index1: 2, Index2: 2, Index3: 1, Conf: "3"},
		{Index1: 2, Index2: 2, Index3: 2, Conf: "4"},
	})
	ctx := profile.ContextWithTestVersion(context.Background())
	center := profile.New(testLoader)
	// 成功获取配置
	p211, err := center.GetTable(ctx, Phase{}).Get(2, 1, 1)
	ast.Nil(err)
	ast.Equal("1", p211.(Phase).Conf)
	// 找不到配置
	p111, err := center.GetTable(ctx, Phase{}).Get(1, 1, 1)
	ast.Equal(errors.ErrNoConfig, err)
	ast.Nil(p111)
	// context未携带版本信息
	ctx = context.Background()
	all, err := center.GetTable(ctx, Phase{}).All()
	ast.Equal(errors.ErrContextMustWithVersion, err)
	ast.Nil(all)
}

func BenchmarkGetConf(b *testing.B) {
	testLoader := loader.NewTestLoader()
	testLoader.AddTable([]Phase{
		{Index1: 2, Index2: 1, Index3: 1, Conf: "1"},
		{Index1: 2, Index2: 1, Index3: 2, Conf: "2"},
		{Index1: 2, Index2: 2, Index3: 1, Conf: "3"},
		{Index1: 2, Index2: 2, Index3: 2, Conf: "4"},
	})
	ctx := profile.ContextWithTestVersion(context.Background())
	center := profile.New(testLoader)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := center.GetTable(ctx, Phase{}).Get(2, 1, 1)
		if err != nil {
			panic(err)
		}
	}
}

func TestGoroutine(t *testing.T) {
	testLoader := loader.NewTestLoader()
	testLoader.AddTable([]Phase{
		{Index1: 2, Index2: 1, Index3: 1, Conf: "1"},
		{Index1: 2, Index2: 1, Index3: 2, Conf: "2"},
		{Index1: 2, Index2: 2, Index3: 1, Conf: "3"},
		{Index1: 2, Index2: 2, Index3: 2, Conf: "4"},
	})
	ctx := profile.ContextWithTestVersion(context.Background())
	center := profile.New(testLoader)
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conf, err := center.GetTable(ctx, Phase{}).Get(2, 1, 1)
			fmt.Println(i, conf, err)
		}(i)
	}
	wg.Wait()
}
