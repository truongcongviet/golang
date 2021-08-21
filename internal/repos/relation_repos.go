package repos

type RelationRepo interface {
	CheckIfExist(id1 string, id2 string, status string) bool
	GetIdFromEmail(email string) (string, error)
	GetEmailByStatus(id string, status string) ([]string, error)
	GetRetrivableEmails(id string) ([]string, error)
	AddRelation(ids []string, status string) (bool, error)
}
