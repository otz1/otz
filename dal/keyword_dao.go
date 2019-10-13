package dal

type KeywordDao interface {
	Insert()
	Delete()
	Create()
}

var KeywordDAO = &KeywordDaoImpl{}

type KeywordDaoImpl struct{}

func (k *KeywordDaoImpl) Insert() {

}

func (k *KeywordDaoImpl) Delete(id uint64) {

}

func (k *KeywordDaoImpl) Create(id uint64, word string) {

}
