package io.uninterface.examples;

import javax.inject.Inject;
import javax.ws.rs.Consumes;
import javax.ws.rs.DELETE;
import javax.ws.rs.GET;
import javax.ws.rs.NotFoundException;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

import io.uninterface.examples.gqlcrud.MutationExecutor;
import io.uninterface.examples.gqlcrud.QueryExecutor;

@Path("/books")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class BooksResource {

    @Inject
    MutationExecutor booksMutations;
    
    @Inject
    QueryExecutor booksQueries;

    @POST
    public Book create(Book book) {
        // return booksStore.create createDocument(book);
        return null;
    }

    @GET
    @Path("{id}")
    public Book get(@PathParam("id") String id) {
        try {
            return booksQueries.getBook(queryResponseDef, id, paramsAndValues);
        } catch (DocumentStoreException e) {
            if (e.getStatusCode() == 404) {
                throw new NotFoundException("Book not found");
            }
            throw e;
        }
    }

    @DELETE
    @Path("{id}")
    public void delete(@PathParam("id") String id) {
        booksStore.deleteDocument(id);
    }

}
