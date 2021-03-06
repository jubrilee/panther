/**
 * Panther is a Cloud-Native SIEM for the Modern Security Team.
 * Copyright (C) 2020 Panther Labs Inc
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

import * as Types from '../../../../../../__generated__/schema';

import { GraphQLError } from 'graphql';
import gql from 'graphql-tag';
import * as ApolloReactCommon from '@apollo/client';
import * as ApolloReactHooks from '@apollo/client';

export type GetLogCfnTemplateVariables = {
  input: Types.GetS3LogIntegrationTemplateInput;
};

export type GetLogCfnTemplate = {
  getS3LogIntegrationTemplate: Pick<Types.IntegrationTemplate, 'body' | 'stackName'>;
};

export const GetLogCfnTemplateDocument = gql`
  query GetLogCfnTemplate($input: GetS3LogIntegrationTemplateInput!) {
    getS3LogIntegrationTemplate(input: $input) {
      body
      stackName
    }
  }
`;

/**
 * __useGetLogCfnTemplate__
 *
 * To run a query within a React component, call `useGetLogCfnTemplate` and pass it any options that fit your needs.
 * When your component renders, `useGetLogCfnTemplate` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetLogCfnTemplate({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useGetLogCfnTemplate(
  baseOptions?: ApolloReactHooks.QueryHookOptions<GetLogCfnTemplate, GetLogCfnTemplateVariables>
) {
  return ApolloReactHooks.useQuery<GetLogCfnTemplate, GetLogCfnTemplateVariables>(
    GetLogCfnTemplateDocument,
    baseOptions
  );
}
export function useGetLogCfnTemplateLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetLogCfnTemplate, GetLogCfnTemplateVariables>
) {
  return ApolloReactHooks.useLazyQuery<GetLogCfnTemplate, GetLogCfnTemplateVariables>(
    GetLogCfnTemplateDocument,
    baseOptions
  );
}
export type GetLogCfnTemplateHookResult = ReturnType<typeof useGetLogCfnTemplate>;
export type GetLogCfnTemplateLazyQueryHookResult = ReturnType<typeof useGetLogCfnTemplateLazyQuery>;
export type GetLogCfnTemplateQueryResult = ApolloReactCommon.QueryResult<
  GetLogCfnTemplate,
  GetLogCfnTemplateVariables
>;
export function mockGetLogCfnTemplate({
  data,
  variables,
  errors,
}: {
  data: GetLogCfnTemplate;
  variables?: GetLogCfnTemplateVariables;
  errors?: GraphQLError[];
}) {
  return {
    request: { query: GetLogCfnTemplateDocument, variables },
    result: { data, errors },
  };
}
