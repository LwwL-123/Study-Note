package main

import (
	"fmt"
)

type SlashDict struct {
	dict map[byte]*SlashDict
	end  bool
	data int
}

// Add 添加
func (sd *SlashDict) Add(key string, value int) {
	// 如果key为空，说明已经到末尾
	if len(key) == 0 {
		sd.end = true
		sd.data = value
		return
	}

	if key[0] == '/' {
		key = key[1:]
	}
	if sd.dict == nil {
		sd.dict = make(map[byte]*SlashDict)
	}

	c := key[0]
	if _, ok := sd.dict[c]; !ok {
		sd.dict[c] = &SlashDict{}
	}
	sd.dict[c].Add(key[1:], value)
}

// Get 获取
func (sd *SlashDict) Get(key string) (int, error) {
	if len(key) == 0 {
		return sd.data, nil
	}

	if key[0] == '/' {
		key = key[1:]
	}

	_, ok := sd.dict[key[0]]
	if !ok {
		return 0, fmt.Errorf("key %s not found", key)
	}

	return sd.dict[key[0]].Get(key[1:])
}

// Pop 弹出
func (sd *SlashDict) Pop(key string) error {
	if len(key) == 0 {
		return fmt.Errorf("key %s not found", key)
	}

	if key[0] == '/' {
		key = key[1:]
	}

	now, ok := sd.dict[key[0]]
	if !ok {
		return fmt.Errorf("key %s not found", key)
	}
	// 如果现在是结尾
	if now.end {
		delete(sd.dict, key[0])
		return nil
	}

	return sd.dict[key[0]].Pop(key[1:])
}

var keys []string

func (sd *SlashDict) deep_keys() []string {
	// dfs深度优先遍历
	sd.dfs([]byte{})
	res := []string{}
	for _, v := range keys {
		tmp := ""
		for _, vv := range v {
			tmp += string(vv)
			tmp += "/"
		}
		res = append(res, tmp[:len(tmp)-1])
	}
	return res
}

func (sd *SlashDict) dfs(path []byte) {
	for k, v := range sd.dict {
		path = append(path, k)
		if v.end {
			keys = append(keys, string(path))
		}
		v.dfs(path)
		path = path[:len(path)-1]
	}
}

func main() {
	sd := &SlashDict{}
	sd.Add("a/b/c", 6)
	sd.Add("a/b/f", 3)
	sd.Add("a/b/f/e", 3)

	// 获取
	tmp1, err := sd.Get("a/b/c")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tmp1)
	}
	// 弹出
	sd.Pop("a/b/c")
	tmp2, err := sd.Get("a/b/c")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tmp2)
	}

	// 深度遍历
	str := sd.deep_keys()
	fmt.Println(str)
}
