package util

// IsArgsNum コマンド引数の数をチェック
func IsArgsNum(command []string, argsNum int) bool {
	return len(command) != argsNum
}
