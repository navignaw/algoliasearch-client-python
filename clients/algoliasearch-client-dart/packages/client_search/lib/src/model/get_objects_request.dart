// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'get_objects_request.g.dart';

@JsonSerializable()
final class GetObjectsRequest {
  /// Returns a new [GetObjectsRequest] instance.
  const GetObjectsRequest({
    this.attributesToRetrieve,
    required this.objectID,
    required this.indexName,
  });

  /// List of attributes to retrieve. By default, all retrievable attributes are returned.
  @JsonKey(name: r'attributesToRetrieve')
  final List<String>? attributesToRetrieve;

  /// ID of the object within that index.
  @JsonKey(name: r'objectID')
  final String objectID;

  /// name of the index containing the object.
  @JsonKey(name: r'indexName')
  final String indexName;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is GetObjectsRequest &&
          other.attributesToRetrieve == attributesToRetrieve &&
          other.objectID == objectID &&
          other.indexName == indexName;

  @override
  int get hashCode =>
      attributesToRetrieve.hashCode + objectID.hashCode + indexName.hashCode;

  factory GetObjectsRequest.fromJson(Map<String, dynamic> json) =>
      _$GetObjectsRequestFromJson(json);

  Map<String, dynamic> toJson() => _$GetObjectsRequestToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
