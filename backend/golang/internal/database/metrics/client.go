package metrics

import (
	"log/slog"
	"minemetrics_golang/internal/database"
	"minemetrics_golang/internal/database/entity"
	"minemetrics_golang/internal/request"
)

func InsertClientData(clientRequest request.ClientRequest) error {
	clientEntity := entity.ClientEntity{
		UserID: clientRequest.UserID,

		OSBitness: clientRequest.OS.Bitness,
		OSName:    clientRequest.OS.Name,

		CPUThreads: clientRequest.CPU.Threads,
		CPUCores:   clientRequest.CPU.Cores,

		RAMTotal:     clientRequest.RAM.TotalRam,
		RAMAvailable: clientRequest.RAM.AvailableRam,

		FPS:           clientRequest.FPS,
		ViewDistance:  clientRequest.ViewDistance,
		EntityCount:   clientRequest.EntityCount,
		ParticleCount: clientRequest.ParticleCount,

		DimensionNamespace: clientRequest.Dimension.Namespace,
		DimensionPath:      clientRequest.Dimension.Path,
	}

	for _, gpu := range clientRequest.GPUs.GPUs {
		clientEntity.GPUs = append(
			clientEntity.GPUs,
			entity.GPU{
				Name:        gpu.Name,
				DeviceID:    gpu.DeviceID,
				Vendor:      gpu.Vendor,
				VersionInfo: gpu.VersionInfo,
				VRAM:        gpu.VRAM,
			},
		)
	}

	for _, chunk := range clientRequest.ChunkPosList {
		if chunk.X == 0 && chunk.Z == 0 {
			continue
		}

		clientEntity.ChunkPositions = append(
			clientEntity.ChunkPositions,
			entity.ChunkPosition{
				X: chunk.X,
				Z: chunk.Z,
			},
		)
	}

	transaction := database.GetTransaction()

	if err := transaction.Create(&clientEntity).Error; err != nil {
		slog.Error("Failed to insert client data", "error", err)
		transaction.Rollback()
		return err
	}

	if err := transaction.Commit().Error; err != nil {
		slog.Error("Transaction failed", "error", err)
		return err
	}

	return nil
}
