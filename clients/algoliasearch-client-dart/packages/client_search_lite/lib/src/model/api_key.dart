// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_search_lite/src/model/acl.dart';

import 'package:json_annotation/json_annotation.dart';

part 'api_key.g.dart';

@JsonSerializable()
final class ApiKey {
  /// Returns a new [ApiKey] instance.
  const ApiKey({
    required this.acl,
    this.description,
    this.indexes,
    this.maxHitsPerQuery,
    this.maxQueriesPerIPPerHour,
    this.queryParameters,
    this.referers,
    this.validity,
  });

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
      other is ApiKey &&
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
      acl.hashCode +
      description.hashCode +
      indexes.hashCode +
      maxHitsPerQuery.hashCode +
      maxQueriesPerIPPerHour.hashCode +
      queryParameters.hashCode +
      referers.hashCode +
      validity.hashCode;

  factory ApiKey.fromJson(Map<String, dynamic> json) => _$ApiKeyFromJson(json);

  Map<String, dynamic> toJson() => _$ApiKeyToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
