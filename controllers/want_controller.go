package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/li-ming-lei/AccountShareC2C/model"
	"github.com/li-ming-lei/AccountShareC2C/service"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func OptionsRequest(c *gin.Context) {
	c.JSON(200, "")
}

func AddWant(c *gin.Context) {
	var want model.Want
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&want)
	want.Id = bson.NewObjectId()
	now := time.Now()
	want.CreateTime = &now
	if err != nil {
		fmt.Println("bind plan from request error")
		c.JSON(400, model.Response{Code: 10001, Msg: err.Error()})
		return
	}
	err = service.AddWant(want)
	if err != nil {
		fmt.Println("add plan to db error")
		c.JSON(500, model.Response{Code: 100002, Msg: err.Error()})
		return
	}
	c.JSON(200, (&model.Response{Code: 0, Msg: "success"}).AddContent("want", want))

}

func ListWants(c *gin.Context) {
	var want model.Want
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&want)
	if err != nil {
		fmt.Printf("bind forsell from request error, err: %+v", err)
		c.JSON(400, model.Response{Code: 10001, Msg: err.Error()})
		return
	}
	wants, err := service.ListWants(want)
	if err != nil {
		c.JSON(500, &model.Response{Code: 10001, Msg: "failed"})
		return
	}
	if wants == nil {
		wants = make([]model.Want, 0)
	}
	c.JSON(200, (&model.Response{Code: 0, Msg: "success"}).AddContent("wants", wants))
}

func DeleteWant(c *gin.Context) {
	var want model.Want
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&want)
	want.IsDel = false
	if err != nil {
		fmt.Println("bind want from request error")
		c.JSON(400, model.Response{Code: 10001, Msg: err.Error()})
		return
	}
	if !want.Id.Valid() {
		fmt.Println("want id is not specified")
		c.JSON(400, model.Response{Code: 10001, Msg: "want id is not specified"})
		return
	}
	err = service.DeleteWant(want.Id)
	if err != nil {
		c.JSON(500, &model.Response{Code: 10001, Msg: "failed"})
		return
	}
	c.JSON(200, &model.Response{Code: 0, Msg: "success"})
}
