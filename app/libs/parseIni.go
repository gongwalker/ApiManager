/**********************************************
** @Des: parse config file
** @Author: gongwen [https://www.gwalker.cn]
** @Date:   2018-10-03 15:42:43
** @Last Modified by:   gongwen
** @Last Modified time: 2019-01-29 11:49:17
***********************************************/
package libs

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// 两个配置key map内部分隔符(可自定义)
const seq = "#-V-#"

var cfg = &ini{make(map[string]string)}

// 配置存储结构体
type ini struct {
	cfg map[string]string
}

// 指定key读配置文件中的某项值
func (c *ini) GetConfig(key string) (string, error) {
	key = strings.Replace(key, ".", seq, 1)
	val, ok := c.cfg[key]
	if ok {
		return val, nil
	} else {
		return val, fmt.Errorf("%s", "指定的key不存在")
	}
}

// GetConfig 方法别名
func (c *ini) GetConfigToString(key string) (string, error) {
	return c.GetConfig(key)
}

// 读配置文件的某项key,且转为Int型
func (c *ini) GetConfigToInt(key string) (int, error) {
	val, error := c.GetConfig(key)
	if error == nil {
		intNum, err := strconv.Atoi(val)
		if err != nil {
			return 0, nil
		}
		return intNum, nil
	}
	return 0, error
}

// 读配置文件的某项key,且转为Int64型
func (c *ini) GetConfigToInt64(key string) (int64, error) {
	val, error := c.GetConfig(key)
	if error == nil {
		intNum, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return 0, nil
		}
		return intNum, nil
	}
	return 0, error
}

// 读配置文件的某项key,且转为bool型
func (c *ini) GetConfigToBool(key string) (bool, error) {
	val, error := c.GetConfig(key)
	if error == nil {
		if val == "false" || val == "" {
			return false, nil
		} else if val != "0" {
			return true, nil
		}
	}
	return false, error
}

// 读配置文件的某项key,且转为Float64型
func (c *ini) GetConfigToFloat64(key string) (float64, error) {
	val, error := c.GetConfig(key)
	if error == nil {
		floatNum, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, nil
		}
		return floatNum, nil
	}
	return 0, error
}

// 解析ini配置文件形成map键值对
func ReadIniFile(filePath string) (*ini, error) {
	var scope, key string

	str, err := ioutil.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	arr := strings.Split(string(str), "\n")

	for _, v := range arr {

		row := strings.TrimSpace(v)
		l := len(row)

		// 空行忽略
		if l == 0 {
			continue
		}

		// 注释行忽略
		if row[0:1] == ";" || row[0:1] == "#" || l >= 2 && row[0:2] == "//" {
			continue
		}

		// 如果行小于三个字符,则忽略
		if l < 3 {
			continue
		}

		// 查看是否为范围段
		if row[0:1] == "[" && row[l-1:l] == "]" {
			scope = row[1 : l-1]
		}

		// 解析键值对
		isHasEqual := strings.Index(row, "=")
		if isHasEqual > 1 {
			rKey := strings.TrimSpace(row[0:isHasEqual])
			rVal := strings.TrimSpace(row[isHasEqual+1 : l])
			if scope != "" {
				key = scope + seq + rKey
			} else {
				key = rKey
			}

			rLen := len(rVal)
			if rLen > 1 {
				if rVal[0:1] == "\"" && rVal[len(rVal)-1:rLen] == "\"" {
					// rVal如果两边都有双引号，去掉双引号
					rVal = rVal[1 : rLen-1]
				} else if rVal[0:1] == "'" && rVal[len(rVal)-1:rLen] == "'" {
					// rVal如果两边都有单号引号，去掉单引号
					rVal = rVal[1 : rLen-1]
				}
			}
			cfg.cfg[key] = rVal
		}
	}
	return cfg, nil
}
