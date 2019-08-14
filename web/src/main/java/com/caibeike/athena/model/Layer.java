package com.caibeike.athena.model;

public class Layer {

    private String name;

    private Long layerId;

    @Override
    public String toString() {
        return "{" +
                "name: '" + name + '\'' +
                ", layerId: " + layerId +
                '}';
    }
}
