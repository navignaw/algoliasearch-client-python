// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

import 'package:algolia_client_core/algolia_client_core.dart';
import 'package:algoliasearch/src/deserialize.dart';
import 'package:algoliasearch/src/version.dart';

import 'package:algoliasearch/src/model/search_method_params.dart';
import 'package:algoliasearch/src/model/search_responses.dart';

final class SearchClient implements ApiClient {
  @override
  final String apiKey;

  @override
  final String appId;

  @override
  final ClientOptions options;

  final RetryStrategy _retryStrategy;

  SearchClient({
    required this.appId,
    required this.apiKey,
    this.options = const ClientOptions(),
  }) : _retryStrategy = RetryStrategy.create(
          segment:
              AgentSegment(value: "Algoliasearch", version: packageVersion),
          appId: appId,
          apiKey: apiKey,
          options: options,
          defaultHosts: () =>
              [
                Host(url: '$appId-dsn.algolia.net', callType: CallType.read),
                Host(url: '$appId.algolia.net', callType: CallType.write),
              ] +
              ([
                Host(url: '$appId-1.algolianet.com'),
                Host(url: '$appId-2.algolianet.com'),
                Host(url: '$appId-3.algolianet.com'),
              ]..shuffle()),
        ) {
    assert(appId.isNotEmpty, '`appId` is missing.');
    assert(apiKey.isNotEmpty, '`apiKey` is missing.');
  }

  /// Send requests to the Algolia REST API.
  /// This method allow you to send requests to the Algolia REST API.
  ///
  /// Parameters:
  /// * [path] Path of the endpoint, anything after \"/1\" must be specified.
  /// * [parameters] Query parameters to apply to the current query.
  /// * [body] Parameters to send with the custom request.
  /// * [requestOptions] additional request configuration.
  Future<Object> post({
    required String path,
    Map<String, Object>? parameters,
    Object? body,
    RequestOptions? requestOptions,
  }) async {
    assert(
      path.isNotEmpty,
      'Parameter `path` is required when calling `post`.',
    );
    final request = ApiRequest(
      method: RequestMethod.post,
      path: r'/1{path}'.replaceAll('{' r'path' '}', path),
      queryParams: {
        ...?parameters,
      },
      body: body,
    );
    final response = await _retryStrategy.execute(
      request: request,
      options: requestOptions,
    );
    return deserialize<Object, Object>(
      response,
      'Object',
      growable: true,
    );
  }

  /// Search multiple indices.
  /// Send multiple search queries to one or more indices.
  ///
  /// Parameters:
  /// * [searchMethodParams] Query requests and strategies. Results will be received in the same order as the queries.
  /// * [requestOptions] additional request configuration.
  Future<SearchResponses> search({
    required SearchMethodParams searchMethodParams,
    RequestOptions? requestOptions,
  }) async {
    final request = ApiRequest(
      method: RequestMethod.post,
      path: r'/1/indexes/*/queries',
      isRead: true,
      body: searchMethodParams.toJson(),
    );
    final response = await _retryStrategy.execute(
      request: request,
      options: requestOptions,
    );
    return deserialize<SearchResponses, SearchResponses>(
      response,
      'SearchResponses',
      growable: true,
    );
  }

  @override
  void dispose() => _retryStrategy.dispose();
}
