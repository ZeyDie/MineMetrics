package com.zeydie.minemetrics.common.managers;

import dev.architectury.injectables.annotations.ExpectPlatform;

public final class EnvironmentManager {
    @ExpectPlatform
    public static boolean isClient() {
        throw new AssertionError("Not implemented!");
    }

    @ExpectPlatform
    public static boolean isServer() {
        throw new AssertionError("Not implemented!");
    }
}