package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	httpNet "net/http"
	pb "xiuyiPro/api/turtle/v1"
)

func listTurtles(c *gin.Context) {
	var (
		req = new(pb.GetTurtleListReq)
	)
	err := c.ShouldBindQuery(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := tSvr.GetTurtleList(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, err)
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))

}

func setBatchTurtles(c *gin.Context) {
	var (
		req = new(pb.SetTurtleBatchReq)
	)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := tSvr.SetTurtleBatch(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, err)
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))
}
