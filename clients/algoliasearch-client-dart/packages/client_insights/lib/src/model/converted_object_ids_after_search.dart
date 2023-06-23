// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_insights/src/model/conversion_event.dart';

import 'package:json_annotation/json_annotation.dart';

part 'converted_object_ids_after_search.g.dart';

@JsonSerializable()
final class ConvertedObjectIDsAfterSearch {
  /// Returns a new [ConvertedObjectIDsAfterSearch] instance.
  const ConvertedObjectIDsAfterSearch({
    required this.eventName,
    required this.eventType,
    required this.index,
    required this.objectIDs,
    required this.queryID,
    required this.userToken,
    this.timestamp,
  });

  /// Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
  @JsonKey(name: r'eventName')
  final String eventName;

  @JsonKey(name: r'eventType')
  final ConversionEvent eventType;

  /// Name of the Algolia index.
  @JsonKey(name: r'index')
  final String index;

  /// List of object identifiers for items of an Algolia index.
  @JsonKey(name: r'objectIDs')
  final List<String> objectIDs;

  /// Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
  @JsonKey(name: r'queryID')
  final String queryID;

  /// Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
  @JsonKey(name: r'userToken')
  final String userToken;

  /// Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
  @JsonKey(name: r'timestamp')
  final int? timestamp;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is ConvertedObjectIDsAfterSearch &&
          other.eventName == eventName &&
          other.eventType == eventType &&
          other.index == index &&
          other.objectIDs == objectIDs &&
          other.queryID == queryID &&
          other.userToken == userToken &&
          other.timestamp == timestamp;

  @override
  int get hashCode =>
      eventName.hashCode +
      eventType.hashCode +
      index.hashCode +
      objectIDs.hashCode +
      queryID.hashCode +
      userToken.hashCode +
      timestamp.hashCode;

  factory ConvertedObjectIDsAfterSearch.fromJson(Map<String, dynamic> json) =>
      _$ConvertedObjectIDsAfterSearchFromJson(json);

  Map<String, dynamic> toJson() => _$ConvertedObjectIDsAfterSearchToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
