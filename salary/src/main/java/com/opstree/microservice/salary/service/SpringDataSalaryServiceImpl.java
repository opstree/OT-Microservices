package com.opstree.microservice.salary.service;

import com.opstree.microservice.salary.entity.SalaryDef;
import com.opstree.microservice.salary.repository.SalaryRepository;
import java.util.List;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class SpringDataSalaryServiceImpl implements SpringDataSalaryService {

    private final SalaryRepository salaryRepository;


    public List<SalaryDef> getSalary() {
        return salaryRepository.findAllSalary();
    }
}
