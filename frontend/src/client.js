import { ApolloClient } from 'apollo-client'
import { InMemoryCache } from 'apollo-cache-inmemory';
import { createHttpLink } from "apollo-link-http";

const link = createHttpLink({
    uri: 'http://roboapi.empatter.com/graphql',
    withCredentials: true
})

export const gqlClient = () => {
    return new ApolloClient({
        credentials: 'include',
        link,
        cache: new InMemoryCache()
    })
}
