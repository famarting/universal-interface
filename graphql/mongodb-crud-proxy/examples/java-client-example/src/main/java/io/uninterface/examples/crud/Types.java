package io.uninterface.examples.crud;

import java.util.List;
import java.util.HashMap;
import java.util.Map;

public class Types {
  
  public static class BookFilterInput {
    private GraphbackObjectIdInput _id;
    private StringInput name;
    private StringInput author;
    private Iterable<BookFilterInput> and;
    private Iterable<BookFilterInput> or;
    private BookFilterInput not;
  
    public BookFilterInput() {}
  
    public GraphbackObjectIdInput get_Id() { return this._id; }
    public StringInput getName() { return this.name; }
    public StringInput getAuthor() { return this.author; }
    public Iterable<BookFilterInput> getAnd() { return this.and; }
    public Iterable<BookFilterInput> getOr() { return this.or; }
    public BookFilterInput getNot() { return this.not; }
    public void set_Id(GraphbackObjectIdInput _id) { this._id = _id; }
    public void setName(StringInput name) { this.name = name; }
    public void setAuthor(StringInput author) { this.author = author; }
    public void setAnd(Iterable<BookFilterInput> and) { this.and = and; }
    public void setOr(Iterable<BookFilterInput> or) { this.or = or; }
    public void setNot(BookFilterInput not) { this.not = not; }
  }
  
  public static class BookSubscriptionFilterInput {
    private Iterable<BookSubscriptionFilterInput> and;
    private Iterable<BookSubscriptionFilterInput> or;
    private BookSubscriptionFilterInput not;
    private GraphbackObjectIdInput _id;
    private StringInput name;
    private StringInput author;
  
    public BookSubscriptionFilterInput() {}
  
    public Iterable<BookSubscriptionFilterInput> getAnd() { return this.and; }
    public Iterable<BookSubscriptionFilterInput> getOr() { return this.or; }
    public BookSubscriptionFilterInput getNot() { return this.not; }
    public GraphbackObjectIdInput get_Id() { return this._id; }
    public StringInput getName() { return this.name; }
    public StringInput getAuthor() { return this.author; }
    public void setAnd(Iterable<BookSubscriptionFilterInput> and) { this.and = and; }
    public void setOr(Iterable<BookSubscriptionFilterInput> or) { this.or = or; }
    public void setNot(BookSubscriptionFilterInput not) { this.not = not; }
    public void set_Id(GraphbackObjectIdInput _id) { this._id = _id; }
    public void setName(StringInput name) { this.name = name; }
    public void setAuthor(StringInput author) { this.author = author; }
  }
  public static class CreateBookInput {
    private String _id;
    private String name;
    private String author;
  
    public CreateBookInput() {}
  
    public String get_Id() { return this._id; }
    public String getName() { return this.name; }
    public String getAuthor() { return this.author; }
    public void set_Id(String _id) { this._id = _id; }
    public void setName(String name) { this.name = name; }
    public void setAuthor(String author) { this.author = author; }
  }
  public static class GraphbackObjectIdInput {
    private String ne;
    private String eq;
    private String le;
    private String lt;
    private String ge;
    private String gt;
    private Iterable<String> in;
    private Iterable<String> between;
  
    public GraphbackObjectIdInput() {}
  
    public String getNe() { return this.ne; }
    public String getEq() { return this.eq; }
    public String getLe() { return this.le; }
    public String getLt() { return this.lt; }
    public String getGe() { return this.ge; }
    public String getGt() { return this.gt; }
    public Iterable<String> getIn() { return this.in; }
    public Iterable<String> getBetween() { return this.between; }
    public void setNe(String ne) { this.ne = ne; }
    public void setEq(String eq) { this.eq = eq; }
    public void setLe(String le) { this.le = le; }
    public void setLt(String lt) { this.lt = lt; }
    public void setGe(String ge) { this.ge = ge; }
    public void setGt(String gt) { this.gt = gt; }
    public void setIn(Iterable<String> in) { this.in = in; }
    public void setBetween(Iterable<String> between) { this.between = between; }
  }
  public static class MutateBookInput {
    private String _id;
    private String name;
    private String author;
  
    public MutateBookInput() {}
  
    public String get_Id() { return this._id; }
    public String getName() { return this.name; }
    public String getAuthor() { return this.author; }
    public void set_Id(String _id) { this._id = _id; }
    public void setName(String name) { this.name = name; }
    public void setAuthor(String author) { this.author = author; }
  }
  public static class MutationCreateBookArgs {
    private CreateBookInput input;
  
    public MutationCreateBookArgs() {}
  
    public CreateBookInput getInput() { return this.input; }
    public void setInput(CreateBookInput input) { this.input = input; }
  }
  public static class MutationUpdateBookArgs {
    private MutateBookInput input;
  
    public MutationUpdateBookArgs() {}
  
    public MutateBookInput getInput() { return this.input; }
    public void setInput(MutateBookInput input) { this.input = input; }
  }
  public static class MutationDeleteBookArgs {
    private MutateBookInput input;
  
    public MutationDeleteBookArgs() {}
  
    public MutateBookInput getInput() { return this.input; }
    public void setInput(MutateBookInput input) { this.input = input; }
  }
  public static class OrderByInput {
    private String field;
    private SortDirectionEnum order;
  
    public OrderByInput() {}
  
    public String getField() { return this.field; }
    public SortDirectionEnum getOrder() { return this.order; }
    public void setField(String field) { this.field = field; }
    public void setOrder(SortDirectionEnum order) { this.order = order; }
  }
  public static class PageRequestInput {
    private Integer limit;
    private Integer offset;
  
    public PageRequestInput() {}
  
    public Integer getLimit() { return this.limit; }
    public Integer getOffset() { return this.offset; }
    public void setLimit(Integer limit) { this.limit = limit; }
    public void setOffset(Integer offset) { this.offset = offset; }
  }
  public static class QueryGetBookArgs {
    private String id;
  
    public QueryGetBookArgs() {}
  
    public String getId() { return this.id; }
    public void setId(String id) { this.id = id; }
  }
  public static class QueryFindBooksArgs {
    private BookFilterInput filter;
    private PageRequestInput page;
    private OrderByInput orderBy;
  
    public QueryFindBooksArgs() {}
  
    public BookFilterInput getFilter() { return this.filter; }
    public PageRequestInput getPage() { return this.page; }
    public OrderByInput getOrderBy() { return this.orderBy; }
    public void setFilter(BookFilterInput filter) { this.filter = filter; }
    public void setPage(PageRequestInput page) { this.page = page; }
    public void setOrderBy(OrderByInput orderBy) { this.orderBy = orderBy; }
  }
  public enum SortDirectionEnum {
    DESC,
    ASC
    
  }
  
  public static class StringInput {
    private String ne;
    private String eq;
    private String le;
    private String lt;
    private String ge;
    private String gt;
    private Iterable<String> in;
    private String contains;
    private String startsWith;
    private String endsWith;
  
    public StringInput() {}
  
    public String getNe() { return this.ne; }
    public String getEq() { return this.eq; }
    public String getLe() { return this.le; }
    public String getLt() { return this.lt; }
    public String getGe() { return this.ge; }
    public String getGt() { return this.gt; }
    public Iterable<String> getIn() { return this.in; }
    public String getContains() { return this.contains; }
    public String getStartsWith() { return this.startsWith; }
    public String getEndsWith() { return this.endsWith; }
    public void setNe(String ne) { this.ne = ne; }
    public void setEq(String eq) { this.eq = eq; }
    public void setLe(String le) { this.le = le; }
    public void setLt(String lt) { this.lt = lt; }
    public void setGe(String ge) { this.ge = ge; }
    public void setGt(String gt) { this.gt = gt; }
    public void setIn(Iterable<String> in) { this.in = in; }
    public void setContains(String contains) { this.contains = contains; }
    public void setStartsWith(String startsWith) { this.startsWith = startsWith; }
    public void setEndsWith(String endsWith) { this.endsWith = endsWith; }
  }
  public static class SubscriptionNewBookArgs {
    private BookSubscriptionFilterInput filter;
  
    public SubscriptionNewBookArgs() {}
  
    public BookSubscriptionFilterInput getFilter() { return this.filter; }
    public void setFilter(BookSubscriptionFilterInput filter) { this.filter = filter; }
  }
  public static class SubscriptionUpdatedBookArgs {
    private BookSubscriptionFilterInput filter;
  
    public SubscriptionUpdatedBookArgs() {}
  
    public BookSubscriptionFilterInput getFilter() { return this.filter; }
    public void setFilter(BookSubscriptionFilterInput filter) { this.filter = filter; }
  }
  public static class SubscriptionDeletedBookArgs {
    private BookSubscriptionFilterInput filter;
  
    public SubscriptionDeletedBookArgs() {}
  
    public BookSubscriptionFilterInput getFilter() { return this.filter; }
    public void setFilter(BookSubscriptionFilterInput filter) { this.filter = filter; }
  }
}
