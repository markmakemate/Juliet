package com.caibeike.athena.service;


import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target({ElementType.METHOD, ElementType.FIELD})
@Retention(RetentionPolicy.RUNTIME)
public @interface RpcService {
    /**
     * 声明一个rpc服务，并指定调用的方法的路径
     * @return
     */
    String value() default "";
}
