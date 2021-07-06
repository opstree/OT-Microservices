package com.opstree.salary.service;

import com.opstree.salary.model.SalaryModel;
import com.fasterxml.jackson.databind.ObjectMapper;
import lombok.extern.slf4j.Slf4j;
import org.apache.lucene.search.join.ScoreMode;
import org.elasticsearch.action.bulk.BulkItemResponse;
import org.elasticsearch.action.bulk.BulkRequest;
import org.elasticsearch.action.bulk.BulkResponse;
import org.elasticsearch.action.delete.DeleteRequest;
import org.elasticsearch.action.delete.DeleteResponse;
import org.elasticsearch.action.get.GetRequest;
import org.elasticsearch.action.get.GetResponse;
import org.elasticsearch.action.index.IndexRequest;
import org.elasticsearch.action.index.IndexResponse;
import org.elasticsearch.action.search.SearchRequest;
import org.elasticsearch.action.search.SearchResponse;
import org.elasticsearch.action.update.UpdateRequest;
import org.elasticsearch.action.update.UpdateResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestHighLevelClient;

import org.elasticsearch.common.xcontent.XContentType;
import org.elasticsearch.index.query.MatchQueryBuilder;
import org.elasticsearch.index.query.Operator;
import org.elasticsearch.index.query.QueryBuilder;
import org.elasticsearch.index.query.QueryBuilders;
import org.elasticsearch.search.SearchHit;
import org.elasticsearch.search.builder.SearchSourceBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.*;

import static com.opstree.salary.util.Constant.INDEX;
import static com.opstree.salary.util.Constant.TYPE;

@Service
@Slf4j
public class SalaryService {


    private RestHighLevelClient client;


    private ObjectMapper objectMapper;

    @Autowired
    public SalaryService(RestHighLevelClient client, ObjectMapper objectMapper) {
        this.client = client;
        this.objectMapper = objectMapper;
    }

    public String createSalaryModel(SalaryModel document) throws Exception {

        UUID uuid = UUID.randomUUID();
        document.setId(uuid.toString());

        IndexRequest indexRequest = new IndexRequest(INDEX, TYPE, document.getId())
                .source(convertSalaryModelToMap(document), XContentType.JSON);

        IndexResponse indexResponse = client.index(indexRequest, RequestOptions.DEFAULT);
        return indexResponse.getResult().name();
    }

    public SalaryModel findById(String id) throws Exception {

            GetRequest getRequest = new GetRequest(INDEX, TYPE, id);

            GetResponse getResponse = client.get(getRequest, RequestOptions.DEFAULT);
            Map<String, Object> resultMap = getResponse.getSource();

            return convertMapToSalaryModel(resultMap);

    }



    public String updateProfile(SalaryModel document) throws Exception {

            SalaryModel resultDocument = findById(document.getId());

            UpdateRequest updateRequest = new UpdateRequest(
                    INDEX,
                    TYPE,
                    resultDocument.getId());

            updateRequest.doc(convertSalaryModelToMap(document));
            UpdateResponse updateResponse = client.update(updateRequest, RequestOptions.DEFAULT);

            return updateResponse
                    .getResult()
                    .name();

    }

    public String bulkUpdateProfile(List<SalaryModel> documents) throws Exception {
        List<SalaryModel> documentsFake = new ArrayList<>();
        if(documents == null){
            for(int i = 0; i < 10; i++){
                SalaryModel mySalaryModel = new SalaryModel();
                String id = UUID.randomUUID().toString();
                mySalaryModel.setId(id);
                mySalaryModel.setName(String.format("Name %s", id));
                mySalaryModel.setMonth(String.format("Month %s", id));
            }
        }

        BulkRequest bulkRequest = new BulkRequest();

        documentsFake.forEach(salaryModel -> {
            IndexRequest indexRequest = new IndexRequest(INDEX,TYPE,salaryModel.getId()).
                    source(objectMapper.convertValue(salaryModel, Map.class));

            bulkRequest.add(indexRequest);
        });

        BulkResponse bulkResponse = client.bulk(bulkRequest, RequestOptions.DEFAULT);
        if(bulkResponse.hasFailures()) {
            for (BulkItemResponse bulkItemResponse : bulkResponse) {
                if (bulkItemResponse.isFailed()) {
                    BulkItemResponse.Failure failure = bulkItemResponse.getFailure();
                    System.out.println("Error "+failure.toString());
                }
            }
        }

        return null;
    }


    public List<SalaryModel> findAll() throws Exception {


        SearchRequest searchRequest = buildSearchRequest(INDEX,TYPE);
        SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
        searchSourceBuilder.query(QueryBuilders.matchAllQuery());
        searchRequest.source(searchSourceBuilder);

        SearchResponse searchResponse =
                client.search(searchRequest, RequestOptions.DEFAULT);

        return getSearchResult(searchResponse);
    }


    public List<SalaryModel> findProfileByName(String name) throws Exception{


        SearchRequest searchRequest = new SearchRequest();
        searchRequest.indices(INDEX);
        searchRequest.types(TYPE);

        SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();

        MatchQueryBuilder matchQueryBuilder = QueryBuilders
                .matchQuery("name",name)
                .operator(Operator.AND);

        searchSourceBuilder.query(matchQueryBuilder);

        searchRequest.source(searchSourceBuilder);

        SearchResponse searchResponse =
                client.search(searchRequest, RequestOptions.DEFAULT);

        return getSearchResult(searchResponse);

    }


    public String deleteSalaryModel(String id) throws Exception {

        DeleteRequest deleteRequest = new DeleteRequest(INDEX, TYPE, id);
        DeleteResponse response = client.delete(deleteRequest,RequestOptions.DEFAULT);

        return response
                .getResult()
                .name();

    }

    private Map<String, Object> convertSalaryModelToMap(SalaryModel salaryModel) {
        return objectMapper.convertValue(salaryModel, Map.class);
    }

    private SalaryModel convertMapToSalaryModel(Map<String, Object> map){
        return objectMapper.convertValue(map,SalaryModel.class);
    }


    public List<SalaryModel> searchByTechnology(String technology) throws Exception{

        SearchRequest searchRequest = buildSearchRequest(INDEX,TYPE);
        SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();

        QueryBuilder queryBuilder = QueryBuilders
                .boolQuery()
                .must(QueryBuilders
                        .matchQuery("technologies.name",technology));

        searchSourceBuilder.query(QueryBuilders.nestedQuery("technologies",queryBuilder,ScoreMode.Avg));

        searchRequest.source(searchSourceBuilder);

        SearchResponse response = client.search(searchRequest,RequestOptions.DEFAULT);

        return getSearchResult(response);
    }

    private List<SalaryModel> getSearchResult(SearchResponse response) {

        SearchHit[] searchHit = response.getHits().getHits();

        List<SalaryModel> salaryModel = new ArrayList<>();

        for (SearchHit hit : searchHit){
            salaryModel
                    .add(objectMapper
                            .convertValue(hit
                                    .getSourceAsMap(), SalaryModel.class));
        }

        return salaryModel;
    }

    private SearchRequest buildSearchRequest(String index, String type) {

        SearchRequest searchRequest = new SearchRequest();
        searchRequest.indices(index);
        searchRequest.types(type);

        return searchRequest;
    }
}
