package nodeListen

import (
	"bytes"
	"context"
	"strings"
	"sync"
)

// 确保每个不同服务地址zk的只有一个监听者
var (
	storer = &nodeListenStore{
		store: make(map[string]*NodeListenClient),
	}
)

type NodeListenClient struct {
	hostList []string
	zkPath   string
	zkHost   []string
	zkAuth   []string
	ctx      *context.Context
}

type nodeListenStore struct {
	sync.RWMutex
	store map[string]*NodeListenClient
}

func (store *nodeListenStore) getListenClient(source string) *NodeListenClient {
	store.RLock()
	client, ok := store.store[source]
	store.RUnlock()
	if ok {
		return client
	}
	return nil
}

func (store *nodeListenStore) addListenClient(source string, client *NodeListenClient) {
	store.Lock()
	store.store[source] = client
	store.Unlock()
}

func NewNodeListenClient(zkPath string, zkHost, zkAuth []string) *NodeListenClient {
	var (
		source bytes.Buffer
		client *NodeListenClient
		ctx    context.Context
		cancel context.CancelFunc
	)
	ctx, cancel = context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "Done", cancel)

	source.WriteString(zkPath)
	source.WriteString(strings.Join(zkHost, ""))
	source.WriteString(strings.Join(zkAuth, ""))
	if client = storer.getListenClient(source.String()); client != nil {
		return client
	}
	client = new(NodeListenClient)
	client.zkHost = zkHost
	client.zkPath = zkPath
	client.zkAuth = zkAuth
	client.ctx = &ctx
	client.startNodeListen()
	storer.addListenClient(source.String(), client)
	return client
}

func (z *NodeListenClient) setHostList(data []string) {
	z.hostList = data
}

func (z *NodeListenClient) GetHostList() []string {
	return z.hostList
}

func (z *NodeListenClient) startNodeListen() error {
	getHostList := func(data []string) {
		z.setHostList(data)
	}
	return StartNodeListen(z.ctx, z.zkHost, z.zkAuth, z.zkPath, getHostList)
}

func (z *NodeListenClient) nodeListenDone() bool {
	if cancel, ok := (*z.ctx).Value("Done").(context.CancelFunc); ok {
		cancel()
		return true
	} else {
		return false
	}
}