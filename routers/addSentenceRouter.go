package routers

import "open-api-provider/api"

func addSentenceRouter() {
	apiGroup.GET("sentence", api.RootApi.GetSentence)
}
