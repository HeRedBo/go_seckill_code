package main

import (
	"ImoocIrisShop/common"
	"ImoocIrisShop/encrypt"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var hostArray = []string{"127.0.0.1", "127.0.0.1"}

var losthost = "127.0.0.1"

var port = "8081"

type AccessControl struct {
	sourceArray map[int]interface{}
	sync.RWMutex
}

// 创建全局变量
var accessControl = &AccessControl{
	sourceArray: make(map[int]interface{}),
}

// 获取定制的数据
func (m *AccessControl) GetNameRecord(uid int) interface{} {
	m.RWMutex.RLock()
	defer m.RWMutex.RUnlock()
	data := m.sourceArray[uid]
	return data
}

// 设置记录
func (m *AccessControl) SetNewRecord(uid int) {
	m.RWMutex.Lock()
	m.sourceArray[uid] = "hello Index123"
	m.RWMutex.Unlock()
}

// 获取本机map 并且处理业务逻，返回结果类型为 bool 类型
func (m *AccessControl) GetDataFormMap(uid string) (isOk bool) {
	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return false
	}
	data := m.GetNameRecord(uidInt)
	if data != nil {
		return true
	}
	return
}

// 获取其他节点处理结果
func GetDataFormOtherMap(host string, request *http.Request) bool {
	uidPre, err := request.Cookie("uid")
	if err != nil {
		return false
	}
	uidSign, err := request.Cookie("sign")
	if err != nil {
		return false
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://"+host+":"+port+"/check", nil)
	if err != nil {
		return false
	}
	// 手动指定，排查多余的cookie
	cookieUid := &http.Cookie{Name: "uid", Value: uidPre.Value, Path: "/"}
	cookieSign := &http.Cookie{Name: "sign", Value: uidSign.Value, Path: "/"}
	// 添加cookie 到模拟的请求中
	req.AddCookie(cookieUid)
	req.AddCookie(cookieSign)

	// 获取返回结果
	response, err := client.Do(req)
	if err != nil {
		return false
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return false
	}
	// 判断状态
	if response.StatusCode == 200 {
		if string(body) == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}

// 执行正常的业务逻辑
func Check(w http.ResponseWriter, r *http.Request) {
	// 执行正常的业务逻辑
	fmt.Println("执行check")
}

//统一验证拦截器，每个接口都要提前验证
func Auth(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("执行验证")
	return nil
}

// 用户身份验证
func CheckUserInfo(r *http.Request) error {
	uidCookie, err := r.Cookie("uid")
	if err != nil {
		return errors.New("用户UID Cookie 获取失败！")
	}
	// 获取用户加密串
	signCookie, err := r.Cookie("sign")
	if err != nil {
		return errors.New("用户加密串 Cookie 获取失败！")
	}
	signByte, err := encrypt.DePwdCode(signCookie.Value)
	if err != nil {
		return errors.New("加密串已被篡改！")
	}
	if checkInfo(uidCookie.Value, string(signByte)) {
		return nil
	}
	return errors.New("身份校验失败！")
	//return nil
}

// 自定义逻辑判断
func checkInfo(checkStr string, signStr string) bool {
	if checkStr == signStr {
		return true
	}
	return false
}

func main() {

	hashConsistent := common.NewConsistent()
	for _, v := range hostArray {
		hashConsistent.Add(v)
	}

	filter := common.NewFilter()

	filter.RegisterFilterUri("/check", Auth)

	http.HandleFunc("/check", filter.Handle(Check))
	//启动服务
	http.ListenAndServe(":8083", nil)
}
