package io.uninterface.client;

public class DocumentStoreException extends RuntimeException {
   
    private static final long serialVersionUID = -2411745048885757263L;
    
    private int statusCode;

    public DocumentStoreException(int statusCode, String message) {
        super(message);
        this.statusCode = statusCode;
    }

    public DocumentStoreException(String message, Throwable cause) {
        super(message, cause);
    }

    public DocumentStoreException(Throwable cause) {
        super(cause);
    }

    public int getStatusCode() {
        return statusCode;
    }

}
