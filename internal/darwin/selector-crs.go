package darwin

import "strings"

const crsTemplate = `<ldb:crs>||CRS||</ldb:crs>`

func CRSSelector(crs string) string {
	return strings.Replace(crsTemplate, "||CRS||", crs, 1)
}
