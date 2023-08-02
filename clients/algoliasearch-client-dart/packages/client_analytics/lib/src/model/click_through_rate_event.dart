// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'click_through_rate_event.g.dart';

@JsonSerializable()
final class ClickThroughRateEvent {
  /// Returns a new [ClickThroughRateEvent] instance.
  const ClickThroughRateEvent({
    required this.rate,
    required this.clickCount,
    required this.trackedSearchCount,
    required this.date,
  });

  /// [Click-through rate (CTR)](https://www.algolia.com/doc/guides/search-analytics/concepts/metrics/#click-through-rate).
  // minimum: 0
  // maximum: 1
  @JsonKey(name: r'rate')
  final double rate;

  /// Number of click events.
  @JsonKey(name: r'clickCount')
  final int clickCount;

  /// Number of tracked searches. This is the number of search requests where the `clickAnalytics` parameter is `true`.
  @JsonKey(name: r'trackedSearchCount')
  final int trackedSearchCount;

  /// Date of the event in the format YYYY-MM-DD.
  @JsonKey(name: r'date')
  final String date;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is ClickThroughRateEvent &&
          other.rate == rate &&
          other.clickCount == clickCount &&
          other.trackedSearchCount == trackedSearchCount &&
          other.date == date;

  @override
  int get hashCode =>
      rate.hashCode +
      clickCount.hashCode +
      trackedSearchCount.hashCode +
      date.hashCode;

  factory ClickThroughRateEvent.fromJson(Map<String, dynamic> json) =>
      _$ClickThroughRateEventFromJson(json);

  Map<String, dynamic> toJson() => _$ClickThroughRateEventToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}