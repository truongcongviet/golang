package service

import (
	"errors"
	"friend-management-v1/internal/repos"
	"friend-management-v1/internal/utils"
	"friend-management-v1/model"
)

type RelationServiceImp struct {
	repo repos.RelationRepo
}

func NewRelationService(rp repos.RelationRepo) RelationService {
	return &RelationServiceImp{
		repo: rp,
	}
}

func (s *RelationServiceImp) GetFriendsEmail(rq model.GetFriendsRequest) ([]string, error) {
	ids, err := s.repo.GetIdFromEmail(rq.Email)
	if err != nil {
		return nil, err
	}
	return s.repo.GetEmailByStatus(ids, "FRIEND")
}

func (s *RelationServiceImp) Addfriend(rq model.AddAndGetCommonRequest) (bool, error) {
	id1, err1 := s.repo.GetIdFromEmail(rq.Friends[0])
	id2, err2 := s.repo.GetIdFromEmail(rq.Friends[1])

	if err1 != nil {
		return false, err1
	}
	if err2 != nil {
		return false, err2
	}

	if re := s.repo.CheckIfExist(id1, id2, "FRIEND"); re {
		return false, errors.New("2 emails are already being friend")
	}
	ids := make([]string, 2)
	ids[0] = id1
	ids[1] = id2
	_, err := s.repo.AddRelation(ids, "FRIEND")
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *RelationServiceImp) GetCommonFriends(rq model.AddAndGetCommonRequest) ([]string, error) {
	id1, err1 := s.repo.GetIdFromEmail(rq.Friends[0])
	id2, err2 := s.repo.GetIdFromEmail(rq.Friends[1])

	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}

	slice1, err1 := s.repo.GetEmailByStatus(id1, "FRIEND")
	slice2, err2 := s.repo.GetEmailByStatus(id2, "FRIEND")
	if err1 != nil {
		return nil, err1
	}
	if err2 != nil {
		return nil, err2
	}
	result := utils.RetainSlices(slice1, slice2)
	return result, nil
}

func (s *RelationServiceImp) SubcribeToEmail(rq model.SubcribeAndBlockRequest) (bool, error) {

	id1, err1 := s.repo.GetIdFromEmail(rq.Requestor)
	id2, err2 := s.repo.GetIdFromEmail(rq.Target)

	if err1 != nil {
		return false, err1
	}
	if err2 != nil {
		return false, err2
	}

	if re := s.repo.CheckIfExist(id1, id2, "BLOCK"); re {
		return false, errors.New("target email has been blocked")
	}
	if re := s.repo.CheckIfExist(id1, id2, "FRIEND"); re {
		return false, errors.New("already being friend to the target email, no need to subcribe")
	}
	if re := s.repo.CheckIfExist(id1, id2, "SUBCRIBE"); re {
		return false, errors.New("already subcribe to the target email")
	}

	ids := make([]string, 2)
	ids[0] = id1
	ids[1] = id2
	result, err := s.repo.AddRelation(ids, "SUBCRIBE")
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *RelationServiceImp) BlockEmail(rq model.SubcribeAndBlockRequest) (bool, error) {

	id1, err1 := s.repo.GetIdFromEmail(rq.Requestor)
	id2, err2 := s.repo.GetIdFromEmail(rq.Target)

	if err1 != nil {
		return false, err1
	}
	if err2 != nil {
		return false, err2
	}

	if re := s.repo.CheckIfExist(id1, id2, "BLOCK"); re {
		return false, errors.New("target email has already being blocked")
	}
	ids := make([]string, 2)
	ids[0] = id1
	ids[1] = id2
	result, err := s.repo.AddRelation(ids, "BLOCK")
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *RelationServiceImp) RetrieveContactEmail(rq model.RetrieveRequest) ([]string, error) {
	id, err := s.repo.GetIdFromEmail(rq.Sender)
	emails := utils.GetEmailsFromText(rq.Text)

	if err != nil {
		return nil, err
	}

	emails2, err := s.repo.GetRetrivableEmails(id)
	if err != nil {
		return nil, err
	}
	result := append(emails, emails2...)

	return utils.Unique(result), nil
}
