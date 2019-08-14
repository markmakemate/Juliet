package com.caibeike.athena.model;

public class Domain {

    private String name;

    private Long domainId;

    @Override
    public String toString() {
        return "{" +
                "name: '" + name + '\'' +
                ", domainId: " + domainId +
                '}';
    }
}
