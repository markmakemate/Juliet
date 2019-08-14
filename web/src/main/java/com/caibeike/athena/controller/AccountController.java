package com.caibeike.athena.controller;


import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController(value = "/athena")
public class AccountController {

    @RequestMapping("/login")
    public String Login(Map<String, Object> reqMap){
        return null;
    }
}
