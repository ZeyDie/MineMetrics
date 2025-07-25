package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientEntity struct {
	gorm.Model

	UserID             uuid.UUID
	OSBitness          int
	OSName             string
	CPUThreads         int
	CPUCores           int
	RAMTotal           uint64
	RAMAvailable       uint64
	FPS                int
	ViewDistance       int
	EntityCount        int
	ParticleCount      int
	DimensionNamespace string
	DimensionPath      string

	GPUs           []GPU           `gorm:"foreignKey:ClientEntityID"`
	ChunkPositions []ChunkPosition `gorm:"foreignKey:ClientEntityID"`
}

type GPU struct {
	gorm.Model

	ClientEntityID uint

	Name        string
	DeviceID    string
	Vendor      string
	VersionInfo string
	VRAM        uint64
}

type ChunkPosition struct {
	gorm.Model

	ClientEntityID uint

	X int
	Z int
}
