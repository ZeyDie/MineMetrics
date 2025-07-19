package com.zeydie.minemetrics.common.level;

import lombok.Data;
import lombok.RequiredArgsConstructor;
import net.minecraft.resources.ResourceLocation;
import net.minecraft.world.level.ChunkPos;
import net.minecraft.world.level.Level;
import org.jetbrains.annotations.Nullable;

import java.util.List;

@Data
@RequiredArgsConstructor
public abstract class Levelnfo<T extends Level> {
    private final @Nullable T level;

    public boolean isValid() {
        return this.level != null;
    }

    public abstract @Nullable ResourceLocation getDimension();

    public abstract int getMobCount();

    public abstract int getParticleCount();

    public abstract List<ChunkPos> getChunkPosList();
}