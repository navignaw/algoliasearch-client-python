// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/synonym_hit.dart';

import 'package:collection/collection.dart';
import 'package:json_annotation/json_annotation.dart';

part 'search_synonyms_response.g.dart';

@JsonSerializable(createFieldMap: true)
final class SearchSynonymsResponse extends DelegatingMap<String, dynamic> {
  /// Returns a new [SearchSynonymsResponse] instance.
  const SearchSynonymsResponse({
    required this.hits,
    required this.nbHits,
    Map<String, dynamic> additionalProperties = const {},
  }) : super(additionalProperties);

  /// Array of synonym objects.
  @JsonKey(name: r'hits')
  final List<SynonymHit> hits;

  /// Number of hits that the search query matched.
  @JsonKey(name: r'nbHits')
  final int nbHits;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is SearchSynonymsResponse &&
          other.hits == hits &&
          other.nbHits == nbHits &&
          const MapEquality<String, dynamic>().equals(this, this);

  @override
  int get hashCode =>
      hits.hashCode +
      nbHits.hashCode +
      const MapEquality<String, dynamic>().hash(this);

  factory SearchSynonymsResponse.fromJson(Map<String, dynamic> json) {
    final instance = _$SearchSynonymsResponseFromJson(json);
    final additionalProperties = Map<String, dynamic>.from(json)
      ..removeWhere(
          (key, value) => _$SearchSynonymsResponseFieldMap.containsKey(key));
    return SearchSynonymsResponse(
      hits: instance.hits,
      nbHits: instance.nbHits,
      additionalProperties: additionalProperties,
    );
  }

  Map<String, dynamic> toJson() =>
      _$SearchSynonymsResponseToJson(this)..addAll(this);

  @override
  String toString() {
    return toJson().toString();
  }
}
