package requestmodels_posnrel

type AddPostData struct {
	Caption   *string
	UserId    *string
	MediaURLs []string
}

type EditPost struct {
	Caption string
	UserId  string
	PostId  string
}
