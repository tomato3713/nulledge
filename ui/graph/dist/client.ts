import { GraphQLClient } from 'graphql-request';
import { GraphQLClientRequestHeaders } from 'graphql-request/build/cjs/types';
import gql from 'graphql-tag';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export enum MarkupLanguage {
  Asciidoc = 'ASCIIDOC',
  Markdown = 'MARKDOWN'
}

export type Mutation = {
  __typename?: 'Mutation';
  createPage: Page;
  createUser: User;
};


export type MutationCreatePageArgs = {
  input: NewPageInput;
};


export type MutationCreateUserArgs = {
  input: NewUserInput;
};

export type NewPageInput = {
  text: Scalars['String']['input'];
  userId: Scalars['String']['input'];
};

export type NewUserInput = {
  name: Scalars['String']['input'];
};

export type Page = {
  __typename?: 'Page';
  format: MarkupLanguage;
  id: Scalars['ID']['output'];
  text: Scalars['String']['output'];
  user: User;
};

export type Query = {
  __typename?: 'Query';
  page?: Maybe<Page>;
  pages: Array<Page>;
  user?: Maybe<User>;
  users: Array<User>;
};


export type QueryPageArgs = {
  id: Scalars['ID']['input'];
};


export type QueryUserArgs = {
  id: Scalars['ID']['input'];
};

export type User = {
  __typename?: 'User';
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type GetUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type GetUsersQuery = { __typename?: 'Query', users: Array<{ __typename?: 'User', id: string, name: string }> };

export type GetPagesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetPagesQuery = { __typename?: 'Query', pages: Array<{ __typename?: 'Page', id: string, text: string, format: MarkupLanguage, user: { __typename?: 'User', id: string, name: string } }> };


export const GetUsersDocument = gql`
    query getUsers {
  users {
    id
    name
  }
}
    `;
export const GetPagesDocument = gql`
    query getPages {
  pages {
    id
    text
    format
    user {
      id
      name
    }
  }
}
    `;

export type SdkFunctionWrapper = <T>(action: (requestHeaders?:Record<string, string>) => Promise<T>, operationName: string, operationType?: string) => Promise<T>;


const defaultWrapper: SdkFunctionWrapper = (action, _operationName, _operationType) => action();

export function getSdk(client: GraphQLClient, withWrapper: SdkFunctionWrapper = defaultWrapper) {
  return {
    getUsers(variables?: GetUsersQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<GetUsersQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetUsersQuery>(GetUsersDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getUsers', 'query');
    },
    getPages(variables?: GetPagesQueryVariables, requestHeaders?: GraphQLClientRequestHeaders): Promise<GetPagesQuery> {
      return withWrapper((wrappedRequestHeaders) => client.request<GetPagesQuery>(GetPagesDocument, variables, {...requestHeaders, ...wrappedRequestHeaders}), 'getPages', 'query');
    }
  };
}
export type Sdk = ReturnType<typeof getSdk>;