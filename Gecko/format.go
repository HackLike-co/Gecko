package gecko

import (
	"fmt"
	"strings"
)

// return a string with the payload in a C formated array
func C_FormatArray(b []byte, n string) string {
	// convert payload into comma separated string
	sPayload := formatBytes(b)

	return fmt.Sprintf("unsigned char %s[] = {\n%s\n};\n", n, sPayload)
}

func formatBytes(b []byte) string {
	s := fmt.Sprintf("% #x", b)
	sl := strings.Split(s, " ")

	return strings.Join(sl, ", ")
}
