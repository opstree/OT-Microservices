package com.opstree.microservice.salary.service;

import com.opstree.microservice.salary.entity.SalaryDef;
import java.util.List;
import java.util.Map;
import org.elasticsearch.search.aggregations.Aggregation;

public interface HighLevelClientSalaryService {

    SalaryDef getSalary();
}
