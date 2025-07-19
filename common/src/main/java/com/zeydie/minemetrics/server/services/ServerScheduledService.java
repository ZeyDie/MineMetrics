package com.zeydie.minemetrics.server.services;

import com.google.common.net.MediaType;
import com.google.common.util.concurrent.AbstractScheduledService;
import com.zeydie.minemetrics.MineMetrics;
import com.zeydie.minemetrics.common.services.MetricScheduledService;
import com.zeydie.minemetrics.server.api.metrics.data.ServerMetricData;
import com.zeydie.minemetrics.common.managers.ConfigManager;
import lombok.extern.slf4j.Slf4j;
import lombok.val;
import org.jetbrains.annotations.NotNull;
import org.jetbrains.annotations.Nullable;

import java.net.URI;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.concurrent.TimeUnit;

@Slf4j
public final class ServerScheduledService extends MetricScheduledService<ServerMetricData> {
    @Override
    public ServerMetricData getMetricData() {
        return new ServerMetricData();
    }
}