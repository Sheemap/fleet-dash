package data

import (
	"context"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	SaveSession(session *Session) error
	GetSession(sessionId string) (*Session, error)
	GetCharacterActiveSession(characterID string) (*Session, error)
	GetSessionByFleet(fleetID string) (*Session, error)
	SaveEveLogEvent(event *Event) error
}

func NewRepository() Repository {
	db := initGorm()
	return &repository{
		db: db,
	}
}

func initGorm() *gorm.DB{
	db, err := gorm.Open(postgres.Open("postgresql://jamis:JFc8yI2euSGTRtNY2Vracw@free-tier4.aws-us-west-2.cockroachlabs.cloud:26257/fleet-dash-2252.defaultdb"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Session{}, &Event{})
	if err != nil {
		panic(err)
	}

	return db
}

func (r *repository) SaveSession(session *Session) error {
	err := crdbgorm.ExecuteTx(context.Background(), r.db, nil,
		func(tx *gorm.DB) error {
			return r.db.Create(session).Error
		},
	)
	return err
}

func (r *repository) GetSession(sessionId string) (*Session, error) {
	var session Session
	err := crdbgorm.ExecuteTx(context.Background(), r.db, nil,
		func(tx *gorm.DB) error {
			return r.db.First(session, sessionId).Error
		},
	)
	return &session, err
}

func (r *repository) GetCharacterActiveSession(characterID string) (*Session, error) {
	var session Session
	err := crdbgorm.ExecuteTx(context.Background(), r.db, nil,
		func(tx *gorm.DB) error {
			return r.db.Where("ended_at IS NULL").Where("character_id = ?", characterID).First(&session).Error
		},
	)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &session, err
}

func (r *repository) GetSessionByFleet(fleetID string) (*Session, error){
	var session Session
	err := crdbgorm.ExecuteTx(context.Background(), r.db, nil,
		func(tx *gorm.DB) error {
			return r.db.Where("ended_at IS NULL").Where("fleet_id = ?", fleetID).First(&session).Error
		},
	)
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &session, err
}

func (r *repository) SaveEveLogEvent(event *Event) error {
	err := crdbgorm.ExecuteTx(context.Background(), r.db, nil,
		func(tx *gorm.DB) error {
			return r.db.Create(event).Error
		},
	)
	return err
}