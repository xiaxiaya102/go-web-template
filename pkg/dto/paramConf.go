package dto

type SaveParamConfRequest struct {
	Key        string `json:"key" binding:"required"`   // 配置键
	Param      string `json:"param" binding:"required"` // 参数名称
	ParamValue string `json:"param_value"`              // 参数值
}

type UpdateParamConfRequest struct {
	ID         int    `json:"id" binding:"required"`    // 配置ID
	Key        string `json:"key" binding:"required"`   // 配置键
	Param      string `json:"param" binding:"required"` // 参数名称
	ParamValue string `json:"param_value"`              // 参数值
}

type ParamConfPageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type ParamConfPageResponse struct {
	List  []ParamConfResponse `json:"list"`
	Total int64               `json:"total"`
}

type ParamConfResponse struct {
	ID         int    `json:"id"`          // 配置ID
	Key        string `json:"key"`         // 配置键
	Param      string `json:"param"`       // 参数名称
	ParamValue string `json:"param_value"` // 参数值
}

type ParamConfQueryRequest struct {
	Key   string `json:"key" form:"key"`     // 配置键（模糊查询）
	Param string `json:"param" form:"param"` // 参数名称（模糊查询）
	Page  int    `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type UpdateParamValueRequest struct {
	ID         int    `json:"id" binding:"required"` // 配置ID
	ParamValue string `json:"param_value"`           // 参数值
}

type BatchUpdateParamConfRequest struct {
	Configs []UpdateParamConfRequest `json:"configs" binding:"required"` // 配置列表
}

type ParamConfByKeyResponse struct {
	Key    string                    `json:"key"`    // 配置键
	Params []ParamConfItemResponse   `json:"params"` // 参数列表
}

type ParamConfItemResponse struct {
	ID         int    `json:"id"`          // 配置ID
	Param      string `json:"param"`       // 参数名称
	ParamValue string `json:"param_value"` // 参数值
}

type ParamConfKeysResponse struct {
	Keys []string `json:"keys"` // 所有配置键列表
}
