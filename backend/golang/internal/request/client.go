package request

import "github.com/google/uuid"

type ClientRequest struct {
	UserID uuid.UUID `json:"userId" validate:"required"`

	OS struct {
		Bitness int8   `json:"bitness" validate:"required,oneof=32 64"`
		Name    string `json:"name" validate:"required"`
	} `json:"os"`
	CPU struct {
		Threads uint8 `json:"threads" validate:"required,min=1"`
		Cores   uint8 `json:"cores" validate:"required,min=1"`
	} `json:"cpu"`
	GPUs struct {
		GPUs []GPU `json:"gpus" validate:"dive"`
	} `json:"gpus"`
	RAM struct {
		TotalRam     uint64 `json:"totalRam" validate:"required,min=1073741824"`
		AvailableRam uint64 `json:"availableRam"`
	} `json:"ram"`

	FPS           uint16 `json:"fps" validate:"min=1,max=1000"`
	ViewDistance  uint8  `json:"viewDistance" validate:"min=2,max=32"`
	EntityCount   uint16 `json:"entityCount" validate:"min=0"`
	ParticleCount uint32 `json:"particleCount" validate:"min=0"`

	Dimension struct {
		Namespace string `json:"namespace" validate:"required"`
		Path      string `json:"path" validate:"required"`
	} `json:"dimension"`

	Position struct {
		X int16 `json:"x" validate:"required"`
		Y int16 `json:"y" validate:"required"`
		Z int16 `json:"z" validate:"required"`
	} `json:"position"`
}

type GPU struct {
	Name        string `json:"name" validate:"required"`
	DeviceID    string `json:"deviceId" validate:"required"`
	Vendor      string `json:"vendor" validate:"required"`
	VersionInfo string `json:"versionInfo" validate:"required"`
	VRAM        uint64 `json:"vram" validate:"min=0"`
}
