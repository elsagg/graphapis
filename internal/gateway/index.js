const { ApolloServer } = require("apollo-server");
const { ApolloGateway } = require("@apollo/gateway");

const gateway = new ApolloGateway({
  serviceList: [
    { name: "authors", url: "http://localhost:8181/query" },
    { name: "books", url: "http://localhost:8182/query" },
  ],
});

const server = new ApolloServer({
  gateway,
  subscriptions: false,
});

server.listen(8180).then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`);
});
