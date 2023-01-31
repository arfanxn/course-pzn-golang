package apis

type CategoryUpdateRequest struct {
	Id   int32  `validate:"required" json:"id"`
	Name string `validate:"required,max=50,min=2" json:"name"`
}
