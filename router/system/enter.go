package system

type RouterGroup struct {
	BaseRouter
	MediaConfigRouter
	PlanTemplateRouter
	TaskManageRouter
	AlarmManageRouter
	NetworkConfigRouter
	DeviceListRouter
	ParamConfRouter
	ModelsInfoRouter
	AlarmCategoryRouter
	ThresholdConfRouter
}
