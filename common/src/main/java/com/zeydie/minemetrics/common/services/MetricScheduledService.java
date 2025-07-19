package com.zeydie.minemetrics.common.services;

import com.google.common.net.MediaType;
import com.google.common.util.concurrent.AbstractScheduledService;
import com.zeydie.minemetrics.MineMetrics;
import com.zeydie.minemetrics.common.api.metrics.data.MetricData;
import com.zeydie.minemetrics.common.managers.ConfigManager;
import com.zeydie.minemetrics.server.api.metrics.data.ServerMetricData;
import lombok.extern.slf4j.Slf4j;
import lombok.val;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.net.URI;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.concurrent.TimeUnit;

@Slf4j
public abstract class MetricScheduledService<T extends MetricData> extends AbstractScheduledService {
    @Override
    protected void runOneIteration() throws Exception {
        @Nullable val config = ConfigManager.getConfig();

        if (config != null)
            try {
                @NotNull val metricData = this.getMetricData();
                @NotNull val json = metricData.toJson();

                if (metricData.isAvialable()) {
                    if (config.isDebug())
                        log.info("Data {}", json);
                    
                    @NotNull val response = MineMetrics.getHttpClient()
                            .sendAsync(
                                    HttpRequest.newBuilder()
                                            .uri(URI.create(config.getEndpoint()))
                                            .header("Content-Type", MediaType.JSON_UTF_8.type())
                                            .POST(HttpRequest.BodyPublishers.ofString(json))
                                            .build(),
                                    HttpResponse.BodyHandlers.ofString()
                            ).get();

                    if (config.isDebug())
                        log.info("Response {}", response.body());
                }
            } catch (@NotNull final Exception exception) {
                if (config.isDebug())
                    log.error("Error in scheduler", exception);
            }
    }

    @Override
    protected @NotNull Scheduler scheduler() {
        return Scheduler.newFixedRateSchedule(0, 30, TimeUnit.SECONDS);
    }

    public abstract T getMetricData();
}