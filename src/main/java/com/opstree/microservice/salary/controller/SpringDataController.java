package com.opstree.microservice.salary.controller;

import com.opstree.microservice.salary.entity.SalaryDef;
import com.opstree.microservice.salary.service.SpringDataSalaryService;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
@RequestMapping("/salary")
@RequiredArgsConstructor
public class SpringDataController {

    private final SpringDataSalaryService springDataSalaryService;

    @GetMapping("/search/all")
    public List<SalaryDef> findAllSalary() {
        return springDataSalaryService.getSalary();
    }
}
