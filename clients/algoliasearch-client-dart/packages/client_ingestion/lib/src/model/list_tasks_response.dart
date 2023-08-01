// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_ingestion/src/model/pagination.dart';
import 'package:algolia_client_ingestion/src/model/task.dart';

import 'package:json_annotation/json_annotation.dart';

part 'list_tasks_response.g.dart';

@JsonSerializable()
final class ListTasksResponse {
  /// Returns a new [ListTasksResponse] instance.
  const ListTasksResponse({
    required this.tasks,
    required this.pagination,
  });

  @JsonKey(name: r'tasks')
  final List<Task> tasks;

  @JsonKey(name: r'pagination')
  final Pagination pagination;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is ListTasksResponse &&
          other.tasks == tasks &&
          other.pagination == pagination;

  @override
  int get hashCode => tasks.hashCode + pagination.hashCode;

  factory ListTasksResponse.fromJson(Map<String, dynamic> json) =>
      _$ListTasksResponseFromJson(json);

  Map<String, dynamic> toJson() => _$ListTasksResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
