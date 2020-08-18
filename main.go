package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"loger/tools"
	"loger/watchdog"
)

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("template/*")
	engine.Any("/", index)
	engine.GET("/log/:project", logger)
	engine.GET("/cap", Cap)
	engine.Run(":8088")
}

func Cap(ctx *gin.Context) {
	host := tools.Env("redis.HOST")
	port := tools.Env("redis.PORT")
	config := tools.EnvSection("sms_prefix")
	scanner, err := tools.NewScanner(host, port)
	if err != nil {
		errorHandel(ctx, err)
		return
	}
	codes := scanner.ScanCode(config)
	ctx.HTML(http.StatusOK, "codes.tpl", gin.H{
		"codes": codes,
	})
}

func index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tpl", "")
}

/**
获取文件位置
*/
func getBase(project string) (string, error) {
	project = strings.ToUpper(project)
	base_path := tools.Env("log_path." + project)
	if base_path == "" {
		return "", errors.New(project + "项目地址未配置")
	}
	dateStr := time.Now().Format("2006-01-02.log")
	file := base_path + dateStr
	_, e := os.Stat(file)
	if e != nil {
		return "", errors.New(project + "文件不存在:" + e.Error())
	}
	return base_path + dateStr, nil
}

/**
项目日志
*/
func logger(ctx *gin.Context) {
	lineString := ctx.DefaultQuery("line", "5")
	line, e := strconv.Atoi(lineString)
	if e != nil {
		errorHandel(ctx, errors.New("请输入正确的行数"))
		return
	}

	project := ctx.Param("project")
	path, e := getBase(project)
	if e != nil {
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
	ctx.HTML(http.StatusOK, "project_log.tpl", gin.H{
		"Title": "ieas",
		"Data":  logs,
	})
}

/**

 */
func errorHandel(ctx *gin.Context, err error) {
	ctx.HTML(http.StatusOK, "error.tpl", gin.H{
		"code": 10086,
		"msg":  err.Error(),
	})
	return
}
