package com.zeydie.minemetrics.common.api.oshi.data;

import lombok.Data;
import lombok.RequiredArgsConstructor;
import oshi.hardware.GraphicsCard;

import java.util.List;

public record GpusData(
        List<GraphicsCard> gpus
) {
}