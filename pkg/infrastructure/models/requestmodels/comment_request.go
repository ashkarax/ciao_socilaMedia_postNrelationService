package requestmodels_posnrel

type CommentRequest struct {
	PostId          uint64
	UserId          string
	CommentText     string
	ParentCommentId uint64
}
