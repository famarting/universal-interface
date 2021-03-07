package io.uninterface.client;

import java.io.IOException;
import java.io.InputStream;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpRequest.BodyPublishers;
import java.net.http.HttpResponse;
import java.net.http.HttpResponse.BodyHandlers;

import com.fasterxml.jackson.databind.ObjectMapper;

/**
 * Hello world!
 */
public final class DocumentStoreClient {
    
    private static final String API_BASE_PATH = "api/documentstores/v1/";

    private HttpClient client;
    private ObjectMapper mapper;

    private String baseUrl;
    private String idField;

    public DocumentStoreClient(String endpoint, String storeName, String idField) {
        if (endpoint == null) {
            endpoint = "http://localhost:7878";
        }
        if (!endpoint.endsWith("/")) {
            endpoint += "/";
        }
        this.baseUrl = endpoint + API_BASE_PATH + storeName;
        this.idField = idField;
        this.client = HttpClient.newHttpClient();
        this.mapper = new ObjectMapper();
    }

    public <T> T createDocument(T document) {
        try {
            HttpRequest req = HttpRequest.newBuilder()
                    .uri(URI.create(baseUrl))
                    .header("Content-Type", "application/json")
                    .POST(BodyPublishers.ofByteArray(this.mapper.writeValueAsBytes(document)))
                    .build();

            HttpResponse<InputStream> res = client.send(req, BodyHandlers.ofInputStream());
            if (res.statusCode() == 200) {
                return (T) this.mapper.readValue(res.body(), document.getClass());
            }
            throw new DocumentStoreException(res.statusCode(), res.toString());
        } catch ( IOException | InterruptedException e ) {
            throw new DocumentStoreException(e);
        }
    }

    public <T> T getDocument(Object id, Class<T> clazz) {
        HttpRequest req = HttpRequest.newBuilder()
                .uri(URI.create(baseUrl + "/" + id))
                .GET()
                .build();
        try {
            HttpResponse<InputStream> res = client.send(req, BodyHandlers.ofInputStream());
            if (res.statusCode() == 200) {
                return this.mapper.readValue(res.body(), clazz);
            }
            throw new DocumentStoreException(res.statusCode(), res.toString());
        } catch ( IOException | InterruptedException e ) {
            throw new DocumentStoreException(e);
        }
    }

    public void deleteDocument(Object id) {
        HttpRequest req = HttpRequest.newBuilder()
                .uri(URI.create(baseUrl + "/" + id))
                .DELETE()
                .build();
        try {
            HttpResponse<InputStream> res = client.send(req, BodyHandlers.ofInputStream());
            if (res.statusCode() == 200) {
                return;
            }
            throw new DocumentStoreException(res.statusCode(), res.toString());
        } catch ( IOException | InterruptedException e ) {
            throw new DocumentStoreException(e);
        }
    }

}