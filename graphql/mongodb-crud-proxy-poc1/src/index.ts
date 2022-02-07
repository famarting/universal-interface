// eslint-disable-next-line @typescript-eslint/no-require-imports
require('dotenv').config();
import { ApolloServer } from 'apollo-server-express';
import { buildGraphbackAPI } from 'graphback';
import { createMongoDbProvider } from '@graphback/runtime-mongo';
import cors from 'cors';
import express from 'express';
import http from 'http';
import { loadConfigSync } from 'graphql-config';
import { connectDB } from './db';

async function start() {
  const app = express();

  app.use(cors());

  const { CONFIG_FILE } = process.env;

  const graphbackExtension = 'graphback';
  const config = loadConfigSync({
    filepath: CONFIG_FILE,
    extensions: [
      () => ({
        name: graphbackExtension
      })
    ]
  });

  const projectConfig = config.getDefault();
  const graphbackConfig = projectConfig.extension(graphbackExtension);

  const modelDefs = projectConfig.loadSchemaSync(graphbackConfig.model);

  const db = await connectDB();

  const { typeDefs, resolvers, contextCreator } = buildGraphbackAPI(modelDefs, {
    dataProviderCreator: createMongoDbProvider(db)
  });

  const apolloServer = new ApolloServer({
    typeDefs,
    resolvers: [resolvers],
    context: contextCreator
  });

  apolloServer.applyMiddleware({ app });

  const httpServer = http.createServer(app);
  apolloServer.installSubscriptionHandlers(httpServer);

  httpServer.listen({ port: 4000 }, () => {
    console.log(`🚀  Server ready at http://localhost:4000/graphql`);
  });
}

start().catch((err: any) => console.log(err));
