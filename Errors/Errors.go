package Errors

func getString(err []string) string {
	if len(err) == 0 {
		return ""
	}
	return err[0]
}
