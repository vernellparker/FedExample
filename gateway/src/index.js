import { ApolloServer } from 'apollo-server-express';
import { ApolloGateway,RemoteGraphQLDataSource  } from "@apollo/gateway";
import { applyMiddleware }  from "graphql-middleware";
import express from "express";
import expressJwt from"express-jwt";
import  permissions  from "./permissions";

const port = 4000;
const app = express();

app.use(
    expressJwt({
        secret: "secret",
        algorithms: ["HS256"],
        credentialsRequired: false
    })
);

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'bolt-wrapper', url: 'http://localhost:4001/query' },
        { name: 'authorization', url: 'http://localhost:4002/query' },
    ],
    buildService({ name, url }) {
        return new RemoteGraphQLDataSource({
            url,
            willSendRequest({request, context}) {
                request.http.headers.set(
                    "user",
                    context.user ? JSON.stringify(context.user) : null
                );
            }
        });
    }
    // buildService({ url}) {
    //     return new RemoteGraphQLDataSource({
    //         url,
    //         willSendRequest({ request, context }) {
    //             // Only add the token if a token exists
    //             if(context.token) {
    //                 request.http.headers.set('Authorization', context.token);
    //             }
    //         },
    //     });
    // },

});

const server = new ApolloServer({
    schema: applyMiddleware(
        permissions
    ),
    gateway,
    subscriptions: false,
    context: ({ req }) => {
        const user = req.user || null;
        return { user };
    }
});

server.applyMiddleware({ app });
// (async () => {
//     const {schema, executor} = await gateway.load();
//
//     const server = new ApolloServer({
//         schema,
//         executor,
//         context: ({ req }) => {
//             const token = req.headers.authorization || null;
//             return {token: token}
//         }
//     });
//
//     server.listen().then(({url}) => {
//         console.log(`🚀 Server ready at ${url}`);
//     });
// })();

app.listen({ port }, () =>
    console.log(`Server ready at http://localhost:${port}${server.graphqlPath}`)
);