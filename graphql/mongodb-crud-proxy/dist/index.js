#!/usr/bin/env node
"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const tslib_1 = require("tslib");
require('dotenv').config();
const runtime_mongo_1 = require("@graphback/runtime-mongo");
const apollo_server_express_1 = require("apollo-server-express");
const cors_1 = require("cors");
const express_1 = require("express");
const graphback_1 = require("graphback");
const http_1 = require("http");
const mongodb_1 = require("mongodb");
const loadModel_1 = require("./loadModel");
const connectDB = () => tslib_1.__awaiter(void 0, void 0, void 0, function* () {
    let url;
    const { DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_DATABASE, DB_AUTHSOURCE } = process.env;
    if (!DB_HOST || !DB_PORT || !DB_DATABASE) {
        throw new Error("DB connection info is missing");
    }
    if (DB_USER && DB_PASSWORD) {
        url = `mongodb://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?authSource=${DB_AUTHSOURCE}`;
    }
    else {
        url = `mongodb://${DB_HOST}:${DB_PORT}/${DB_DATABASE}`;
    }
    const mongoClient = yield mongodb_1.MongoClient.connect(url, { useUnifiedTopology: true });
    // get db instance
    const db = mongoClient.db(process.env.DB_DATABASE);
    return db;
});
function start() {
    return tslib_1.__awaiter(this, void 0, void 0, function* () {
        const { GRAPHQL_MODEL } = process.env;
        const model = loadModel_1.loadModel(GRAPHQL_MODEL);
        const db = yield connectDB();
        const dataProvider = runtime_mongo_1.createMongoDbProvider(db);
        const graphbackApi = graphback_1.buildGraphbackAPI(model, { dataProviderCreator: dataProvider });
        const app = express_1.default();
        app.use(cors_1.default());
        const apolloConfig = {
            typeDefs: graphbackApi.typeDefs,
            resolvers: graphbackApi.resolvers,
            // eslint-disable-next-line @typescript-eslint/unbound-method
            context: graphbackApi.contextCreator,
            playground: true,
            resolverValidationOptions: {
                requireResolversForResolveType: false
            }
        };
        const apolloServer = new apollo_server_express_1.ApolloServer(apolloConfig);
        apolloServer.applyMiddleware({ app });
        const httpServer = http_1.default.createServer(app);
        apolloServer.installSubscriptionHandlers(httpServer);
        httpServer.listen({ port: 4000 }, () => {
            console.log(`🚀  Server ready at http://localhost:4000/graphql`);
        });
    });
}
start().catch((err) => console.log(err));
//# sourceMappingURL=index.js.map