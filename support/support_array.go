package support

func GetStringFromArray(src []string, target int) string {
	if len(src)-1 < target {
		return "未知参数"
	}
	return src[target]
}
