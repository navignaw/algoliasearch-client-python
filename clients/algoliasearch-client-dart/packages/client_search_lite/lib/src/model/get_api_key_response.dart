// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_search_lite/src/model/acl.dart';

import 'package:json_annotation/json_annotation.dart';

part 'get_api_key_response.g.dart';

@JsonSerializable()
final class GetApiKeyResponse {
  /// Returns a new [GetApiKeyResponse] instance.
  const GetApiKeyResponse({
    this.value,
    required this.createdAt,
    required this.acl,
    this.description,
    this.indexes,
    this.maxHitsPerQuery,
    this.maxQueriesPerIPPerHour,
    this.queryParameters,
    this.referers,
    this.validity,
  });

  /// The API key.
  @JsonKey(name: r'value')
  final String? value;

  /// Time of the event expressed in milliseconds since the Unix epoch.
  @JsonKey(name: r'createdAt')
  final int createdAt;

  /// Set of permissions associated with the key.
  @JsonKey(name: r'acl')
  final List<Acl> acl;

  /// A comment used to identify a key more easily in the dashboard. It is not interpreted by the API.
  @JsonKey(name: r'description')
  final String? description;

  /// Restrict this new API key to a list of indices or index patterns. If the list is empty, all indices are allowed.
  @JsonKey(name: r'indexes')
  final List<String>? indexes;

  /// Maximum number of hits this API key can retrieve in one query. If zero, no limit is enforced.
  @JsonKey(name: r'maxHitsPerQuery')
  final int? maxHitsPerQuery;

  /// Maximum number of API calls per hour allowed from a given IP address or a user token.
  @JsonKey(name: r'maxQueriesPerIPPerHour')
  final int? maxQueriesPerIPPerHour;

  /// URL-encoded query string. Force some query parameters to be applied for each query made with this API key.
  @JsonKey(name: r'queryParameters')
  final String? queryParameters;

  /// Restrict this new API key to specific referers. If empty or blank, defaults to all referers.
  @JsonKey(name: r'referers')
  final List<String>? referers;

  /// Validity limit for this key in seconds. The key will automatically be removed after this period of time.
  @JsonKey(name: r'validity')
  final int? validity;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is GetApiKeyResponse &&
          other.value == value &&
          other.createdAt == createdAt &&
          other.acl == acl &&
          other.description == description &&
          other.indexes == indexes &&
          other.maxHitsPerQuery == maxHitsPerQuery &&
          other.maxQueriesPerIPPerHour == maxQueriesPerIPPerHour &&
          other.queryParameters == queryParameters &&
          other.referers == referers &&
          other.validity == validity;

  @override
  int get hashCode =>
      value.hashCode +
      createdAt.hashCode +
      acl.hashCode +
      description.hashCode +
      indexes.hashCode +
      maxHitsPerQuery.hashCode +
      maxQueriesPerIPPerHour.hashCode +
      queryParameters.hashCode +
      referers.hashCode +
      validity.hashCode;

  factory GetApiKeyResponse.fromJson(Map<String, dynamic> json) =>
      _$GetApiKeyResponseFromJson(json);

  Map<String, dynamic> toJson() => _$GetApiKeyResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
