package com.zeydie.minemetrics;

import com.zeydie.minemetrics.MineMetrics;
import dev.architectury.platform.forge.EventBuses;
import net.minecraftforge.fml.common.Mod;
import net.minecraftforge.fml.javafmlmod.FMLJavaModLoadingContext;
import net.minecraftforge.fml.util.thread.EffectiveSide;

@Mod(MineMetrics.MOD_ID)
public final class MineMetricsForge {
    public MineMetricsForge() {
        EventBuses.registerModEventBus(MineMetrics.MOD_ID, FMLJavaModLoadingContext.get().getModEventBus());

        MineMetrics.init();
    }
}