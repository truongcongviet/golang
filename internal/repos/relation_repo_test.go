package repos

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCheckIfExistFalse(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}
	ids := []string{"1", "2"}

	status := "FRIEND"

	sql_query := `select fr.relation_id 
	from friend_relationship fr 
	where ((fr.your_id = $1 and fr.friend_id = $2) or (fr.your_id =$2 and fr.friend_id =$1)) and status = $3`

	mock.ExpectQuery(regexp.QuoteMeta(sql_query)).
		WillReturnRows(sqlmock.NewRows([]string{"relation_id"}))

	resp := repo.CheckIfExist(ids[0], ids[1], status)

	assert.NotNil(t, resp)
	assert.Equal(t, false, resp)
}

func TestCheckIfExistTrue(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}
	ids := []string{"1", "2"}

	status := "FRIEND"

	sql_query := `select fr.relation_id 
	from friend_relationship fr 
	where ((fr.your_id = $1 and fr.friend_id = $2) or (fr.your_id =$2 and fr.friend_id =$1)) and status = $3`

	mock.ExpectQuery(regexp.QuoteMeta(sql_query)).
		WillReturnRows(sqlmock.NewRows([]string{"relation_id"}).AddRow("1"))

	resp := repo.CheckIfExist(ids[0], ids[1], status)

	assert.NotNil(t, resp)
	assert.Equal(t, true, resp)
}
func TestGetIdFromEmail(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}
	email := "quan12yt@gmail.com"
	sql_query := "select e.email_id from email e where e.email = '" + email + "'"

	mock.ExpectQuery(sql_query).
		WillReturnRows(sqlmock.NewRows([]string{"email_id"}).
			AddRow("2"))

	resp, err := repo.GetIdFromEmail("quan12yt@gmail.com")

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "2", resp)
}

func TestGetEmailByStatus(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}

	id := "1"
	status := "FRIEND"

	sql_query := `select distinct e.email
	from email e left join friend_relationship fr 
	on (e.email_id = fr.friend_id) or (e.email_id = fr.your_id)
	where (fr.your_id = $1 or fr.friend_id = $1) and fr.status = $2 and e.email_id != $1`

	mock.ExpectQuery(regexp.QuoteMeta(sql_query)).
		WillReturnRows(sqlmock.NewRows([]string{"email"}).
			AddRow("quan12yt@gmail.com"))

	resp, err := repo.GetEmailByStatus(id, status)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "quan12yt@gmail.com", resp[0])
}

func TestGetRetrivableEmails(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}
	id := "1"

	sql_query := `select distinct e.email 
	from friend_relationship fr left join email e
	on e.email_id = fr.your_id 
	where fr.friend_id = $1 and ((fr.status = 'FRIEND' or fr.status = 'SUBCRIBE') and fr.status != 'BLOCK')`

	mock.ExpectQuery(regexp.QuoteMeta(sql_query)).
		WillReturnRows(sqlmock.NewRows([]string{"email"}).
			AddRow("quan12yt@gmail.com"))

	resp, err := repo.GetRetrivableEmails(id)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "quan12yt@gmail.com", resp[0])
}

func TestAddRelationSucceed(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}
	id := []string{"1", "2"}
	status := "FRIEND"
	sql_query := `insert into friend_relationship (your_id, friend_id, status)
	values ($1, $2, $3), ($2, $1, $3)`
	result := sqlmock.NewResult(1, 1)

	mock.ExpectExec(regexp.QuoteMeta(sql_query)).WithArgs(id[0], id[1], status).WillReturnResult(result).WillReturnError(nil)

	resp, err := repo.AddRelation(id, status)

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, result, result)
}

func TestAddRelationFailed(t *testing.T) {
	db, mock := DbMock()
	repo := RelationRepoImp{Db: db}
	id := []string{"1", "2"}
	status := "FRIEND"
	sql_query := `insert into friend_relationship (your_id, friend_id, status)
	values ($1, $2, $3), ($2, $1, $3)`

	mock.ExpectExec(regexp.QuoteMeta(sql_query)).WithArgs(id[0], id[1], status).WillReturnResult(nil).WillReturnError(errors.New(""))

	_, err := repo.AddRelation(id, status)

	assert.NotNil(t, err)
}
