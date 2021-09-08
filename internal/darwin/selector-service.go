package darwin

import "strings"

const serviceTemplate = `<ldb:serviceID>||SERVICE||</ldb:serviceID>`

func ServiceSelector(service string) string {
	return strings.Replace(serviceTemplate, "||SERVICE||", service, 1)
}
