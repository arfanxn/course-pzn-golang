package entities

type CategoryEntity struct {
	Id   string
	Name string
}

// Mark struct implements the Mark interface
var _ EntityInterface = CategoryEntity{}       
var _ EntityInterface = (*CategoryEntity)(nil) 
