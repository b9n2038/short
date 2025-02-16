// pkg/short/domain/ports/storage.go
package secondary

import "github.com/b9n2038/short/pkg/domain/model"

type ListStorage interface {
	Save(list *model.ShortList) error
	Load(name string) (*model.ShortList, error)
	Exists(name string) bool
}
