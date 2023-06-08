// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'delete_source_response.g.dart';

@JsonSerializable()
final class DeleteSourceResponse {
  /// Returns a new [DeleteSourceResponse] instance.
  const DeleteSourceResponse({
    required this.deletedAt,
  });

  /// Date of deletion (ISO-8601 format).
  @JsonKey(name: r'deletedAt')
  final String deletedAt;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is DeleteSourceResponse && other.deletedAt == deletedAt;

  @override
  int get hashCode => deletedAt.hashCode;

  factory DeleteSourceResponse.fromJson(Map<String, dynamic> json) =>
      _$DeleteSourceResponseFromJson(json);

  Map<String, dynamic> toJson() => _$DeleteSourceResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}