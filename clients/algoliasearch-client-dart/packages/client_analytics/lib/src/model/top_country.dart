// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'top_country.g.dart';

@JsonSerializable()
final class TopCountry {
  /// Returns a new [TopCountry] instance.
  const TopCountry({
    required this.country,
    required this.count,
  });

  /// Country.
  @JsonKey(name: r'country')
  final String country;

  /// Number of occurrences.
  @JsonKey(name: r'count')
  final int count;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is TopCountry && other.country == country && other.count == count;

  @override
  int get hashCode => country.hashCode + count.hashCode;

  factory TopCountry.fromJson(Map<String, dynamic> json) =>
      _$TopCountryFromJson(json);

  Map<String, dynamic> toJson() => _$TopCountryToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}