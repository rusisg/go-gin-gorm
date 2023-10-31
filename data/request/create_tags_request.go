package request

type CreateRequest struct {
	Name string `validate:"required, min=1, max=10" json:"name"`
}
