/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * GetNoResultsRateResponse
 *
 * @param rate The click-through rate.
 * @param count The number of occurrences.
 * @param noResultCount The number of occurrences.
 * @param dates A list of searches without results with their date, rate and counts.
 */
@Serializable
public data class GetNoResultsRateResponse(

  /** The click-through rate. */
  @SerialName(value = "rate") val rate: Double,

  /** The number of occurrences. */
  @SerialName(value = "count") val count: Int,

  /** The number of occurrences. */
  @SerialName(value = "noResultCount") val noResultCount: Int,

  /** A list of searches without results with their date, rate and counts. */
  @SerialName(value = "dates") val dates: List<NoResultsRateEvent>,
)
