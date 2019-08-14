package com.caibeike.athena.config;


import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.stereotype.Component;

@Component
@ConfigurationProperties(prefix = "mybatis")
public class MybatisConfiguration {

}
