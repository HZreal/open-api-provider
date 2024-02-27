package routers

import "open-api-interface/api"

func addSentenceRouter() {
	apiGroup.GET("sentence", api.RootApi.GetSentence)
}
