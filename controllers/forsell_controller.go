package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/li-ming-lei/AccountShareC2C/model"
	"github.com/li-ming-lei/AccountShareC2C/service"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func AddForSell(c *gin.Context) {
	var forsell model.ForSell
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&forsell)
	forsell.Id = bson.NewObjectId()
	now := time.Now()
	forsell.CreateTime = &now
	if err != nil {
		fmt.Println("bind plan from request error")
		c.JSON(400, model.Response{Code: 10001, Msg: err.Error()})
		return
	}
	err = service.AddForsell(forsell)
	if err != nil {
		fmt.Println("add plan to db error")
		c.JSON(500, model.Response{Code: 100002, Msg: err.Error()})
		return
	}
	c.JSON(200, (&model.Response{Code: 0, Msg: "success"}).AddContent("forsell", forsell))
}

func ListForSells(c *gin.Context) {
	var forsell model.ForSell
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&forsell)
	if err != nil {
		fmt.Printf("bind forsell from request error, err: %+v", err)
		c.JSON(400, model.Response{Code: 10001, Msg: err.Error()})
		return
	}
	forsells, err := service.ListForsells(forsell)
	if err != nil {
		c.JSON(500, &model.Response{Code: 10001, Msg: "failed"})
		return
	}
	if forsells == nil {
		forsells = make([]model.ForSell, 0)
	}
	c.JSON(200, (&model.Response{Code: 0, Msg: "success"}).AddContent("forsells", forsells))
}

func DeleteForSell(c *gin.Context) {
	var forsell model.ForSell
	var err error
	// This will infer what binder to use depending on the content-type header.
	err = c.Bind(&forsell)
	forsell.IsDel = false
	if err != nil {
		fmt.Println("bind forsell from request error")
		c.JSON(400, model.Response{Code: 10001, Msg: err.Error()})
		return
	}
	if !forsell.Id.Valid() {
		fmt.Println("forsell id is not specified")
		c.JSON(400, model.Response{Code: 10001, Msg: "forsell id is not specified"})
		return
	}
	err = service.DeleteForsell(forsell.Id)
	if err != nil {
		c.JSON(500, &model.Response{Code: 10001, Msg: "failed"})
		return
	}
	c.JSON(200, &model.Response{Code: 0, Msg: "success"})
}
