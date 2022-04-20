package cache_service

import (
	"strconv"
	"strings"
)

type Tag struct {
	ID   int
	Name string
}

// 获取标签缓存key的方法
func (t *Tag) GetTagsKey() string {
	keys := []string{
		cacheTag,
		"List",
	}
	if t.ID > 0 {
		keys = append(keys, strconv.Itoa(t.ID))
	}
	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	return strings.Join(keys, "_")
}
