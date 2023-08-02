// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algolia_client_ingestion/src/model/on_demand_trigger_type.dart';

import 'package:json_annotation/json_annotation.dart';

part 'on_demand_trigger.g.dart';

@JsonSerializable()
final class OnDemandTrigger {
  /// Returns a new [OnDemandTrigger] instance.
  const OnDemandTrigger({
    required this.type,
    this.lastRun,
  });

  @JsonKey(name: r'type')
  final OnDemandTriggerType type;

  /// The last time the scheduled task ran (RFC3339 format).
  @JsonKey(name: r'lastRun')
  final String? lastRun;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is OnDemandTrigger &&
          other.type == type &&
          other.lastRun == lastRun;

  @override
  int get hashCode => type.hashCode + lastRun.hashCode;

  factory OnDemandTrigger.fromJson(Map<String, dynamic> json) =>
      _$OnDemandTriggerFromJson(json);

  Map<String, dynamic> toJson() => _$OnDemandTriggerToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}