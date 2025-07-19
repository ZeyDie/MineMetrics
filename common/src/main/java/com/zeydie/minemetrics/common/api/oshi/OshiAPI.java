package com.zeydie.minemetrics.common.api.oshi;

import com.zeydie.minemetrics.common.api.oshi.data.CpuData;
import com.zeydie.minemetrics.common.api.oshi.data.GpusData;
import com.zeydie.minemetrics.common.api.oshi.data.OSData;
import com.zeydie.minemetrics.common.api.oshi.data.RamData;
import lombok.Getter;
import lombok.val;
import org.jetbrains.annotations.NotNull;
import oshi.SystemInfo;
import oshi.hardware.HardwareAbstractionLayer;
import oshi.software.os.OperatingSystem;

public final class OshiAPI {
    static {
        init();
    }

    @Getter
    private static OSData os;
    @Getter
    private static CpuData cpu;
    @Getter
    private static GpusData gpus;
    @Getter
    private static RamData ram;

    public static void init() {
        initOS();
        initCPU();
        initGPU();
        initRam();
    }

    public static void update() {
        initRam();
    }

    public static void initOS() {
        val oshiOS = getOshiOS();

        os = new OSData(oshiOS.getBitness(), oshiOS.getFamily() + oshiOS.getVersionInfo().toString());
    }

    public static void initCPU() {
        val oshiCPU = getOshiHardware().getProcessor();

        cpu = new CpuData(
                oshiCPU.getLogicalProcessorCount(),
                oshiCPU.getPhysicalProcessorCount()
        );
    }

    public static void initGPU() {
        gpus = new GpusData(getOshiHardware().getGraphicsCards());
    }

    public static void initRam() {
        val oshiRam = getOshiHardware().getMemory();

        ram = new RamData(
                oshiRam.getTotal(),
                oshiRam.getAvailable()
        );
    }

    private static @NotNull OperatingSystem getOshiOS() {
        return new SystemInfo().getOperatingSystem();
    }

    private static @NotNull HardwareAbstractionLayer getOshiHardware() {
        return new SystemInfo().getHardware();
    }
}