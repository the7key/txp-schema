package main

import (
	"github.com/GIT_USER_ID/GIT_REPO_ID/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

    //todo: handle the error!
	c, _ := handlers.NewContainer()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// CreateCall - 一斉架電を登録する
	e.POST("/jians/:jianId/calls", c.CreateCall)

	// CreateInst - 施設を登録する
	e.POST("/insts", c.CreateInst)

	// GetInst - 施設情報を返す
	e.GET("/insts/:instId", c.GetInst)

	// GetInsts - 施設一覧を返す
	e.GET("/insts", c.GetInsts)

	// CreateInstGroup - 施設グループを作成する
	e.POST("/inst-groups", c.CreateInstGroup)

	// GetInstGroup - 施設グループ情報を返す
	e.GET("/inst-groups/:instGroupId", c.GetInstGroup)

	// GetInstGroups - 施設グループ一覧を返す
	e.GET("/inst-groups", c.GetInstGroups)

	// CreateJian - 事案を作成する
	e.POST("/jians", c.CreateJian)

	// GetJians - 事案一覧を返す
	e.GET("/jians", c.GetJians)


	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}