package com.zeydie.minemetrics.client.api.metrics;

import com.zeydie.minemetrics.client.level.ClientLevelInfo;
import net.minecraft.client.Minecraft;
import net.minecraft.resources.ResourceLocation;
import net.minecraft.world.level.ChunkPos;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.UUID;

public final class ClientMetricAPI {
    public static @NotNull Minecraft getMinecraft() {
        return Minecraft.getInstance();
    }

    public static @NotNull UUID getUserId() {
        return getMinecraft().getUser().getProfileId();
    }

    public static int getFps() {
        return getMinecraft().getFps();
    }

    public static int getViewDistance() {
        var options = getMinecraft().options;

        return options == null ? 0 : options.renderDistance().get();
    }

    public static @Nullable ResourceLocation getDimension() {
        return new ClientLevelInfo().getDimension();
    }

    public static int getEntityCount() {
        return new ClientLevelInfo().getMobCount();
    }

    public static int getParticleCount() {
        return new ClientLevelInfo().getParticleCount();
    }

    public static @NotNull List<ChunkPos> getChunkPosList() {
        return new ClientLevelInfo().getChunkPosList();
    }
}