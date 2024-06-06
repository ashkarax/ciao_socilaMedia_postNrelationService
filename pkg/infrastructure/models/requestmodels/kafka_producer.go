package requestmodels_posnrel

import "time"

type KafkaNotificationTopicModel struct {
	UserID      string
	ActorID     string
	ActionType  string
	TargetID    string
	TargetType  string
	CommentText string
	CreatedAt   time.Time
}
