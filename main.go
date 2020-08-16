package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"loger/watchdog"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	engine.Any("/",index)
	engine.GET("/log/:project", logger)
	engine.Run(":8088")
}


func index(context *gin.Context) {
	context.HTML(http.StatusOK,"index.tpl","")
}

/**
获取文件位置
 */
func getBase(project string)(string,error)  {
	project_map := map[string]string{"ieas": `D:\code\loger\ieas.api.web.`, "php_error": "/data/logs/php_error.log"}
	base_path, ok := project_map[project]
	if ok == false {
		return "", errors.New(project + "项目不存在")
	}
	dateStr := time.Now().Format("2006-01-02.log")
	return base_path+dateStr,nil
}

/**
项目日志
 */
func logger(ctx *gin.Context) {
	lineString := ctx.DefaultQuery("line", "5")
	line, e := strconv.Atoi(lineString)
	if e != nil {
		errorHandel(ctx,errors.New("请输入正确的行数"))
		return
	}

	project := ctx.DefaultQuery("project", "ieas")
	path, e := getBase(project)
	if e !=nil {
		errorHandel(ctx, e)
		return
	}

	projectLog, e := watchdog.NewProjectLog(path, line, "")
	if e != nil {
		errorHandel(ctx, e)
		return
	}
	logs, e := projectLog.GetLogs()
	if e != nil {
		errorHandel(ctx, e)
		return
	}
	ctx.HTML(http.StatusOK, "project_log.tpl",gin.H{
		"Title":"ieas",
		"Data":logs,
	})
}

/**

 */
func errorHandel(ctx *gin.Context,err error)  {
	ctx.HTML(http.StatusOK, "error.tpl", gin.H{
		"code": 10086,
		"msg": err.Error(),
	})
	return
}

