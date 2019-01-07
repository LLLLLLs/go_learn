//FROM:https://gitlab.dianchu.cc/go_chaos/service_discovery/tree/master/go_sd
package utils

import (
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type ServiceRegister struct {
	ZKClient *zk.Conn
	//ZKChan <-chan zk.Event
}

func NewServiceRegister(servers []string, watcher zk.EventCallback, logger zk.Logger) *ServiceRegister {
	var (
		sr  *ServiceRegister
		err error
	)
	sr = new(ServiceRegister)
	sr.ZKClient, _, err = zk.Connect(servers, time.Second, zk.WithEventCallback(watcher))
	if err != nil {
		panic(err)
	}
	if logger != nil {
		sr.ZKClient.SetLogger(logger)
	}
	return sr
}

func (sr *ServiceRegister) Close() {
	sr.ZKClient.Close()
}

func (sr *ServiceRegister) Lock(path string) (*zk.Lock, error) {
	var (
		lock *zk.Lock
		err  error
	)
	lock = zk.NewLock(sr.ZKClient, path, zk.WorldACL(zk.PermAll))
	err = lock.Lock()
	return lock, err
}

func (sr *ServiceRegister) Unlock(lock *zk.Lock) error {
	return lock.Unlock()
}

func (sr *ServiceRegister) Exists(path string) (bool, int32, <-chan zk.Event, error) {
	var (
		ok   bool
		stat *zk.Stat
		ech  <-chan zk.Event
		err  error
	)
	ok, stat, ech, err = sr.ZKClient.ExistsW(path)
	return ok, stat.Version, ech, err
}

// flag: 0 Normal Nodes, 1 Ephemeral Nodes, 2 Sequence Nodes
func (sr *ServiceRegister) Create(path string, data string, flag int32) (string, error) {
	return sr.ZKClient.Create(path, []byte(data), flag, zk.WorldACL(zk.PermAll))
}

func (sr *ServiceRegister) Delete(path string) error {
	return sr.ZKClient.Delete(path, -1)
}

func (sr *ServiceRegister) Get(path string) (string, int32, error) {
	var (
		data []byte
		stat *zk.Stat
		err  error
	)
	data, stat, err = sr.ZKClient.Get(path)
	if err != nil {
		return "", -2, err
	}
	return string(data), stat.Version, err
}

func (sr *ServiceRegister) GetW(path string) (string, int32, <-chan zk.Event, error) {
	var (
		data []byte
		stat *zk.Stat
		ech  <-chan zk.Event
		err  error
	)
	data, stat, ech, err = sr.ZKClient.GetW(path)
	if err != nil {
		return "", -2, nil, err
	}
	return string(data), stat.Version, ech, err
}

func (sr *ServiceRegister) GetChildren(path string) ([]string, int32, error) {
	var (
		children []string
		stat     *zk.Stat
		err      error
	)
	children, stat, err = sr.ZKClient.Children(path)
	if err != nil {
		return []string{}, -2, err
	}
	return children, stat.Version, err
}

func (sr *ServiceRegister) GetChildrenW(path string) ([]string, int32, <-chan zk.Event, error) {
	var (
		children []string
		stat     *zk.Stat
		err      error
		wCh      <-chan zk.Event
	)
	children, stat, wCh, err = sr.ZKClient.ChildrenW(path)
	if err != nil {
		return []string{}, -2, wCh, err
	}
	return children, stat.Version, wCh, err
}

func (sr *ServiceRegister) Set(path string, data string, version int32) (int32, error) {
	var (
		stat *zk.Stat
		err  error
	)
	stat, err = sr.ZKClient.Set(path, []byte(data), version)
	if err != nil {
		return -2, err
	}
	return stat.Version, nil
}
