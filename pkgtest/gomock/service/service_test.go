// Time        : 2019/10/17
// Description :

package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	repomock "golearn/pkgtest/gomock/repo-mock"
	"testing"
)

func TestService_Method(t *testing.T) {
	ast := assert.New(t)
	ctrl := gomock.NewController(t)
	repo := repomock.NewMockRepository(ctrl)

	// 创建，打印输入参数，并返回err=nil
	repo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Do(func(key string, value interface{}) {
			t.Logf("创建成功,key=%s,value=%v", key, value)
		}).Return(nil).AnyTimes()
	// 获取，参数="1"时，返回 "special value 1", true，只允许被调用1次
	repo.EXPECT().Get("1").Return("special value 1", true).Times(1)
	// 获取，参数="2"时，返回 nil, false，也只允许被调用1次
	repo.EXPECT().Get("2").Return(nil, false).Times(1)
	// 获取，针对其他参数，返回key,true，可被调用任意次数
	repo.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string) (interface{}, bool) {
		return key, true
	}).AnyTimes()
	// 更新，打印更新日志，返回err=nil
	repo.EXPECT().Update(gomock.Any(), gomock.Any()).Do(
		func(key string, value interface{}) {
			t.Logf("更新成功,key=%s,value=%v", key, value)
		}).Return(nil).AnyTimes()

	service := Service{repo: repo}
	// 这里应该在控制台看到创建成功的输出
	err := service.CreateTest("123", "value123")
	ast.Nil(err)

	// 这里应该先看到Service获取数据成功的输出，然后是更新成功的输出
	err = service.UpdateGetTest("1", "value111")
	ast.Nil(err)

	// 这里应返回key不存在的错误信息
	err = service.UpdateGetTest("2", "value111")
	ast.Equal(ErrKeyNotExist, err)

	// 这里应该先看到Service获取数据成功的输出，然后是更新成功的输出
	err = service.UpdateGetTest("1", "value111")
	ast.Nil(err)
}
