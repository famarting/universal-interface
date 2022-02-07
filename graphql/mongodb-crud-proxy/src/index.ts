#!/usr/bin/env node
require('dotenv').config();
import { createMongoDbProvider } from "@graphback/runtime-mongo";
import { ApolloServer } from "apollo-server-express";
import cors from "cors";
import express from "express";
import { buildGraphbackAPI, GraphbackAPI } from "graphback";
import http from 'http';
import { Db, MongoClient } from "mongodb";
import { loadModel } from "./loadModel";

const connectDB =  async (): Promise<Db> => {
  let url: string

  const { DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE, DB_AUTHSOURCE } = process.env;

  if (!DB_HOST || !DB_PORT || !DB_DATABASE) {
    throw new Error("DB connection info is missing")
  }

  if (DB_USER && DB_PASSWORD) {
    url = `mongodb://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?authSource=${DB_AUTHSOURCE}`;
  } else {
    url = `mongodb://${DB_HOST}:${DB_PORT}/${DB_DATABASE}`
  }

  const mongoClient = await MongoClient.connect(url, { useUnifiedTopology: true });

  // get db instance
  const db = mongoClient.db(process.env.DB_DATABASE);

  return db;
}

async function start() {

  const { GRAPHQL_MODEL } = process.env;
  const model = loadModel(`${GRAPHQL_MODEL}`);

  const db = await connectDB()
  const dataProvider = createMongoDbProvider(db)

  const graphbackApi: GraphbackAPI = buildGraphbackAPI(model, { dataProviderCreator: dataProvider });

  const app = express();

  app.use(cors());

  const apolloConfig = {
    typeDefs: graphbackApi.typeDefs,
    resolvers: graphbackApi.resolvers,
    // eslint-disable-next-line @typescript-eslint/unbound-method
    context: graphbackApi.contextCreator,
    playground: true,
    resolverValidationOptions: {
      requireResolversForResolveType: false
    }
  }

  const apolloServer = new ApolloServer(apolloConfig);

  apolloServer.applyMiddleware({ app });

  const httpServer = http.createServer(app);
  apolloServer.installSubscriptionHandlers(httpServer);
  
  httpServer.listen({ port: 7878 }, () => {
    console.log(`🚀  Server ready at http://localhost:7878/graphql`);
  });

}

start().catch((err: any) => console.log(err));

