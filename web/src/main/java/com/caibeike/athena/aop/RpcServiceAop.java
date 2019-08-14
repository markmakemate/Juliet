package com.caibeike.athena.aop;


import com.caibeike.athena.service.RpcService;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.After;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Pointcut;
import org.springframework.stereotype.Component;

/**
 * Rpc服务切面实现
 */
@Aspect
@Component
public class RpcServiceAop {

    @Pointcut(value = "@annotation(rpcService)", argNames = "rpcService")
    public void pointcut(RpcService rpcService){
    }

    @AfterReturning(value = "@annotaion(com.caibeike.athena.service.RpcService)",
            returning = "data")
    public void after(JoinPoint joinPoint, String data){

    }

}
