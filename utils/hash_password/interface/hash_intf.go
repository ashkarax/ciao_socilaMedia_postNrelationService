package interface_hash_postNrelSvc

type IhashPassword interface {
	HashPassword(password string) string
	CompairPassword(hashedPassword string, plainPassword string) error
}
