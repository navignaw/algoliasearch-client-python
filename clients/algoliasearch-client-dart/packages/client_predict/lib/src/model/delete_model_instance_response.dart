// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'delete_model_instance_response.g.dart';

@JsonSerializable()
final class DeleteModelInstanceResponse {
  /// Returns a new [DeleteModelInstanceResponse] instance.
  const DeleteModelInstanceResponse({
    required this.modelID,
    required this.deletedUntil,
  });

  /// The ID of the model.
  @JsonKey(name: r'modelID')
  final String modelID;

  /// The date until which you can safely consider the data as being deleted.
  @JsonKey(name: r'deletedUntil')
  final String deletedUntil;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is DeleteModelInstanceResponse &&
          other.modelID == modelID &&
          other.deletedUntil == deletedUntil;

  @override
  int get hashCode => modelID.hashCode + deletedUntil.hashCode;

  factory DeleteModelInstanceResponse.fromJson(Map<String, dynamic> json) =>
      _$DeleteModelInstanceResponseFromJson(json);

  Map<String, dynamic> toJson() => _$DeleteModelInstanceResponseToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
