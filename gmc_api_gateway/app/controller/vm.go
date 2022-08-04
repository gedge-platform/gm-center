package controller

import (
	"encoding/json"
	"log"
	"fmt"

	"net/http"

	"gmc_api_gateway/app/common"

	"gmc_api_gateway/app/model"

	"github.com/labstack/echo/v4"
)

// GetCloudOS godoc
// @Summary Cloudos
// @Description get CloudOS
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Router /spider/cloudos [get]
func GetCloudOS(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "cloudos",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	cloudos := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": cloudos,
	})

}

// GetALLCredential godoc
// @Summary Credential
// @Description get ALLCredential
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Router /spider/credentials [get]
func GetALLCredential(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "credential",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	credential := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})

}

// GetCredential godoc
// @Summary Credential
// @Description get Credential
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Router /spider/credentials/{credentialName} [get]
// @Param credentialName path string true "Name of the credentials"
func GetCredential(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "credential",
		Name:   c.Param("credentialName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	credential := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})
}

func GetALLCredentialCount(c echo.Context)(err error) {

	params := model.PARAMS{
		Kind:   "credential",
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)

	var P model.CredentialCount
	json.Unmarshal([]byte(getData), &P)
	log.Printf("[#Credential Count] is %s", P.CredentialNames)

	return c.JSON(http.StatusOK, echo.Map{
		"credentialCnt": len(P.CredentialNames),
	})
}

// CreateCredential godoc
// @Summary Credential
// @Description get Credential
// @Param CredentialBody body string true "Credential Info Body"
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Router /spider/credentials [post]
func CreateCredential(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:   "credential",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	credential := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})
}

// DeleteCredential godoc
// @Summary Credential
// @Description get Credential
// @ApiImplicitParam
// @Accept  json
// @Produce  json
// @Security   Bearer
// @Router /spider/credentials/{credentialName} [delete]
// @Param credentialName path string true "Name of the credentials"
func DeleteCredential(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "credential",
		Name:   c.Param("credentialName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	credential := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": credential,
	})

}

func GetALLConnectionconfig(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "connectionconfig",
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)
	connectionconfig := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": connectionconfig,
	})
}

func GetConnectionconfig(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "connectionconfig",
		Name:   c.Param("configName"),
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)
	connectionconfig := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": connectionconfig,
	})
}

func CreateConnectionconfig(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "connectionconfig",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	connectionconfig := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": connectionconfig,
	})
}

func DeleteConnectionconfig(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "connectionconfig",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	connectionconfig := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": connectionconfig,
	})
}

func GetALLClouddriver(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "clouddriver",
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)
	clouddriver := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": clouddriver,
	})
}

func GetClouddriver(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "clouddriver",
		Name:   c.Param("clouddriverName"),
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)
	clouddriver := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": clouddriver,
	})
}

func RegisterClouddriver(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "clouddriver",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	clouddriver := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": clouddriver,
	})
}

func UnregisterClouddriver(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "clouddriver",
		Name:   c.Param("clouddriverName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	clouddriver := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": clouddriver,
	})
}

func GetALLCloudregion(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "cloudregion",
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)
	cloudregion := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": cloudregion,
	})
}

func GetCloudregion(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "cloudregion",
		Name:   c.Param("cloudregionName"),
		Method: c.Request().Method,
	}

	getData, err := common.DataRequest_spider(params)
	cloudregion := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": cloudregion,
	})
}

func RegisterCloudregion(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "cloudregion",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	cloudregion := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": cloudregion,
	})
}

func UnregisterCloudregion(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "cloudregion",
		Name:   c.Param("cloudregionName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	cloudregion := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": cloudregion,
	})
}

func VmControl(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "controlvm",
		Name:   c.Param("vmName"),
		Action: c.QueryParam("action"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})
}

func VmTerminate(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "controlvm",
		Name:   c.Param("vmName"),
		Action: c.QueryParam("action"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})
}

func GetALLVm(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vm",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})

}

func GetALLVmCount(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vm",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)

	var P model.VMCount
	json.Unmarshal([]byte(getData), &P)

	return c.JSON(http.StatusOK, echo.Map{
		"VMCnt": len(P.VMCount),
	})

}

func GetVm(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("vmName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})

}

func CreateVm(c echo.Context) (err error) {
	params := model.PARAMS{
		Kind:   "vm",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)

	vm := StringToInterface(getData)
	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})
}

func DeleteVm(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vm",
		Name:   c.Param("vmName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vm := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vm,
	})

}

func GetALLVMStatus(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmstatus",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmstatus := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmstatus,
	})
}

func GetALLVMStatusCount(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmstatus",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)

	var P model.VMStatusCount
	json.Unmarshal([]byte(getData), &P)

	var running int = 0
	var suspended int = 0
	var failed int = 0

	for i := 0; i < len(P.Vmstatus); i++ {
		str := fmt.Sprintf("%v", P.Vmstatus[i])
		if(str == "{Running}"){
			running++
		}
		if(str == "{Suspended}"){
			suspended++
		}
		if(str == "{Failed}"){
			failed++
		}		
	}
	
	return c.JSON(http.StatusOK, echo.Map{
		"Running": running,
		"Stop": suspended,
		"Paused": failed,
	})
}

func GetVMStatus(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmstatus",
		Name:   c.Param("vmstatusName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmstatus := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmstatus,
	})
}

func GetALLVMSpec(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmspec",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmspec := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmspec,
	})
}

func GetVMSpec(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmspec",
		Name:   c.Param("vmspecName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmspec := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmspec,
	})
}

func GetALLVMOrgSpec(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmorgspec",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmorgspec := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmorgspec,
	})
}

func GetVMOrgSpec(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmorgspec",
		Name:   c.Param("vmspecName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmorgspec := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmorgspec,
	})
}

func GetALLVMImage(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmimage",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmimage := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmimage,
	})
}

func GetVMImage(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vmimage",
		Name:   c.Param("vmImageNameId"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vmimage := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vmimage,
	})
}

func GetALLVPC(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vpc",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vpc := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vpc,
	})
}

func GetVPC(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vpc",
		Name:   c.Param("vpcName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vpc := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vpc,
	})
}

func CreateVPC(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vpc",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vpc := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vpc,
	})
}

func DeleteVPC(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "vpc",
		Name:   c.Param("vpcName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	vpc := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": vpc,
	})
}

func GetALLSecurityGroup(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "securitygroup",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	securitygroup := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": securitygroup,
	})
}

func GetSecurityGroup(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "securitygroup",
		Name:   c.Param("securitygroupName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	securitygroup := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": securitygroup,
	})
}

func CreateSecurityGroup(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "securitygroup",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	securitygroup := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": securitygroup,
	})
}

func DeleteSecurityGroup(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "securitygroup",
		Name:   c.Param("securitygroupName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	securitygroup := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": securitygroup,
	})
}

func RegisterSecurityGroup(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "regsecuritygroup",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	regsecuritygroup := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": regsecuritygroup,
	})
}

func UnregisterSecurityGroup(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "regsecuritygroup",
		Name:   c.Param("securitygroupName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	regsecuritygroup := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": regsecuritygroup,
	})
}

func GetALLKeypair(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "keypair",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	keypair := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": keypair,
	})
}

func GetKeypair(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "keypair",
		Name:   c.Param("keypairName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	keypair := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": keypair,
	})
}

func CreateKeypair(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "keypair",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	keypair := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": keypair,
	})
}

func DeleteKeypair(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "keypair",
		Name:   c.Param("keypairName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	keypair := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": keypair,
	})
}

func RegisterKeypair(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "regkeypair",
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	regkeypair := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": regkeypair,
	})
}

func UnregisterKeypair(c echo.Context) (err error) {

	params := model.PARAMS{
		Kind:   "regkeypair",
		Name:   c.Param("keypairName"),
		Method: c.Request().Method,
		Body:   common.ResponseBody_spider(c.Request().Body),
	}

	getData, err := common.DataRequest_spider(params)
	regkeypair := StringToInterface(getData)

	return c.JSON(http.StatusOK, echo.Map{
		"data": regkeypair,
	})
}
