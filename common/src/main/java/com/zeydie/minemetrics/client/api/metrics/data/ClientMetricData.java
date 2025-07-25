package com.zeydie.minemetrics.client.api.metrics.data;

import com.zeydie.minemetrics.MineMetrics;
import com.zeydie.minemetrics.common.api.metrics.data.MetricData;
import com.zeydie.minemetrics.client.api.metrics.ClientMetricAPI;
import com.zeydie.minemetrics.common.api.oshi.OshiAPI;
import com.zeydie.minemetrics.common.api.oshi.data.CpuData;
import com.zeydie.minemetrics.common.api.oshi.data.GpusData;
import com.zeydie.minemetrics.common.api.oshi.data.OSData;
import com.zeydie.minemetrics.common.api.oshi.data.RamData;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.val;
import net.minecraft.client.Minecraft;
import net.minecraft.resources.ResourceLocation;
import net.minecraft.world.level.ChunkPos;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

@Data
@EqualsAndHashCode(callSuper = false)
public final class ClientMetricData extends MetricData {
    private final @NotNull UUID userId = ClientMetricAPI.getUserId();

    private final @Nullable OSData os = OshiAPI.getOs();
    private final @Nullable CpuData cpu = OshiAPI.getCpu();
    private final @Nullable GpusData gpus = OshiAPI.getGpus();
    private final @Nullable RamData ram = OshiAPI.getRam();

    private final int fps = ClientMetricAPI.getFps();
    private final int viewDistance = ClientMetricAPI.getViewDistance();
    private final int entityCount = ClientMetricAPI.getEntityCount();
    private final int particleCount = ClientMetricAPI.getParticleCount();

    private final @Nullable ResourceLocation dimension = ClientMetricAPI.getDimension();
    private final @NotNull List<ChunkPos> chunkPosList = ClientMetricAPI.getChunkPosList();

    @Override
    public @NotNull String toJson() {
        return MineMetrics.GSON.toJson(this);
    }

    @Override
    public boolean isAvialable() {
        val minecraft = ClientMetricAPI.getMinecraft();

        return !minecraft.isSingleplayer() && !minecraft.isPaused() && minecraft.isWindowActive() && !this.chunkPosList.isEmpty();
    }
}