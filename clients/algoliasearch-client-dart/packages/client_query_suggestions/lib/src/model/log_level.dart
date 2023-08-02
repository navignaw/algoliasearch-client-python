// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:json_annotation/json_annotation.dart';

/// The type of log entry.   - `SKIP`. A query is skipped because it doesn't match the conditions for successful inclusion. For example, when a query doesn't generate enough search results. - `INFO`. An informative log entry. - `ERROR`. The Query Suggestions process encountered an error.
@JsonEnum(valueField: 'raw')
enum LogLevel {
  /// The type of log entry.   - `SKIP`. A query is skipped because it doesn't match the conditions for successful inclusion. For example, when a query doesn't generate enough search results. - `INFO`. An informative log entry. - `ERROR`. The Query Suggestions process encountered an error.
  skip(r'SKIP'),

  /// The type of log entry.   - `SKIP`. A query is skipped because it doesn't match the conditions for successful inclusion. For example, when a query doesn't generate enough search results. - `INFO`. An informative log entry. - `ERROR`. The Query Suggestions process encountered an error.
  info(r'INFO'),

  /// The type of log entry.   - `SKIP`. A query is skipped because it doesn't match the conditions for successful inclusion. For example, when a query doesn't generate enough search results. - `INFO`. An informative log entry. - `ERROR`. The Query Suggestions process encountered an error.
  error(r'ERROR');

  const LogLevel(this.raw);
  final dynamic raw;

  dynamic toJson() => raw;

  static LogLevel fromJson(dynamic json) {
    for (final value in values) {
      if (value.raw == json) return value;
    }
    throw ArgumentError.value(json, "raw", "No enum value with that value");
  }

  @override
  String toString() => raw.toString();
}