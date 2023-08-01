// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_ingestion/src/model/pagination.dart';
import 'package:algolia_client_ingestion/src/model/run.dart';

import 'package:json_annotation/json_annotation.dart';

part 'run_list_response.g.dart';

@JsonSerializable()
final class RunListResponse {
  /// Returns a new [RunListResponse] instance.
  const RunListResponse({
    required this.runs,
    required this.pagination,
  });

  @JsonKey(name: r'runs')
  final List<Run> runs;

  @JsonKey(name: r'pagination')
  final Pagination pagination;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is RunListResponse &&
          other.runs == runs &&
          other.pagination == pagination;

  @override
  int get hashCode => runs.hashCode + pagination.hashCode;

  factory RunListResponse.fromJson(Map<String, dynamic> json) =>
      _$RunListResponseFromJson(json);

  Map<String, dynamic> toJson() => _$RunListResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
