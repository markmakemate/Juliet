package com.caibeike.athena.model;


import java.util.Map;

/**
 * 继承Parameter的所有对象都必须实现toString()方法
 */
public interface Parameter {

   void Receive(Map<String, Object> reqMap);

}
