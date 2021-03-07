package io.uninterface.examples;

import javax.enterprise.inject.Produces;

import com.fasterxml.jackson.databind.JavaType;

import io.uninterface.client.DocumentStoreClient;

public class DocumentStoreClientProducer {
    
    @Produces
    public DocumentStoreClient client() {
        return new DocumentStoreClient(null, "example", null);
    }

}
