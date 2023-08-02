// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element

import 'package:json_annotation/json_annotation.dart';

part 'auth_basic.g.dart';

@JsonSerializable()
final class AuthBasic {
  /// Returns a new [AuthBasic] instance.
  const AuthBasic({
    required this.username,
    required this.password,
  });

  @JsonKey(name: r'username')
  final String username;

  @JsonKey(name: r'password')
  final String password;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is AuthBasic &&
          other.username == username &&
          other.password == password;

  @override
  int get hashCode => username.hashCode + password.hashCode;

  factory AuthBasic.fromJson(Map<String, dynamic> json) =>
      _$AuthBasicFromJson(json);

  Map<String, dynamic> toJson() => _$AuthBasicToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}