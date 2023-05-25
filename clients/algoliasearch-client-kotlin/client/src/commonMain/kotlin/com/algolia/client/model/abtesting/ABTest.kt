/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.abtesting

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * ABTest
 *
 * @param abTestID The A/B test ID.
 * @param clickSignificance A/B test significance based on click data. Should be > 0.95 to be considered significant (no matter which variant is winning).
 * @param conversionSignificance A/B test significance based on conversion data. Should be > 0.95 to be considered significant (no matter which variant is winning).
 * @param endAt End date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ.
 * @param updatedAt Update date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ.
 * @param createdAt Creation date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ.
 * @param name A/B test name.
 * @param status status of the A/B test.
 * @param variants List of A/B test variant.
 */
@Serializable
public data class ABTest(

  /** The A/B test ID. */
  @SerialName(value = "abTestID") val abTestID: Int,

  /** A/B test significance based on click data. Should be > 0.95 to be considered significant (no matter which variant is winning). */
  @SerialName(value = "clickSignificance") val clickSignificance: Double,

  /** A/B test significance based on conversion data. Should be > 0.95 to be considered significant (no matter which variant is winning). */
  @SerialName(value = "conversionSignificance") val conversionSignificance: Double,

  /** End date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ. */
  @SerialName(value = "endAt") val endAt: String,

  /** Update date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ. */
  @SerialName(value = "updatedAt") val updatedAt: String,

  /** Creation date for the A/B test expressed as YYYY-MM-DDThh:mm:ssZ. */
  @SerialName(value = "createdAt") val createdAt: String,

  /** A/B test name. */
  @SerialName(value = "name") val name: String,

  /** status of the A/B test. */
  @SerialName(value = "status") val status: String,

  /** List of A/B test variant. */
  @SerialName(value = "variants") val variants: List<Variant>,
)