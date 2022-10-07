package  com.opstree.microservice.salary;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.ConfigurationPropertiesScan;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import co.elastic.apm.attach.ElasticApmAttacher;

@SpringBootApplication
@ConfigurationPropertiesScan(" com.opstree.microservice.salary.properties")
public class SalaryApplication {
    public static void main(String[] args) {
        ElasticApmAttacher.attach();
        SpringApplication.run(SalaryApplication.class, args);
    }
}
