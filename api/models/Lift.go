package models

import (
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Lift struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Lift) Prepare() {
	p.ID = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Lift) Validate() error {

	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}

	return nil
}

func (p *Lift) SaveLift(db *gorm.DB) (*Lift, error) {
	var err error
	err = db.Debug().Model(&Lift{}).Create(&p).Error
	if err != nil {
		return &Lift{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Error
		if err != nil {
			return &Lift{}, err
		}
	}
	return p, nil
}

func (p *Lift) FindAllLifts(db *gorm.DB) (*[]Lift, error) {
	var err error
	lifts := []Lift{}
	err = db.Debug().Model(&Lift{}).Limit(100).Find(&lifts).Error
	if err != nil {
		return &[]Lift{}, err
	}
	return &lifts, nil
}

func (p *Lift) FindLiftByID(db *gorm.DB, pid uint64) (*Lift, error) {
	var err error
	err = db.Debug().Model(&Lift{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Lift{}, err
	}
	return p, nil
}

func (p *Lift) UpdateALift(db *gorm.DB) (*Lift, error) {

	var err error

	err = db.Debug().Model(&Lift{}).Where("id = ?", p.ID).Updates(Lift{Name: p.Name, Description: p.Description, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Lift{}, err
	}
	return p, nil
}

func (p *Lift) DeleteALift(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Lift{}).Where("id = ? and author_id = ?", pid, uid).Take(&Lift{}).Delete(&Lift{})

	if db.Error != nil {
		fmt.Printf("Error deleting a lift: %+v", db.Error)
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
