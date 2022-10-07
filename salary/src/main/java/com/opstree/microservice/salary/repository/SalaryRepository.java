package com.opstree.microservice.salary.repository;

import com.opstree.microservice.salary.entity.SalaryDef;
import java.util.List;
import org.springframework.data.elasticsearch.annotations.Query;
import org.springframework.data.elasticsearch.repository.ElasticsearchRepository;

public interface SalaryRepository extends ElasticsearchRepository<SalaryDef, String> {

    @Query("{\"match_all\":{}}")
    List<SalaryDef> findAllSalary();
}
