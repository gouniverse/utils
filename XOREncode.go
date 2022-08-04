package utils

func XOREncode(buffer []byte, key []byte) []byte {
	keylen := len(key)
	ciphertext := make([]byte, len(buffer))
	for i, b := range buffer {
		ciphertext[i] = (b ^ key[i%keylen])
	}
	return ciphertext
}
