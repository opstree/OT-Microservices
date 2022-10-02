package com.opstree.microservice.salary.service;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.opstree.microservice.salary.entity.SalaryDef;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Collections;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import lombok.RequiredArgsConstructor;
import lombok.SneakyThrows;
import lombok.extern.slf4j.Slf4j;
import org.elasticsearch.action.search.SearchRequest;
import org.elasticsearch.action.search.SearchResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestHighLevelClient;
import org.elasticsearch.index.query.IdsQueryBuilder;
import org.elasticsearch.index.query.QueryBuilders;
import org.elasticsearch.search.SearchHit;
import org.elasticsearch.search.builder.SearchSourceBuilder;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
@Slf4j
public class HighLevelClientSalaryServiceImpl implements HighLevelClientSalaryService{
    private final RestHighLevelClient restHighLevelClient;
    private final ObjectMapper objectMapper;

    public SalaryDef getSalary() {
        SearchRequest searchRequest = new SearchRequest("employee-management");
        SearchSourceBuilder sourceBuilder = SearchSourceBuilder.searchSource();
        IdsQueryBuilder idsQueryBuilder = QueryBuilders.idsQuery();
        sourceBuilder.query(idsQueryBuilder);
        searchRequest.source(sourceBuilder);
        try {
            SearchResponse response = restHighLevelClient.search(searchRequest, RequestOptions.DEFAULT);
            return toSalaryList(response.getHits().getHits()).stream().findFirst().orElse(null);
        } catch (Exception e) {
            e.printStackTrace();
        }
        return null;
    }

    private List<SalaryDef> toSalaryList(SearchHit[] searchHits) throws Exception {
        List<SalaryDef> salaryList = new ArrayList<>();
        for (SearchHit searchHit : searchHits) {
            salaryList.add(objectMapper.readValue(searchHit.getSourceAsString(), SalaryDef.class));
        }
        return salaryList;
    }
}
