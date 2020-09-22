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
		id := context.Param("id")

		cache := lib.NewsCache()
		defer lib.ReleaseNewsCache(cache)

		cache.Getter = lib.NewsDbGetter(id)
		context.Header("Content-type", "application/json")
		context.String(http.StatusOK, cache.GetCache(fmt.Sprintf("news:%s", id)).(string))
	})

	_ = r.Run()
}
