// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'event_scoring.g.dart';

@JsonSerializable()
final class EventScoring {
  /// Returns a new [EventScoring] instance.
  const EventScoring({
    required this.score,
    required this.eventName,
    required this.eventType,
  });

  /// The score for the event.
  @JsonKey(name: r'score')
  final int score;

  /// The name of the event.
  @JsonKey(name: r'eventName')
  final String eventName;

  /// The type of the event.
  @JsonKey(name: r'eventType')
  final String eventType;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is EventScoring &&
          other.score == score &&
          other.eventName == eventName &&
          other.eventType == eventType;

  @override
  int get hashCode => score.hashCode + eventName.hashCode + eventType.hashCode;

  factory EventScoring.fromJson(Map<String, dynamic> json) =>
      _$EventScoringFromJson(json);

  Map<String, dynamic> toJson() => _$EventScoringToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}