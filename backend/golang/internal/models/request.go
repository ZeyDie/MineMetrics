package models

type Request struct {
	Environment string `json:"environment"`
	OS          struct {
		Bitness int    `json:"bitness"`
		Name    string `json:"name"`
	} `json:"os"`
	CPU struct {
		Threads int `json:"threads"`
		Cores   int `json:"cores"`
	} `json:"cpu"`
	GPUs struct {
		GPUs []GPU `json:"gpus"`
	} `json:"gpus"`
	RAM struct {
		TotalRam     uint64 `json:"totalRam"`
		AvailableRam uint64 `json:"availableRam"`
	} `json:"ram"`
	FPS           int `json:"fps"`
	ViewDistance  int `json:"viewDistance"`
	EntityCount   int `json:"entityCount"`
	ParticleCount int `json:"particleCount"`
	Dimension     struct {
		Namespace string `json:"namespace"`
		Path      string `json:"path"`
	} `json:"dimension"`
	ChunkPosList []ChunkPos `json:"chunkPosList"`
}

type GPU struct {
	Name        string `json:"name"`
	DeviceID    string `json:"deviceId"`
	Vendor      string `json:"vendor"`
	VersionInfo string `json:"versionInfo"`
	VRAM        uint64 `json:"vram"`
}

type ChunkPos struct {
	X int `json:"x"`
	Z int `json:"z"`
}
