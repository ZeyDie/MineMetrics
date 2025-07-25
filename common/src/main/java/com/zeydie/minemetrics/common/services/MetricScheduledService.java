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
import java.util.function.Consumer;

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
                    @NotNull val request = HttpRequest.newBuilder()
                            .uri(URI.create(config.getEndpoint()))
                            .header("Content-Type", MediaType.JSON_UTF_8.toString())
                            .POST(HttpRequest.BodyPublishers.ofString(json))
                            .build();

                    if (config.isDebug()) {
                        log.info("Request {}\n{}", request, json);
                        log.info("Request Headers {}", request.headers());
                    }

                    @NotNull val response = MineMetrics.getHttpClient()
                            .sendAsync(
                                    request,
                                    HttpResponse.BodyHandlers.ofString()
                            ).thenAcceptAsync(
                                    httpResponse -> {
                                        if (config.isDebug())
                                            log.info("Response {}", httpResponse.body());
                                    }
                            );
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