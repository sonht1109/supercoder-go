package utils

func Blue(text string) string {
	return "\033[34m" + text + "\033[0m"
}

func Bold(text string) string {
	return "\033[1m" + text + "\033[0m"
}

func Underline(text string) string {
	return "\033[4m" + text + "\033[0m"
}

func Red(text string) string {
	return "\033[31m" + text + "\033[0m"
}
