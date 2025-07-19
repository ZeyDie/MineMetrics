package com.zeydie.minemetrics.common.managers.fabric;

import net.fabricmc.api.EnvType;
import net.fabricmc.loader.api.FabricLoader;

public final class EnvironmentManagerImpl {
    public static boolean isClient() {
        return FabricLoader.getInstance().getEnvironmentType() == EnvType.CLIENT;
    }

    public static boolean isServer() {
        return FabricLoader.getInstance().getEnvironmentType() == EnvType.SERVER;
    }
}