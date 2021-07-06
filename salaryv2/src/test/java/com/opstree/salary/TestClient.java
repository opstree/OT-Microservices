package com.opstree.salary;

import java.io.IOException;
import java.net.InetAddress;
import java.net.UnknownHostException;

import org.apache.http.HttpHost;
import org.elasticsearch.action.get.GetRequest;
import org.elasticsearch.action.get.GetResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestClient;
import org.elasticsearch.client.RestHighLevelClient;
import org.elasticsearch.client.transport.TransportClient;
import org.elasticsearch.common.settings.Settings;
import org.elasticsearch.common.transport.TransportAddress;
import org.junit.jupiter.api.Test;


@SuppressWarnings({ "deprecation", "resource" })
public class TestClient {

    private static String CLUSTER_NAME = "elasticsearch";
    private static String HOST_IP = "172.17.0.3";
    private static int TCP_PORT = 9300;

    @Test
    public void testRestClient() throws Exception {
        RestHighLevelClient client = new RestHighLevelClient(RestClient.builder(new HttpHost(HOST_IP, 9200, "http")));
        GetRequest getRequest = new GetRequest("books", "1");
        GetResponse getResponse = client.get(getRequest, RequestOptions.DEFAULT);
        System.out.println(getResponse.getSourceAsString());
    }
}
