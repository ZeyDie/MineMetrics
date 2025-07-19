package com.zeydie.minemetrics.common.api.oshi.data;

import lombok.Data;
import lombok.RequiredArgsConstructor;

public record RamData(
        long totalRam,
        long availableRam
) {
}