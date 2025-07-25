package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientEntity struct {
	gorm.Model

	UserID uuid.UUID

	OSBitness    int8
	OSName       string
	CPUThreads   uint8
	CPUCores     uint8
	RAMTotal     uint64
	RAMAvailable uint64

	FPS          uint16
	ViewDistance uint8

	PositionEntity []PositionEntity `gorm:"foreignKey:ClientEntityID"`

	GPUs       string `gorm:"type:json" json:"-"`
	GPUsStruct []GPU  `gorm:"-" json:"gpus_struct"`
}

type GPU struct {
	Name        string
	DeviceID    string
	Vendor      string
	VersionInfo string
	VRAM        uint64
}

type PositionEntity struct {
	gorm.Model

	ClientEntityID uint `json:"-"`

	DimensionNamespace string
	DimensionPath      string

	X int16
	Y int16
	Z int16

	EntityCount   uint16
	ParticleCount uint32
}

func (clientEntity *ClientEntity) BeforeSave(transaction *gorm.DB) error {
	gpusBytes, err := json.Marshal(clientEntity.GPUsStruct)
	if err != nil {
		return err
	}

	clientEntity.GPUs = string(gpusBytes)

	return nil
}

func (clientEntity *ClientEntity) AfterFind(transaction *gorm.DB) error {
	if clientEntity.GPUs != "" {
		err := json.Unmarshal([]byte(clientEntity.GPUs), &clientEntity.GPUsStruct)
		if err != nil {
			return err
		}
	}

	return nil
}
