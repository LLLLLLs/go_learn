// Time        : 2019/10/17
// Description :

package gomock

//go:generate mockgen -destination=repo-mock/repo.go -package=repomock -source repository.go
//go:generate mockgen -destination=repo-mock/repo.go -package=repomock golearn/pkgtest/gomock Repository
type Repository interface {
	Create(key string, value interface{}) error
	Get(key string) (v interface{}, ok bool)
	Update(key string, value interface{}) error
	Remove(key string) error
}
