package com.zeydie.minemetrics.common.api.oshi.data;

import lombok.Data;
import lombok.RequiredArgsConstructor;

public record OSData(
        int bitness,
        String name
) {
}