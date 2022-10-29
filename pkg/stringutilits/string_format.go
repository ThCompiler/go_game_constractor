package stringutilits

import (
	"fmt"
	"strings"
)

func StringFormat(format string, args ...interface{}) string {
	preparedArgs := make([]string, len(args))
	for i, v := range args {
		if i%2 == 0 {
			preparedArgs[i] = fmt.Sprintf("{%v}", v)
		} else {
			preparedArgs[i] = fmt.Sprint(v)
		}
	}
	r := strings.NewReplacer(preparedArgs...)
	return fmt.Sprint(r.Replace(format))
}
