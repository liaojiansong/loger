package tools

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type item = map[string]string
type section = map[string]item

type Config struct {
	file string
	data section
}

/**
读取文件
*/
func (this *Config) readConfigFile() {
	file, e := os.Open(this.file)
	if e != nil {
		log.Fatal("文件不存在:" + e.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	// 初始化 map
	this.data = make(section)
	var section string
	for {
		readString, e := reader.ReadString('\n')
		if e == io.EOF {
			break
		}
		lastS := strings.TrimSpace(readString)
		flag, key, value := this.parse(lastS)

		if flag == 0 {
			_ = value
			_ = key
			continue
		}
		if flag == 1 {
			_ = value
			// 置换新的
			section = key
			continue
		}
		if flag == 2 {
			temp, ok := this.data[section]
			// 初始化 map
			if ok == false {
				temp = make(item)
			}
			temp[key] = value
			this.data[section] = temp
		}
	}
}

/**
解析
*/
func (this *Config) parse(s string) (flag int, key string, value string) {
	l := len(s)
	if l < 2 {
		return 0, "", ""
	}
	if s[0:1] == "[" && s[l-1:l] == "]" {
		return 1, s[1 : l-1], ""
	}

	split := strings.Split(s, "=")
	len_split := len(split)

	if len_split == 1 {
		return 2, split[0], ""
	}
	if len_split == 2 {
		return 2, split[0], split[1]
	}
	return 0, "", ""
}

/**
获取
*/
func (this *Config) get(name string) string {
	split2 := strings.Split(name, ".")
	if len(split2) != 2 {
		return ""
	}
	section, ok := this.data[split2[0]]
	if ok == false {
		return ""
	}
	item, ok := section[split2[1]]
	if ok == false {
		return ""
	}
	return item
}
func (this *Config) getSection(name string) item {
	split2 := strings.Split(name, ".")
	if len(split2) != 1 {
		return item{}
	}
	section, ok := this.data[split2[0]]
	if ok == false {
		return item{}
	}
	return section
}

func genC() *Config {
	c := &Config{
		file: `D:\data\loger\.env`,
	}
	c.readConfigFile()
	return c
}

/**
读取配置信息
*/
func Env(name string) string {
	c := genC()
	return c.get(name)
}

func EnvSection(name string) map[string]string {
	c := genC()
	return c.getSection(name)
}
