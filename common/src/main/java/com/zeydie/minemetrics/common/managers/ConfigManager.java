package com.zeydie.minemetrics.common.managers;

import dev.architectury.injectables.annotations.ExpectPlatform;
import lombok.Data;
import org.jetbrains.annotations.NotNull;

@Data
public class ConfigManager {
    private boolean debug = true;
    private @NotNull String endpoint = "http://localhost:8080/metrics";

    @ExpectPlatform
    public static void init() {
        throw new AssertionError("Not implemented!");
    }

    @ExpectPlatform
    public static ConfigManager getConfig() {
        throw new AssertionError("Not implemented!");
    }
}