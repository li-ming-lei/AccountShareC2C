package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/li-ming-lei/AccountShareC2C/conf"
	"github.com/li-ming-lei/AccountShareC2C/model"
	utils "github.com/li-ming-lei/AccountShareC2C/util"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {
	var user model.User
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&user)
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", conf.GlobalConf.AppId, conf.GlobalConf.AppSecret, user.Code)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, model.Response{Code: 100002, Msg: err.Error()})
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	retUser := model.User{}
	err = utils.UnmarshalJson(body, &retUser)
	if err != nil {
		fmt.Printf("unmarshal user return err, %+v", err)
		c.JSON(500, model.Response{Code: 100002, Msg: err.Error()})
		return
	}
	c.JSON(200, (&model.Response{Code: 0, Msg: "success"}).AddContent("user", retUser))
}
