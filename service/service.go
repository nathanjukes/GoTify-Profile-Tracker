package service

type Service interface {
	FollowUser(id string) error
}
