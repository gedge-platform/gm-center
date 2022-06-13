package controller

import (
	"net/http"

	"gmc_api_gateway/app/common"

	"gmc_api_gateway/app/model"

	"github.com/labstack/echo/v4"
)

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
