package com.zeydie.minemetrics;

import com.google.common.util.concurrent.Service;
import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.zeydie.minemetrics.common.managers.ConfigManager;
import com.zeydie.minemetrics.common.managers.EnvironmentManager;
import com.zeydie.minemetrics.client.services.ClientScheduledService;
import com.zeydie.minemetrics.server.services.ServerScheduledService;
import lombok.Getter;
import org.jetbrains.annotations.NotNull;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.http.HttpClient;
import java.time.Duration;

public final class MineMetrics {
    public static final @NotNull String MOD_ID = "minemetrics";
    public static Gson GSON;

    @Getter
    private static final @NotNull HttpClient httpClient = HttpClient.newHttpClient();

    private static @NotNull Service scheduler;

    public static void init() {
        ConfigManager.init();

        @NotNull var gsonBuilder = new GsonBuilder();

        if (ConfigManager.getConfig().isDebug())
            gsonBuilder = gsonBuilder.setPrettyPrinting();

        GSON = gsonBuilder.create();

        scheduler = EnvironmentManager.isClient() ? new ClientScheduledService() : new ServerScheduledService();
        scheduler.startAsync();
    }

    public static enum Environment {
        CLIENT,
        SERVER
    }
}