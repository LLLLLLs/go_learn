/*
Author : Haoyuan Liu
Time   : 2018/6/25
*/
package dcapi

import (
	"math"
	"time"
	"errors"
)

const abortIndex int8 = math.MaxInt8 / 2

//HandlerFunc 为责任链中的回调函数
type HandlerFunc func(Context)

//HandlersChain 为责任链
type HandlersChain []HandlerFunc

//Last 返回责任链中的最后一个handler，其为主handler
func (h HandlersChain) Last() HandlerFunc {
	if length := len(h); length > 0 {
		return h[length-1]
	}
	return nil
}

//Next 只能用于中间件中，
//它将执行责任链中的下一个handler
func Next(c Context) {
	c.addIndex()
	handlers := c.handlers()
	for s := int8(len(handlers)); c.index() < s; c.addIndex() {
		handlers[c.index()](c)
	}
}

//Context 的基本功能有：
// 1. 保存请求参数、响应内容
// 2. 执行责任链（handler链）
// 3. 缓存上下文中使用的键值
type Context interface {
	//请求参数，可为任意类型，使用时自己做类型转换
	Params() interface{}
	//设置请求参数，由于请求参数是只读的，只能设置一次
	SetParams(interface{})
	//响应
	Resp() (Response, error)
	//设置响应
	SetResp(Response)

	index() int8
	addIndex()
	handlers() HandlersChain

	//*注意，每个Context实现必须重新实现一遍该方法
	Next()
	//最后一个Handler函数
	Last() HandlerFunc
	//责任链是否中断
	IsAborted() bool
	//中断责任链
	Abort()
	AbortWithError(err error)

	//上下文内的存取器
	Store() *Store

	//兼容原生context

	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

type context struct {
	idx    int8
	params interface{}
	resp   Response
	chain  HandlersChain
	store  Store
}

func (c *context) Store() *Store {
	return &c.store
}

/************************************/
/********** Params & Resp ***********/
/************************************/

func (c *context) Params() interface{} {
	return c.params
}

func (c *context) SetParams(p interface{}) {
	if c.params == nil {
		c.params = p
	}
}

func (c *context) Resp() (Response, error) {
	if c.resp == nil {
		return nil, errors.New("must SetResp first")
	}
	return c.resp, nil
}

func (c *context) SetResp(resp Response) {
	c.resp = resp
}

/************************************/
/*********** HandlerChain ***********/
/************************************/

func (c *context) Next() {
	Next(c)
}

// Last returns the main handler.
func (c *context) Last() HandlerFunc {
	return c.handlers().Last()
}

// IsAborted returns true if the current context was aborted.
func (c *context) IsAborted() bool {
	return c.idx >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current params is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this params are not called.
func (c *context) Abort() {
	c.idx = abortIndex
}

// AbortWithError calls `AbortWithStatus()` and `SetError()` internally.
// This method stops the chain, writes the status code and pushes the specified error to `c.Errors`.
func (c *context) AbortWithError(err error) {
	c.resp.SetError(err)
}

func (c *context) index() int8 {
	return c.idx
}

func (c *context) addIndex() {
	c.idx++
}

func (c *context) handlers() HandlersChain {
	return c.chain
}

/************************************/
/***** GOLANG.ORG/X/NET/CONTEXT *****/
/************************************/

func (c *context) Deadline() (deadline time.Time, ok bool) {
	return
}

func (c *context) Done() <-chan struct{} {
	return nil
}

func (c *context) Err() error {
	return nil
}

func (c *context) Value(key interface{}) interface{} {
	if keyAsString, ok := key.(string); ok {
		val := c.store.Get(keyAsString)
		return val
	}
	return nil
}

func NewContext(manager *Manager) Context {
	return &context{
		idx:   0,
		chain: manager.Handlers(),
	}
}
