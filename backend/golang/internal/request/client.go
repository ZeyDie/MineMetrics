package request

import "github.com/google/uuid"

type ClientRequest struct {
	UserID uuid.UUID `json:"userId" validate:"required"`

	OS struct {
		Bitness int    `json:"bitness" validate:"required,oneof=32 64"`
		Name    string `json:"name" validate:"required"`
	} `json:"os"`
	CPU struct {
		Threads int `json:"threads" validate:"required,min=1"`
		Cores   int `json:"cores" validate:"required,min=1"`
	} `json:"cpu"`
	GPUs struct {
		GPUs []GPU `json:"gpus" validate:"dive"`
	} `json:"gpus"`
	RAM struct {
		TotalRam     uint64 `json:"totalRam" validate:"required,min=1073741824"`
		AvailableRam uint64 `json:"availableRam"`
	} `json:"ram"`

	FPS           int `json:"fps" validate:"min=1,max=1000"`
	ViewDistance  int `json:"viewDistance" validate:"min=2,max=32"`
	EntityCount   int `json:"entityCount" validate:"min=0"`
	ParticleCount int `json:"particleCount" validate:"min=0"`

	Dimension struct {
		Namespace string `json:"namespace" validate:"required"`
		Path      string `json:"path" validate:"required"`
	} `json:"dimension"`
	ChunkPosList []ChunkPos `json:"chunkPosList" validate:"dive"`
}

type GPU struct {
	Name        string `json:"name" validate:"required"`
	DeviceID    string `json:"deviceId" validate:"required"`
	Vendor      string `json:"vendor" validate:"required"`
	VersionInfo string `json:"versionInfo" validate:"required"`
	VRAM        uint64 `json:"vram" validate:"min=0"`
}

type ChunkPos struct {
	X int `json:"x" validate:"required"`
	Z int `json:"z" validate:"required"`
}
