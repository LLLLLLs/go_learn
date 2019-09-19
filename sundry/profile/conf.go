// Time        : 2019/09/05
// Description :

package profile

import "context"

type ConfigCenter interface {
	// 获取指定配置表(集合)
	GetTable(ctx context.Context, model interface{}) Table
	// 更新新版本配置
	Update()
}

type Table interface {
	// 根据预定义索引获取指定配置，索引顺序需与模型Tag定义一致
	// 如对于key_value表，v,err := Get("XXX");value := v.(KeyValue)就可获取到key="XXX"的配置项
	Get(index ...interface{}) (interface{}, error)
	// 获取所有配置项
	// 如对于key_value表，all,err:=All();list = all.([]KeyValue)就可以切片形式获取到所有key_value配置
	All() (interface{}, error)
}
