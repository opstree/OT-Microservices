 
package com.opstree.gateway;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.netflix.zuul.EnableZuulProxy;
import org.springframework.context.annotation.Bean;
import com.opstree.gateway.filters.pre.SimpleFilter;
import co.elastic.apm.attach.ElasticApmAttacher;

@SpringBootApplication
@EnableZuulProxy

public class GatewayServiceApplication {
	public static void main(String[] args) {
		ElasticApmAttacher.attach();
		SpringApplication.run(GatewayServiceApplication.class, args);
	}

	@Bean
	public SimpleFilter simpleFilter() {
		return new SimpleFilter();
	}
}
