package ApplicationCore

import "github.com/MrAmperage/GoWebStruct/WebCore"

type ApplicationCore struct {
	ApplicationPort     int64
	WebCore             WebCore.WebCore
	ApplicationSettings map[string]interface{}
}
