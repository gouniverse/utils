package hb

// ImgPlaceholderURL returns a placeholder image
func ImgPlaceholderURL(width, height, text string) string {
	url := "https://via.placeholder.com/" + width + "x" + height + ".png?text=" + text
	return url
}
