package utils

type PicsumURLOptions struct {
	ID        int
	Blur      int // 1 to 10
	Grayscale bool
	Seed      string // if you want random image (but staying the same)
}

// PicsumURL generates an image URL for the Lorem Picsum online service
// More info can be found at its website: https://picsum.photos/
func PicsumURL(width int, height int, opt PicsumURLOptions) string {
	url := "https://picsum.photos/"
	if opt.Seed != "" {
		url += "/seed/" + opt.Seed + "/"
	}
	if opt.ID != 0 {
		url += "/id/" + ToString(opt.ID) + "/"
	}

	mainURL := url + ToString(width) + "/" + ToString(height)

	hasOptions := opt.Blur > 0 || opt.Grayscale
	if hasOptions {
		mainURL += "?"
	}

	if opt.Grayscale {
		mainURL += "&grayscale"
	}

	if opt.Blur > 0 {
		mainURL += "&blur=" + ToString(opt.Blur)
	}

	return mainURL
}
