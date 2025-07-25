package metrics

import (
	"fmt"
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

		FPS:          clientRequest.FPS,
		ViewDistance: clientRequest.ViewDistance,
	}
	positionEntity := entity.PositionEntity{
		DimensionNamespace: clientRequest.Dimension.Namespace,
		DimensionPath:      clientRequest.Dimension.Path,

		X: clientRequest.Position.X,
		Y: clientRequest.Position.Y,
		Z: clientRequest.Position.Z,

		EntityCount:   clientRequest.EntityCount,
		ParticleCount: clientRequest.ParticleCount,
	}

	if positionEntity.X == 0 && positionEntity.Y == 0 && positionEntity.Z == 0 {
		return fmt.Errorf("Invalid position %f %f %f", positionEntity.X, positionEntity.Y, positionEntity.Z)
	}

	for _, gpu := range clientRequest.GPUs.GPUs {
		clientEntity.GPUsStruct = append(
			clientEntity.GPUsStruct,
			entity.GPU{
				Name:        gpu.Name,
				DeviceID:    gpu.DeviceID,
				Vendor:      gpu.Vendor,
				VersionInfo: gpu.VersionInfo,
				VRAM:        gpu.VRAM,
			},
		)
	}

	transaction := database.GetTransaction()

	userEntity := transaction.Find(&clientEntity, "user_id = ?", clientRequest.UserID)

	if userEntity != nil {
		userEntity.Updates(&clientEntity)
	} else {
		if err := transaction.Create(&clientEntity).Error; err != nil {
			slog.Error("Failed to insert client data", "error", err)
			transaction.Rollback()
			return err
		}
	}

	if err := transaction.Create(&positionEntity).Error; err != nil {
		slog.Error("Failed to insert positions", "error", err)
		transaction.Rollback()
		return err
	}

	if err := transaction.Commit().Error; err != nil {
		slog.Error("Transaction failed", "error", err)
		return err
	}

	return nil
}
