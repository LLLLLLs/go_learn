/*
Author : Haoyuan Liu
Time   : 2018/5/17
*/
package env

import (
	"arthur/utils/osutils"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	GAME_NAME      = "arthur"
	StatCache_Name = "ProCenter"
)

var (
	HTTP_ADDR = osutils.GetEnvWithDefault("HTTP_ADDR", "0.0.0.0:20000")
	ZK_HOST   = getZkHost()

	LOG_LEVEL = osutils.GetEnvWithDefault("LOG_LEVEL", "debug")

	// ZK 根节点
	ZK_ROOT = osutils.GetEnvWithDefault("ZK_PATH", "/arthur_dev")
	// ZK auth字符串
	ZK_AUTH = getZkAuth()

	DOC_HOST     = osutils.GetEnvWithDefault("DOC_SERVER_HOST", "")
	DOC_PORT     = 22
	DOC_USER     = "root"
	DOC_PASSWORD = osutils.GetEnvWithDefault("DOC_SERVER_PASSWD", "")

	REDIS_KEY_SEP = "::"
)

// 获取zk hosts
func getZkHost() []string {
	hosts := osutils.GetEnvWithDefault("ZK_HOSTS", "10.46.1.15:2181,10.46.211.190:2181,10.46.140.8:2181")
	return strings.Split(hosts, ",")
}

// getZkAuth 获取加密后的ZkAuth字符串
func getZkAuth() string {
	authString := osutils.GetEnvWithDefault("ZK_AUTH", "user:password")
	auth := strings.Split(authString, ":")
	bAuth := md5.Sum([]byte(auth[1]))
	return auth[0] + ":" + string(bAuth[:])
}

func ProjectRoot() string {
	name := GAME_NAME
	p := os.Getenv("PROJECT_ROOT")
	if p != "" {
		return p
	}
	gopath := os.Getenv("GOPATH")
	mulPath := make([]string, 0)
	//check in windows
	if runtime.GOOS == "windows" {
		mulPath = strings.Split(gopath, ";")
	} else {
		//check in linux
		mulPath = strings.Split(gopath, ":")
	}

	for _, p := range mulPath {
		src := filepath.Join(p, "src")
		dirs := ListDir(src, false, true)
		if findInList(name, dirs) {
			return filepath.Join(src, name)
		}
	}
	panic(fmt.Sprintf("project named %s not found", name))
}

func findInList(name string, list []string) bool {
	for _, s := range list {
		if s == name {
			return true
		}
	}
	return false
}

// ListDir list directories and files in fpath
func ListDir(fpath string, fullPath bool, listDir bool) []string {
	files, err := ioutil.ReadDir(fpath)
	dirs := make([]string, 0)
	fileName := ""
	if err != nil {
		log.Printf("list error path %s", fpath)
		log.Fatal(err)
	}
	for _, f := range files {
		if fullPath {
			fileName = path.Join(fpath, f.Name())
		} else {
			fileName = f.Name()
		}
		if f.IsDir() == listDir {
			dirs = append(dirs, fileName)
		}
	}
	return dirs
}
