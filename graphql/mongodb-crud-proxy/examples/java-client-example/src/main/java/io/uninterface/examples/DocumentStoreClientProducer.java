package io.uninterface.examples;

import javax.enterprise.inject.Produces;

import io.uninterface.examples.gqlcrud.MutationExecutor;
import io.uninterface.examples.gqlcrud.QueryExecutor;

public class DocumentStoreClientProducer {
    
    @Produces
    public MutationExecutor client() {
        Boo
        return new MutationExecutor("http://localhost:7878");
    }

    @Produces
    public QueryExecutor queries() {
        return new QueryExecutor("http://localhost:7878");
    }
}
