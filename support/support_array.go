package support

func GetStringFromArray(src []string, target int) string {
	if len(src)-1 < target {
		return ""
	}
	return src[target]
}
