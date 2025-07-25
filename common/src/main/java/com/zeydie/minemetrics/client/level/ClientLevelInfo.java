package com.zeydie.minemetrics.client.level;

import com.google.common.collect.Lists;
import com.zeydie.minemetrics.client.api.metrics.ClientMetricAPI;
import com.zeydie.minemetrics.common.level.Levelnfo;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.val;
import net.minecraft.client.Minecraft;
import net.minecraft.client.multiplayer.ClientLevel;
import net.minecraft.resources.ResourceLocation;
import net.minecraft.world.level.ChunkPos;
import org.jetbrains.annotations.Nullable;

import java.util.List;

@Data
@EqualsAndHashCode(callSuper = false)
public final class ClientLevelInfo extends Levelnfo<ClientLevel> {
    public ClientLevelInfo() {
        super(ClientMetricAPI.getMinecraft().level);
    }

    @Override
    public @Nullable ResourceLocation getDimension() {
        return this.isValid() ? this.getLevel().dimension().location() : null;
    }

    @Override
    public int getMobCount() {
        return this.isValid() ? this.getLevel().getEntityCount() : 0;
    }

    @Override
    public int getParticleCount() {
        return this.isValid() ? Integer.parseInt(ClientMetricAPI.getMinecraft().particleEngine.countParticles()) : 0;
    }

    @Override
    public List<ChunkPos> getChunkPosList() {
        val list = Lists.<ChunkPos>newArrayList();

        if (this.isValid()) {
            val centeredChunk = new ChunkPos(Minecraft.getInstance().player.blockPosition());
            val renderDistance = ClientMetricAPI.getViewDistance();

            for (int x = -renderDistance; x < renderDistance; x++)
                for (int z = -renderDistance; z < renderDistance; z++)
                    list.add(new ChunkPos(centeredChunk.x + x, centeredChunk.z + z));
        }

        return list;
    }
}