// Time        : 2019/09/09
// Description :

package profile

import (
	"context"
	"fmt"
	"golearn/sundry/profile/errors"
	"golearn/sundry/profile/loader"
	"golearn/sundry/profile/model"
	"reflect"
	"strings"
)

//多版本配置中心
type confCenter struct {
	updateSignal chan struct{}
	versions     map[string]version
	loadData     loader.Loader
}

func New(loader loader.Loader) ConfigCenter {
	center := confCenter{
		updateSignal: make(chan struct{}),
		versions:     map[string]version{},
		loadData:     loader,
	}
	center.load()
	go center.refresh()
	return &center
}

func (vc confCenter) version(v string) version {
	versionConf, ok := vc.versions[v]
	if !ok {
		return version{
			lastErr: errors.ErrNoVersion,
		}
	}
	return versionConf
}

func (vc *confCenter) refresh() {
	for {
		select {
		// TODO 决定更新触发方式
		case <-vc.updateSignal:
			vc.load()
		}
	}
}

func (vc *confCenter) load() {
	dataMap, dataList, err := vc.loadData.Load()
	if err != nil {
		fmt.Printf("静态配置加载失败: %v", err)
		return
	}
	newVersion := version{tables: make(map[string]table)}
	for name := range dataMap {
		newVersion.tables[name] = newTable(dataMap[name], dataList[name])
	}
	v, err := newVersion.version()
	if err != nil {
		fmt.Printf("静态配置加载失败: %v", err)
		return
	}
	vc.versions[v] = newVersion
}

func (vc confCenter) GetTable(ctx context.Context, model interface{}) Table {
	ver, err := getVersion(ctx)
	if err != nil {
		return table{
			lastErr: err,
		}
	}
	return vc.version(ver).table(model)
}

func (vc confCenter) Update() {
	vc.updateSignal <- struct{}{}
}

// 单版本配置
type version struct {
	tables  map[string]table
	lastErr error
}

func (v version) table(model interface{}) table {
	if v.lastErr != nil {
		return table{
			lastErr: v.lastErr,
		}
	}
	name := reflect.ValueOf(model).Type().Name()
	t, ok := v.tables[name]
	if !ok {
		return table{
			lastErr: errors.ErrNoTable,
		}
	}
	return t
}

func (v version) version() (string, error) {
	if v.lastErr != nil {
		return "", v.lastErr
	}
	all, err := v.table(model.Version{}).All()
	if err != nil {
		return "", err
	}
	list := all.([]model.Version)
	if len(list) != 1 {
		return "", errors.ErrVersionConfig
	}
	return list[0].Version, nil
}

// 具体表配置
type table struct {
	data    map[string]interface{}
	list    interface{}
	lastErr error
}

func newTable(data map[string]interface{}, list interface{}) table {
	return table{
		data: data,
		list: list,
	}
}

func (t table) Get(index ...interface{}) (interface{}, error) {
	if t.lastErr != nil {
		return nil, t.lastErr
	}
	indexes := make([]string, len(index))
	for i := range indexes {
		indexes[i] = loader.IndexToString(index[i])
	}
	key := strings.Join(indexes, loader.Separator)
	conf, ok := t.data[key]
	if !ok {
		return nil, errors.ErrNoConfig
	}
	return conf, nil
}

func (t table) All() (interface{}, error) {
	return t.list, t.lastErr
}
