package watchdog

import (
	"errors"
	"strconv"
	"strings"
)

type Project struct {
	Path    string
	Command string
	Seq     string
	Line    int
	Args    []string
	Data    []string
	ParseLog []*LogLine
	Error error
}

type LogLine struct {
	Time   string   `json:"time"`
	Level  string   `json:"Level"`
	Params string   `json:"Params,omitempty"`
	Trace  []string `json:"Trace,omitempty"`
}

func (this *Project) Parse(s string) (*LogLine) {
	line := &LogLine{}
	splitN := strings.SplitN(s, "]", 5)
	if len(splitN) != 5 {
		line.Params = "单行日志不等于5片,数据不完整";
		return line
	}
	// fmt.Printf("%s\n",s)
	// 第一部分(标准头)
	line.Level = strings.Replace(splitN[0], "[", "", 1)
	line.Time = strings.Split(strings.Replace(splitN[1], "[", "", 1), " ")[1]
	// 第二部分(参数,有普通文本,请求参数,函数调用栈)
	mixed := splitN[4]
	flag := false

	// 处理变量
	params_prefix := "     params => ";
	contains := strings.Contains(mixed, params_prefix)
	if contains {
		pars := strings.Replace(mixed, params_prefix, "", 1)
		line.Params = pars
		flag = true
	}
	// 处理调用栈
	trace_prefix := "     #0"
	contains = strings.Contains(mixed, trace_prefix)
	if contains {
		traces := strings.Split(mixed, "#")
		line.Trace = traces[1:len(traces)]
		flag = true
	}
	// 处理普通文本
	if flag == false {
		line.Trace = append(line.Trace, mixed)
	}
	return line
}

/**
初始化
*/
func NewProjectLog(path string, line int, keyword string) (*Project, error) {
	// 查找项目是否存在
	// 组装参数
	// 执行命令
	var args []string
	var num string
	num = strconv.Itoa(line)
	if keyword == "" {
		args = []string{"-n", num, path}
	} else {
		args = []string{"-n", num, path, "|", "grep", keyword}
	}

	project := &Project{
		Path:    path,
		Command: "tail",
		Seq:     "\n",
		Args:    args,
		Line:    line,
	}
	return project, nil
}

/**
获取日志,还未格式化
*/
func (this *Project) ReadLogs() (error) {
	s, err := execCommand(this.Command, this.Args)
	if err != nil {
		return errors.New("读取失败" + err.Error())
	}
	this.Data = strings.Split(s, this.Seq)
    return nil
}



/**
解析日志
 */
func (this *Project) format() {
	tempParseLogs := make([]*LogLine,0,this.Line)
	for _,logLine := range this.Data {
		// fmt.Printf("%v\n", logLine)
		// continue;
		tempParseLogs =  append(tempParseLogs, this.Parse(logLine))
	}
	// os.Exit(7)
	this.ParseLog = tempParseLogs
}

/**
 最后获取日志
 */
func (this *Project) GetLogs() ([]*LogLine,error) {
	e := this.ReadLogs()
	if e != nil {
		return nil,e
	}
	this.format()
	return this.ParseLog,nil
}
