package types

// ImageConfig img config
type ImageConfig struct {
	Context   ImageContext
	FontSize  float64
	ImagePath string
	Color     string
	Command   string
}

// ImageContext define image size
type ImageContext struct {
	Height int
	Width  int
}
