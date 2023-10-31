import { AppProps } from "next/app";
import "../styles/globals.css";
import {
  ApolloClient,
  ApolloProvider,
  InMemoryCache,
  createHttpLink,
} from "@apollo/client";

function MyApp({ Component, pageProps }: AppProps) {
  const link = createHttpLink({
    uri: "http://localhost:8080/query",
    credentials: "include",
  });

  const client = new ApolloClient({
    cache: new InMemoryCache(),
    link,
  });

  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />;
    </ApolloProvider>
  );
}

export default MyApp;
