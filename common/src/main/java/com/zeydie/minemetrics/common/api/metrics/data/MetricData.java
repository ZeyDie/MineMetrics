package com.zeydie.minemetrics.common.api.metrics.data;

import com.zeydie.minemetrics.MineMetrics;
import lombok.val;
import net.minecraft.client.Minecraft;
import org.jetbrains.annotations.NotNull;

public abstract class MetricData {
    public abstract @NotNull String toJson();

    public abstract boolean isAvialable();
}