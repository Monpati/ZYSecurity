package utils

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

//https://www.jb51.net/article/151051.htm

func GetInt64(val int64) string {
	return strconv.FormatInt(val, 10)
}

func GetHeaderByName(ctx *gin.Context, key string) int64 {
	val := ctx.Request.Header.Get(key)
	if result, err := strconv.ParseInt(val, 10, 64); err == nil {
		return result
	} else {
		return 0
	}
}

func GetParamByName(ctx *gin.Context, key string) int64 {
	val := ctx.Param(key)
	if result, err := strconv.ParseInt(val, 10, 64); err == nil {
		return result
	} else {
		return 0
	}
}

//func GetPrimaryKey(db *gorm.DB, tableName string) string {
//	db.Exec(fmt.Sprintf("call proc_auto_increment_id('%s','%s',@result)", tableName, ""))
//	row := db.Raw("select @result").Row()
//	var id string
//	row.Scan(&id)
//	return id
//}

func GetMapIndex(list map[int]string, val string) int {
	for key, value := range list {
		if value == val {
			return key
		}
	}
	return 0
}

func GetMapValue(list map[int]string, val int) string {
	item, err := list[val]
	if err {
		return item
	} else {
		return list[0]
	}
}

// 定义GORM JSON格式
type JSON []byte

func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		errors.New("JSON: Invalid Scan")
	}
	*j = append((*j)[0:0], s...)
	return nil
}
func (m JSON) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}
func (m *JSON) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("JSON: Null Point Exception")
	}
	*m = append((*m)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}
