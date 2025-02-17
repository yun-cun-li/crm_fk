package api

import (
	"crm/models"
	"crm/response"
	"crm/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContractApi struct {
	contractService *service.ContractService
}

func NewContractApi() *ContractApi {
	contractApi := ContractApi{
		contractService: &service.ContractService{},
	}
	return &contractApi
}

// 创建合同
func (c *ContractApi) Create(context *gin.Context) {
	var param models.ContractCreateParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	fmt.Println(param)
	param.Creator = int64(uid)
	errCode := c.contractService.Create(&param)
	response.Result(errCode, nil, context)
}

// 更新合同
func (c *ContractApi) Update(context *gin.Context) {
	var param models.ContractUpdateParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := c.contractService.Update(&param)
	response.Result(errCode, nil, context)
}

// 删除合同
func (c *ContractApi) Delete(context *gin.Context) {
	var param models.ContractDeleteParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	errCode := c.contractService.Delete(&param)
	response.Result(errCode, nil, context)
}

// 查询合同列表
func (c *ContractApi) QueryList(context *gin.Context) {
	var param models.ContractQueryParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil || param.Page.PageNum <= 0 || param.Page.PageSize <= 0 {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	param.Creator = int64(uid)
	productList, rows, errCode := c.contractService.QueryList(&param)
	response.PageResult(errCode, productList, rows, context)
}

// 查询合同信息
func (c *ContractApi) QueryInfo(context *gin.Context) {
	var param models.ContractQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	productInfo, errCode := c.contractService.QueryInfo(&param)
	response.Result(errCode, productInfo, context)
}

// 编辑合同时，查询产品列表
func (p *ContractApi) QueryPlist(context *gin.Context) {
	var param models.ProductQueryParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		return
	}
	productList, errCode := p.contractService.QueryPlist(&param)
	response.Result(errCode, productList, context)
}
