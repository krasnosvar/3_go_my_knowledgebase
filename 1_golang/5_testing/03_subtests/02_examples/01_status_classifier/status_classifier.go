package status

func Classify(code int) string {
	if code >= 500 {
		return "server"
	}
	if code >= 400 {
		return "client"
	}
	return "ok"
}
