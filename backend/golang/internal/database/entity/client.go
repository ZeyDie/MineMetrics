package entity

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientEntity struct {
	gorm.Model

	UserID             uuid.UUID
	OSBitness          int8
	OSName             string
	CPUThreads         uint8
	CPUCores           uint8
	RAMTotal           uint64
	RAMAvailable       uint64
	FPS                uint16
	ViewDistance       uint8
	EntityCount        uint16
	ParticleCount      uint32
	DimensionNamespace string
	DimensionPath      string

	GPUsJSON       string          `gorm:"type:json" json:"-"`
	ChunksJSON     string          `gorm:"type:json" json:"-"`
	GPUs           []GPU           `gorm:"-" json:"gpus"`
	ChunkPositions []ChunkPosition `gorm:"-" json:"chunk_positions"`
}

type GPU struct {
	Name        string
	DeviceID    string
	Vendor      string
	VersionInfo string
	VRAM        uint64
}

type ChunkPosition struct {
	X int16
	Z int16
}

func (clientEntity *ClientEntity) BeforeSave(transaction *gorm.DB) error {
	gpusBytes, err := json.Marshal(clientEntity.GPUs)
	if err != nil {
		return err
	}

	clientEntity.GPUsJSON = string(gpusBytes)

	chunksBytes, err := json.Marshal(clientEntity.ChunkPositions)
	if err != nil {
		return err
	}

	clientEntity.ChunksJSON = string(chunksBytes)

	return nil
}

func (clientEntity *ClientEntity) AfterFind(transaction *gorm.DB) error {
	if clientEntity.GPUsJSON != "" {
		err := json.Unmarshal([]byte(clientEntity.GPUsJSON), &clientEntity.GPUs)
		if err != nil {
			return err
		}
	}

	if clientEntity.ChunksJSON != "" {
		err := json.Unmarshal([]byte(clientEntity.ChunksJSON), &clientEntity.ChunkPositions)
		if err != nil {
			return err
		}
	}

	return nil
}
