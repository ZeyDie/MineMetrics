package com.zeydie.minemetrics.common.managers.forge;

import com.zeydie.minemetrics.MineMetrics;
import com.zeydie.minemetrics.common.managers.ConfigManager;
import me.shedaniel.autoconfig.AutoConfig;
import me.shedaniel.autoconfig.ConfigData;
import me.shedaniel.autoconfig.annotation.Config;
import me.shedaniel.autoconfig.serializer.Toml4jConfigSerializer;
import org.jetbrains.annotations.NotNull;

@Config(name = MineMetrics.MOD_ID)
public final class ConfigManagerImpl extends ConfigManager implements ConfigData {
    private static ConfigManager config;

    public static void init() {
        config = AutoConfig.register(ConfigManagerImpl.class, Toml4jConfigSerializer::new).getConfig();
    }

    public static @NotNull ConfigManager getConfig() {
        return config;
    }
}