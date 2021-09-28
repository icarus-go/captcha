package adapter

type Adapter string

const (
	SMS   Adapter = "ext"
	Email Adapter = "email"
	Image Adapter = "image"
)

//Value
func (a Adapter) Value() string {
	return string(a)
}
