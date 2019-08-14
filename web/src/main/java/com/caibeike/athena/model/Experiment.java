package com.caibeike.athena.model;

public class Experiment {

    private String name;

    private Long exptId;

    private Parameter parameter;


    @Override
    public String toString() {
        return "{" +
                "name:'" + name + '\'' +
                ", exptId: " + exptId + "," + "param: " +
                parameter.toString() + '}';
    }
}
