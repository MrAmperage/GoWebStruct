package ApplicationCore

import "github.com/MrAmperage/GoWebStruct/WebCore"

type ApplicationCore struct {
	WebCore             WebCore.WebCore
	ApplicationSettings map[string]interface{}
}
