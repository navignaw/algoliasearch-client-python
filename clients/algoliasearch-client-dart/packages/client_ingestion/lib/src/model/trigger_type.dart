// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:json_annotation/json_annotation.dart';

/// The type of the task reflect how it can be used:   - onDemand: a task that runs manually   - schedule: a task that runs regularly, following a given cron expression   - subscription: a task that runs after a subscription event is received from an integration (e.g. Webhook).
@JsonEnum(valueField: 'raw')
enum TriggerType {
  /// The type of the task reflect how it can be used:   - onDemand: a task that runs manually   - schedule: a task that runs regularly, following a given cron expression   - subscription: a task that runs after a subscription event is received from an integration (e.g. Webhook).
  onDemand(r'onDemand'),

  /// The type of the task reflect how it can be used:   - onDemand: a task that runs manually   - schedule: a task that runs regularly, following a given cron expression   - subscription: a task that runs after a subscription event is received from an integration (e.g. Webhook).
  schedule(r'schedule'),

  /// The type of the task reflect how it can be used:   - onDemand: a task that runs manually   - schedule: a task that runs regularly, following a given cron expression   - subscription: a task that runs after a subscription event is received from an integration (e.g. Webhook).
  subscription(r'subscription');

  const TriggerType(this.raw);
  final dynamic raw;

  dynamic toJson() => raw;

  static TriggerType fromJson(dynamic json) {
    for (final value in values) {
      if (value.raw == json) return value;
    }
    throw ArgumentError.value(json, "raw", "No enum value with that value");
  }

  @override
  String toString() => raw.toString();
}