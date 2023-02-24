// # Copyright 2023 LightTree <alwaysday1@qq.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/alwaysday1/flashBox.

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"net/http"
	"strconv"
	"time"
)

func curd() {
	dsn := "root:15110307050@tcp(127.0.0.1:3306)/flash-box?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	fmt.Println("db=", db)
	fmt.Println("err=", err)

	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	type List struct {
		gorm.Model
		Name  string `gorm:"type:varchar(20);not null" json:"name" binding:"required"`
		Phone string `gorm:"type:varchar(20);not null" json:"phone" binding:"required"`
	}
	db.AutoMigrate(&List{})
	r := gin.Default()
	v1 := r.Group("api")

	v1.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"msg": "run success",
		})
	})
	// 增
	v1.POST("/user/add", func(context *gin.Context) {
		var data List
		err := context.ShouldBindJSON(&data)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "添加失败",
				"data": err.Error(),
				"code": "400",
			})
		} else {
			db.Create(&data)
			context.JSON(http.StatusOK, gin.H{
				"msg":  "添加成功",
				"data": data,
				"code": "200",
			})
		}
	})
	// 删
	v1.DELETE("/user/delete/:id", func(context *gin.Context) {
		var data []List
		id := context.Param("id")
		db.Where("id =?", id).Find(&data)
		if len(data) == 0 {
			context.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "not found id",
			})
		} else {
			db.Delete(&data)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "delete id success",
			})
		}
	})
	// 改
	v1.PUT("/user/update/:id", func(context *gin.Context) {
		var data List
		id := context.Param("id")
		db.Select("id").Where("id =?", id).Find(&data)
		if data.ID == 0 {
			context.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  "can not find userid",
			})
		} else {
			if err := context.ShouldBindJSON(&data); err != nil {
				context.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":  err.Error(),
				})
				return
			}
			db.Where("id=?", id).Updates(&data)
			context.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "change success",
			})

		}

	})
	//查，不分页
	v1.GET("/user/list/:name", func(context *gin.Context) {
		name := context.Param("name")
		var dataList []List
		db.Where("name=?", name).Find(&dataList)
		if len(dataList) == 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "no data",
				"code": "400",
				"data": gin.H{},
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "query success",
				"code": "200",
				"data": dataList,
			})
		}
	})
	// 查，分页
	v1.GET("/user/list", func(context *gin.Context) {
		var dataList []List
		pageSize, _ := strconv.Atoi(context.Query("pageSize"))
		pageNum, _ := strconv.Atoi(context.Query("pageNum"))

		if pageSize == 0 {
			pageSize = -1
		}
		if pageNum == 0 {
			pageNum = -1
		}
		offsetVal := (pageNum - 1) * pageSize
		if pageNum == -1 && pageSize == -1 {
			offsetVal = -1
		}
		var total int64
		db.Model(dataList).Count(&total).Limit(pageSize).Offset(offsetVal).Find(&dataList)

		if len(dataList) == 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "未查到数据",
				"code": 400,
				"data": gin.H{},
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg":  "查到数据",
				"code": 200,
				"data": gin.H{
					"list":     dataList,
					"total":    total,
					"pageNum":  pageNum,
					"pageSize": pageSize,
				},
			})

		}
	})
	v1.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "pong",
			"code": "200",
		})
	})
	PORT := "3000"
	r.Run(":" + PORT)
}

func main() {
	//curd()
}
