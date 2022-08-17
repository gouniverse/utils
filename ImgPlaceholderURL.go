package utils

// ImgPlaceholderURL returns a placeholder image
func ImgPlaceholderURL(width int, height int, text string) string {
	url := "https://via.placeholder.com/" + ToString(width) + "x" + ToString(height) + ".png?text=" + text
	return url
}
