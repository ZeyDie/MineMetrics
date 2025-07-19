package com.zeydie.minemetrics.server.level;

import com.google.common.collect.Lists;
import com.zeydie.minemetrics.common.level.Levelnfo;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.val;
import net.minecraft.resources.ResourceLocation;
import net.minecraft.server.level.ServerLevel;
import net.minecraft.world.level.ChunkPos;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.util.List;
import java.util.stream.StreamSupport;

@Data
@EqualsAndHashCode(callSuper = false)
public final class ServerLevelInfo extends Levelnfo<ServerLevel> {
    public ServerLevelInfo(@NotNull final ServerLevel level) {
        super(level);
    }

    @Override
    public @Nullable ResourceLocation getDimension() {
        return this.isValid() ? this.getLevel().dimension().location() : null;
    }

    @Override
    public int getMobCount() {
        return this.isValid() ? (int) StreamSupport.stream(this.getLevel().getAllEntities().spliterator(), false).count() : 0;
    }

    @Override
    public int getParticleCount() {
        return 0;
    }

    @Override
    public List<ChunkPos> getChunkPosList() {
        val list = Lists.<ChunkPos>newArrayList();

        if (this.isValid())
            this.getLevel()
                    .getPlayers(serverPlayer -> true)
                    .forEach(
                            serverPlayer -> {
                                val playerChunkPos = serverPlayer.chunkPosition();

                                //TODO Nearby chunks

                                list.add(playerChunkPos);
                            }
                    );

        return list;
    }
}