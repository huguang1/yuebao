package libs

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Configs struct { // 结构体
	config map[string]string  // 地图，类似于字典
	node   string
}

const MidStr = "-_-!"

var Conf *Configs  // 定义结构体指针变量

func init() {
	Conf = new(Configs)  // new方法，返回一个新定义的结构体的指针
	Conf.LoadConfig("config/config.ini")  //
}

// 结构体的方法，也就是类方法，将配置文件中的内容写入到map中
func (conf *Configs) LoadConfig(path string) {
	conf.config = make(map[string]string)  // 初始化一个map
	file, err := os.Open(path)
	// fmt.Print(file, err)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// 读取配置文件中的内容，并将其放到map中
	buf := bufio.NewReader(file)
	for {
		lines, _, err := buf.ReadLine()
		line := strings.TrimSpace(string(lines))
		if err != nil {
			//读完最后一行退出
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		// 处理注释
		if strings.Index(line, "#") == 0 {
			continue
		}
		//如果是[xxx]
		n := strings.Index(line, "[")
		nl := strings.LastIndex(line, "]")

		if n > -1 && nl > -1 && nl > n+1 {
			conf.node = strings.TrimSpace(line[n+1 : nl])
			continue
		}
		if len(conf.node) == 0 || len(line) == 0 {
			continue
		}
		arr := strings.Split(line, "=")
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])
		newKey := conf.node + MidStr + key
		conf.config[newKey] = value
	}
}

// 读取配置文件中的内容
func (conf *Configs) Read(node, key string) string {
	key = node + MidStr + key
	if v, ok := conf.config[key]; !ok {
		return ""
	} else {
		return v
	}
}
