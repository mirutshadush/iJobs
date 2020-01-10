package service

import (
	"errors"
	"fmt"
	"github.com/miruts/iJobs/entity"
	"github.com/miruts/iJobs/usecases/session"
	"time"
)

type SessionServiceImpl struct {
	sessRepo session.SessionRepository
}

func NewSessionServiceImpl(ss session.SessionRepository) *SessionServiceImpl {
	return &SessionServiceImpl{sessRepo: ss}
}
func (ss *SessionServiceImpl) Sessions() ([]entity.Session, error) {
	return ss.sessRepo.Sessions()
}
func (ss *SessionServiceImpl) Session(id int) (entity.Session, error) {
	return ss.sessRepo.Session(id)
}
func (ss *SessionServiceImpl) DeleteSession(id int) (entity.Session, error) {
	return ss.sessRepo.DeleteSession(id)
}
func (ss *SessionServiceImpl) StoreSession(sess *entity.Session) (*entity.Session, error) {
	return ss.sessRepo.StoreSession(sess)
}
func (ss *SessionServiceImpl) SessionByValue(value string) (entity.Session, error) {
	return ss.sessRepo.SessionByValue(value)
}
func (ss *SessionServiceImpl) UpdateSession(sess *entity.Session) (*entity.Session, error) {
	return ss.sessRepo.UpdateSession(sess)
}
func (ss *SessionServiceImpl) Check(sess *entity.Session) (bool, entity.Session, error) {
	session, err := ss.SessionByValue(sess.Uuid)
	if err != nil {
		return false, session, errors.New("session not found")
	} else if time.Now().Sub(session.CreatedAt) < 6*time.Hour {
		session.CreatedAt = time.Now()
		session, err := ss.UpdateSession(&session)
		if err != nil {
			fmt.Println("Storing session error")
			return false, *session, err
		}
		return true, *session, nil
	} else {
		_, err := ss.DeleteSession(int(session.ID))
		if err != nil {
			return false, session, errors.New("invalid session")
		}
		return false, session, nil
	}
}