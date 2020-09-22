package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golf/lib"
	"net/http"
)

func main() {

	r := gin.New()

	r.Handle(http.MethodGet, "news/:id", func(context *gin.Context) {
		// 1.从对象池获取新闻缓存
		cache := lib.NewsCache()
		defer lib.ReleaseNewsCache(cache)
		// 2.获取参数，设置DBGetter
		id := context.Param("id")
		cache.Getter = lib.NewsDbGetter(id) // 当缓存不存在时，从此处获取数据
		// 3.从缓存输出(如果没有缓存，上面的DBGetter会被调用)
		context.Header("Content-type", "application/json")
		context.String(http.StatusOK, cache.GetCache(fmt.Sprintf("news:%s", id)).(string))
	})

	_ = r.Run()
}
