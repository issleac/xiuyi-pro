package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	httpNet "net/http"
	pb "xiuyiPro/api/idiom/v1"
)

func getIdiom(c *gin.Context) {
	var (
		req = new(pb.GetIdiomReq)
	)
	err := c.ShouldBindQuery(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := iSvr.GetIdiom(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, BuildErrorResponse(c, 0)) //todo:errCode
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))

}

func setIdioms(c *gin.Context) {
	var (
		req = new(pb.SetIdiomBatchReq)
	)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := iSvr.SetIdiomBatch(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, BuildErrorResponse(c, 0)) //todo:errCode
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))
}

func getRanking(c *gin.Context) {
	var (
		req = new(pb.GetRankingReq)
	)
	err := c.ShouldBindQuery(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := iSvr.GetRanking(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, BuildErrorResponse(c, 0)) //todo:errCode
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))
}

func updateRanking(c *gin.Context) {
	var (
		req = new(pb.UpdateRankingReq)
	)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := iSvr.UpdateRanking(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, BuildErrorResponse(c, 0)) //todo:errCode
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))
}

func startGame(c *gin.Context) {
	var (
		req = new(pb.StartAppReq)
	)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := iSvr.StartApp(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, BuildErrorResponse(c, 0)) //todo:errCode
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))
}

func endGame(c *gin.Context) {
	var (
		req = new(pb.EndAppReq)
	)
	err := c.ShouldBindJSON(req)
	if err != nil {
		log.Error(err)
		return
	}
	resp, err := iSvr.EndApp(c, req)
	if err != nil {
		c.JSON(httpNet.StatusOK, BuildErrorResponse(c, 0)) //todo:errCode
		return
	}
	c.JSON(httpNet.StatusOK, BuildDataResponse(c, resp))
}
