// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_analytics/src/model/search_event.dart';

import 'package:json_annotation/json_annotation.dart';

part 'get_searches_count_response.g.dart';

@JsonSerializable()
final class GetSearchesCountResponse {
  /// Returns a new [GetSearchesCountResponse] instance.
  const GetSearchesCountResponse({
    required this.count,
    required this.dates,
  });

  /// Number of occurrences.
  @JsonKey(name: r'count')
  final int count;

  /// Search events with their associated dates and hit counts.
  @JsonKey(name: r'dates')
  final List<SearchEvent> dates;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is GetSearchesCountResponse &&
          other.count == count &&
          other.dates == dates;

  @override
  int get hashCode => count.hashCode + dates.hashCode;

  factory GetSearchesCountResponse.fromJson(Map<String, dynamic> json) =>
      _$GetSearchesCountResponseFromJson(json);

  Map<String, dynamic> toJson() => _$GetSearchesCountResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}