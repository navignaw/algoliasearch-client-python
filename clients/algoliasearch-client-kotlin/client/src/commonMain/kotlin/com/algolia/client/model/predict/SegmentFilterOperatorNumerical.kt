/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.predict

import kotlinx.serialization.*

/**
 * The operator used on the numerical filter value.
 */
@Serializable
public enum class SegmentFilterOperatorNumerical(public val value: kotlin.String) {

  @SerialName(value = "EQ")
  EQ("EQ"),

  @SerialName(value = "NEQ")
  NEQ("NEQ"),

  @SerialName(value = "GT")
  GT("GT"),

  @SerialName(value = "GTE")
  GTE("GTE"),

  @SerialName(value = "LT")
  LT("LT"),

  @SerialName(value = "LTE")
  LTE("LTE");

  override fun toString(): kotlin.String = value
}
