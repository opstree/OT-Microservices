
package com.opstree.microservice.salary.service;

import com.opstree.microservice.salary.entity.SalaryDef;
import java.util.List;

public interface SpringDataSalaryService {

    List<SalaryDef> getSalary();

}
