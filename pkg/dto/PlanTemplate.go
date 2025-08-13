package dto

type SavePlanTemplateRequest struct {
	Name string `json:"name" binding:"required"`
	Plan string `json:"plan" binding:"required"`
}

type UpdatePlanTemplateRequest struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Plan string `json:"plan" binding:"required"`
}

type PlanTemplatePageRequest struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

type PlanTemplatePageResponse struct {
	List  []PlanTemplateResponse `json:"list"`
	Total int64                  `json:"total"`
}

type PlanTemplateResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Plan string `json:"plan"`
}
