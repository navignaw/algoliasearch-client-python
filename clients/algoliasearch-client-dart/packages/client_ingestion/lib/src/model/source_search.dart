// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'source_search.g.dart';

@JsonSerializable()
final class SourceSearch {
  /// Returns a new [SourceSearch] instance.
  const SourceSearch({
    required this.sourceIDs,
  });

  @JsonKey(name: r'sourceIDs')
  final List<String> sourceIDs;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is SourceSearch && other.sourceIDs == sourceIDs;

  @override
  int get hashCode => sourceIDs.hashCode;

  factory SourceSearch.fromJson(Map<String, dynamic> json) =>
      _$SourceSearchFromJson(json);

  Map<String, dynamic> toJson() => _$SourceSearchToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}