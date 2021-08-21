package repos

import (
	"database/sql"
	"errors"
)

type RelationRepoImp struct {
	Db *sql.DB
}

func NewRelationRepo(db *sql.DB) RelationRepo {
	return &RelationRepoImp{
		Db: db,
	}
}

func (repo *RelationRepoImp) CheckIfExist(id1 string, id2 string, status string) bool {
	sql_query := `select fr.relation_id 
	from friend_relationship fr 
	where ((fr.your_id = $1 and fr.friend_id = $2) or (fr.your_id =$2 and fr.friend_id =$1)) and status = $3`
	rows, _ := repo.Db.Query(sql_query, id1, id2, status)
	return (rows.Next())
}

func (repo *RelationRepoImp) GetIdFromEmail(email string) (string, error) {
	sql_query := "select e.email_id from email e where e.email = '" + email + "'"

	rows, err := repo.Db.Query(sql_query)
	if err != nil {
		return "", err
	}
	var ids string
	for rows.Next() {
		var i string
		err = rows.Scan(&i)
		if err != nil {
			return "", err
		}
		ids = i
	}
	if ids == "" {
		return "", errors.New("email: " + email + " is not exist in database")
	}
	return ids, nil
}

func (repo *RelationRepoImp) GetEmailByStatus(id string, status string) ([]string, error) {
	sql_query := `select distinct e.email
	from email e left join friend_relationship fr 
	on (e.email_id = fr.friend_id) or (e.email_id = fr.your_id)
	where (fr.your_id = $1 or fr.friend_id = $1) and fr.status = $2 and e.email_id != $1`

	rows, err := repo.Db.Query(sql_query, id, status)
	if err != nil {
		return nil, err
	}
	var friends []string
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return nil, err
		}
		friends = append(friends, email)
	}
	// defer pro.Db.Close()
	return friends, nil
}

func (repo *RelationRepoImp) GetRetrivableEmails(id string) ([]string, error) {
	sql_query := `select distinct e.email 
	from friend_relationship fr left join email e
	on e.email_id = fr.your_id 
	where fr.friend_id = $1 and ((fr.status = 'FRIEND' or fr.status = 'SUBCRIBE') and fr.status != 'BLOCK')`

	rows, err := repo.Db.Query(sql_query, id)
	if err != nil {
		return nil, err
	}
	var friends []string
	for rows.Next() {
		var email string
		err = rows.Scan(&email)
		if err != nil {
			return nil, err
		}
		friends = append(friends, email)
	}
	// defer pro.Db.Close()
	return friends, nil
}

func (repo *RelationRepoImp) AddRelation(ids []string, status string) (bool, error) {
	sql_query := `insert into friend_relationship (your_id, friend_id, status)
	values ($1, $2, $3), ($2, $1, $3)`

	result, err := repo.Db.Exec(sql_query, ids[0], ids[1], status)
	rs := (result == nil)
	if err != nil {
		return rs, err
	}
	return rs, nil
}
