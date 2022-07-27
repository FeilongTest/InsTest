package utils

//CopyMap 复制Map
func CopyMap(dst, src map[string]string) map[string]string {
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
