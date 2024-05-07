package interface_usecase_postnrel

type IRelationUseCase interface {
	GetCountsForUserProfile(userId *string) (*uint, *uint, *uint, *error)
	Follow(userId, userBId *string) *error
	UnFollow(userId, userBId *string) *error
	GetFollowersIds(*string) (*[]uint64, error)
	GetFollowingsIds(userId *string) (*[]uint64, error)
}
