package com.zeydie.minemetrics;

import com.zeydie.minemetrics.MineMetrics;
import net.fabricmc.api.ModInitializer;

public final class MineMetricsFabric implements ModInitializer {
    @Override
    public void onInitialize() {
        MineMetrics.init();
    }
}