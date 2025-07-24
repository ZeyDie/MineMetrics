package entity

import (
	"gorm.io/gorm"
	"minemetrics_golang/internal/models/dto"
)

func ClientDTOToEntity(dto *dto.ClientDTO) (*ClientEntity, error) {
	clientEntity := ClientEntity{
		OS: OS{
			Bitness: dto.OS.Bitness,
			Name:    dto.OS.Name,
		},
		CPU: CPU{
			Threads: dto.CPU.Threads,
			Cores:   dto.CPU.Cores,
		},
		RAM: RAM{
			TotalRam:     dto.RAM.TotalRam,
			AvailableRam: dto.RAM.AvailableRam,
		},
		Dimension: Dimension{
			Namespace: dto.Dimension.Namespace,
			Path:      dto.Dimension.Path,
		},
		FPS:           dto.FPS,
		ViewDistance:  dto.ViewDistance,
		EntityCount:   dto.EntityCount,
		ParticleCount: dto.ParticleCount,
	}

	for _, gpu := range dto.GPUs.GPUs {
		clientEntity.GPUs = append(
			clientEntity.GPUs,
			GPU{
				Name:        gpu.Name,
				DeviceID:    gpu.DeviceID,
				Vendor:      gpu.Vendor,
				VersionInfo: gpu.VersionInfo,
				VRAM:        gpu.VRAM,
			},
		)
	}

	for _, chunk := range dto.ChunkPosList {
		clientEntity.ChunkPosList = append(
			clientEntity.ChunkPosList,
			ChunkPos{
				X: chunk.X,
				Z: chunk.Z,
			},
		)
	}

	return &clientEntity, nil
}

type ClientEntity struct {
	gorm.Model

	OS        `gorm:"embedded;embeddedPrefix:os_"`
	CPU       `gorm:"embedded;embeddedPrefix:cpu_"`
	RAM       `gorm:"embedded;embeddedPrefix:ram_"`
	Dimension `gorm:"embedded;embeddedPrefix:dimension_"`

	FPS           int `json:"fps"`
	ViewDistance  int `json:"viewDistance"`
	EntityCount   int `json:"entityCount"`
	ParticleCount int `json:"particleCount"`

	GPUs         []GPU      `gorm:"foreignKey:ClientID"`
	ChunkPosList []ChunkPos `gorm:"foreignKey:ClientID"`
}

type OS struct {
	Bitness int    `json:"bitness" gorm:"column:bitness"`
	Name    string `json:"name" gorm:"column:name"`
}

type CPU struct {
	Threads int `json:"threads" gorm:"column:threads"`
	Cores   int `json:"cores" gorm:"column:cores"`
}

type RAM struct {
	TotalRam     uint64 `json:"totalRam" gorm:"column:total_ram"`
	AvailableRam uint64 `json:"availableRam" gorm:"column:available_ram"`
}

type Dimension struct {
	Namespace string `json:"namespace" gorm:"column:namespace"`
	Path      string `json:"path" gorm:"column:path"`
}

type GPU struct {
	gorm.Model

	ClientID    uint   `json:"-"`
	Name        string `json:"name"`
	DeviceID    string `json:"deviceId"`
	Vendor      string `json:"vendor"`
	VersionInfo string `json:"versionInfo"`
	VRAM        uint64 `json:"vram"`
}

type ChunkPos struct {
	gorm.Model

	ClientID uint `json:"-"`
	X        int  `json:"x"`
	Z        int  `json:"z"`
}
