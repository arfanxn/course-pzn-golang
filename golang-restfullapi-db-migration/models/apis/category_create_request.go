package apis

type CategoryCreateRequest struct {
	Name string `validate:"required,max=50,min=2" json:"name"`
}
