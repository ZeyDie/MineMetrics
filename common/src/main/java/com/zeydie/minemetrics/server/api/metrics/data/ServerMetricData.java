package com.zeydie.minemetrics.server.api.metrics.data;

import com.zeydie.minemetrics.MineMetrics;
import com.zeydie.minemetrics.common.api.metrics.data.MetricData;
import lombok.Data;
import lombok.EqualsAndHashCode;
import org.jetbrains.annotations.NotNull;

@Data
@EqualsAndHashCode(callSuper = false)
public final class ServerMetricData extends MetricData {
    private final @NotNull MineMetrics.Environment environment = MineMetrics.Environment.SERVER;

    @Override
    public @NotNull String toJson() {
        return MineMetrics.GSON.toJson(this);
    }

    @Override
    public boolean isAvialable() {
        return false;//!this.chunkPosList.isEmpty();
    }
}
