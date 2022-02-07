package documentstore

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/famartinrh/uninterface/pkg/config"
	"github.com/google/uuid"
	"gocloud.dev/docstore"

	// Import the docstore driver packages we want to be able to open.
	_ "gocloud.dev/docstore/awsdynamodb"
	_ "gocloud.dev/docstore/gcpfirestore"
	_ "gocloud.dev/docstore/memdocstore"
	_ "gocloud.dev/docstore/mongodocstore"
)

// type Message struct {
// 	ID               string // unique ID of each document
// 	Date             string
// 	Content          string
// 	DocstoreRevision interface{}
// }

// func DoSomething() error {

// 	ctx := context.TODO()

// 	collectionURL := "mem://sample-coll/ID?filename=save"

// 	// Open a *docstore.Collection using the collectionURL.
// 	collection, err := docstore.OpenCollection(ctx, collectionURL)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to open collection: %v\n", err)
// 		return err
// 	}
// 	defer collection.Close()

// 	id := uuid.New().String()
// 	date := time.Now().Format("2006-01-02")
// 	content := "abcde"
// 	msg := &Message{
// 		ID:      id,
// 		Date:    date,
// 		Content: content,
// 	}
// 	if err := collection.Put(ctx, msg); err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to put message: %v\n", err)
// 		return err
// 	}
// 	fmt.Printf("Put message: %s\n", msg)

// 	q := collection.Query()
// 	q = q.Where("Date", "=", date).Where("Content", "=", "abcde")
// 	iter := q.Get(ctx)
// 	defer iter.Stop()
// 	for {
// 		var msg Message
// 		err := iter.Next(ctx, &msg)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed to list: %v\n", err)
// 			return err
// 		}
// 		fmt.Println("Doc query")
// 		fmt.Println(msg)
// 	}

// 	getdoc := &Message{ID: id}
// 	err = collection.Get(ctx, getdoc, "Content")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to GET: %v\n", err)
// 		return err
// 	}
// 	fmt.Println("Doc get")
// 	fmt.Println(getdoc)

// 	msg = &Message{ID: id}
// 	mods := docstore.Mods{"Content": "testupdate"}
// 	if errs := collection.Actions().Update(msg, mods).Get(msg).Do(ctx); errs != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to update message: %v\n", errs)
// 		return errs
// 	}
// 	fmt.Printf("updated: %s\n", msg)

// 	q = collection.Query()
// 	q = q.Where("Date", "=", date)
// 	iter = q.Get(ctx, "ID")
// 	dels := collection.Actions()
// 	for {
// 		var msg Message
// 		err := iter.Next(ctx, &msg)
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		dels.Delete(&msg)
// 	}
// 	if err := dels.Do(ctx); err != nil {
// 		fmt.Fprintf(os.Stderr, "Failed to delete: %v\n", err)
// 		return err
// 	}
// 	return nil
// }

type DocumentStore struct {
	name       string
	IDField    string
	collection *docstore.Collection
}

type Query struct {
	Conditions       []queryCondition `json:"conditions,omitempty"`
	Limit            int              `json:"limit,omitempty"`
	OrderByField     string           `json:"orderBy"`
	OrderByDirection string           `json:"sort"`
}

type queryCondition struct {
	field string      `json:"field"`
	op    string      `json:"op"`
	value interface{} `json:"value"`
}

func NewDocumentStore(config config.Docstore) (*DocumentStore, error) {
	if config.ConnectionString != "" {
		os.Setenv("MONGO_SERVER_URL", config.ConnectionString)
	}
	collection, err := docstore.OpenCollection(context.TODO(), config.URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open collection: %v\n", err)
		return nil, err
	}
	if config.IDField == "" {
		return nil, errors.New("idField is required")
	}
	return &DocumentStore{
		name:       config.Name,
		IDField:    config.IDField,
		collection: collection,
	}, nil
}

func (store DocumentStore) Create(document map[string]interface{}) error {
	idValue, exists := document[store.IDField]
	idProvided := exists && idValue != nil
	if !idProvided {
		fmt.Println("Auto generating id")
		document[store.IDField] = uuid.New().String()
	}
	return store.collection.Create(context.TODO(), document)
}

func (store DocumentStore) Get(document map[string]interface{}) error {
	return store.collection.Get(context.TODO(), document)
}

func (store DocumentStore) Delete(document map[string]interface{}) error {
	return store.collection.Delete(context.TODO(), document)
}

func (store DocumentStore) Query(query Query) error {
	collection := store.collection
	q := collection.Query()

	for _, cond := range query.Conditions {
		q = q.Where(docstore.FieldPath(cond.field), cond.op, cond.value)
	}

	limit := 100
	if query.Limit != 0 {
		limit = query.Limit
	}
	q = q.Limit(limit)

	orderByField := store.IDField
	if query.OrderByField != "" {
		orderByField = query.OrderByField
	}
	orderByDirection := "asc"
	if query.OrderByDirection != "" {
		orderByDirection = query.OrderByDirection
	}
	q = q.OrderBy(orderByField, orderByDirection)

	iter := q.Get(context.TODO())
	iter.

	defer iter.Stop()
	for {
		var msg Message
		err := iter.Next(ctx, &msg)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to list: %v\n", err)
			return err
		}
		fmt.Println("Doc query")
		fmt.Println(msg)
	}
}
