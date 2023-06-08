// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
// ignore_for_file: unused_element
import 'package:algoliasearch/src/model/matched_geo_location.dart';
import 'package:algoliasearch/src/model/personalization.dart';

import 'package:json_annotation/json_annotation.dart';

part 'ranking_info.g.dart';

@JsonSerializable()
final class RankingInfo {
  /// Returns a new [RankingInfo] instance.
  const RankingInfo({
    required this.filters,
    required this.firstMatchedWord,
    required this.geoDistance,
    this.geoPrecision,
    this.matchedGeoLocation,
    this.personalization,
    required this.nbExactWords,
    required this.nbTypos,
    required this.promoted,
    this.proximityDistance,
    required this.userScore,
    required this.words,
    this.promotedByReRanking,
  });

  /// This field is reserved for advanced usage.
  @JsonKey(name: r'filters')
  final int filters;

  /// Position of the most important matched attribute in the attributes to index list.
  @JsonKey(name: r'firstMatchedWord')
  final int firstMatchedWord;

  /// Distance between the geo location in the search query and the best matching geo location in the record, divided by the geo precision (in meters).
  @JsonKey(name: r'geoDistance')
  final int geoDistance;

  /// Precision used when computing the geo distance, in meters.
  @JsonKey(name: r'geoPrecision')
  final int? geoPrecision;

  @JsonKey(name: r'matchedGeoLocation')
  final MatchedGeoLocation? matchedGeoLocation;

  @JsonKey(name: r'personalization')
  final Personalization? personalization;

  /// Number of exactly matched words.
  @JsonKey(name: r'nbExactWords')
  final int nbExactWords;

  /// Number of typos encountered when matching the record.
  @JsonKey(name: r'nbTypos')
  final int nbTypos;

  /// Present and set to true if a Rule promoted the hit.
  @JsonKey(name: r'promoted')
  final bool promoted;

  /// When the query contains more than one word, the sum of the distances between matched words (in meters).
  @JsonKey(name: r'proximityDistance')
  final int? proximityDistance;

  /// Custom ranking for the object, expressed as a single integer value.
  @JsonKey(name: r'userScore')
  final int userScore;

  /// Number of matched words, including prefixes and typos.
  @JsonKey(name: r'words')
  final int words;

  /// Wether the record are promoted by the re-ranking strategy.
  @JsonKey(name: r'promotedByReRanking')
  final bool? promotedByReRanking;

  @override
  bool operator ==(Object other) =>
      identical(this, other) ||
      other is RankingInfo &&
          other.filters == filters &&
          other.firstMatchedWord == firstMatchedWord &&
          other.geoDistance == geoDistance &&
          other.geoPrecision == geoPrecision &&
          other.matchedGeoLocation == matchedGeoLocation &&
          other.personalization == personalization &&
          other.nbExactWords == nbExactWords &&
          other.nbTypos == nbTypos &&
          other.promoted == promoted &&
          other.proximityDistance == proximityDistance &&
          other.userScore == userScore &&
          other.words == words &&
          other.promotedByReRanking == promotedByReRanking;

  @override
  int get hashCode =>
      filters.hashCode +
      firstMatchedWord.hashCode +
      geoDistance.hashCode +
      geoPrecision.hashCode +
      matchedGeoLocation.hashCode +
      personalization.hashCode +
      nbExactWords.hashCode +
      nbTypos.hashCode +
      promoted.hashCode +
      proximityDistance.hashCode +
      userScore.hashCode +
      words.hashCode +
      promotedByReRanking.hashCode;

  factory RankingInfo.fromJson(Map<String, dynamic> json) =>
      _$RankingInfoFromJson(json);

  Map<String, dynamic> toJson() => _$RankingInfoToJson(this);

  @override
  String toString() {
    return toJson().toString();
  }
}