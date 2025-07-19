package com.zeydie.minemetrics.common.api.oshi.data;

import lombok.Data;
import lombok.RequiredArgsConstructor;

public record CpuData(
        int threads,
        int cores
) {
}