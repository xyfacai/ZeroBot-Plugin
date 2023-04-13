package query_material

import (
	"regexp"
	"testing"
)

func TestMatch(t *testing.T) {
	regex := regexp.MustCompile(`((b23|acg).tv|bili2233.cn)/[0-9a-zA-Z]+`)
	if matched := regex.FindStringSubmatch("https://b23.tv/mzDvDet?share_medium=android&share_source=qq&bbid=XU99A819C2045E20D9AA79F5542A7C4F6572B&ts=1681390837391"); matched != nil {
		return
	}
	return
}
