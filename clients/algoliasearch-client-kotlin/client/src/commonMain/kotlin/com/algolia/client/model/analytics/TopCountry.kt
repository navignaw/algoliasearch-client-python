/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * TopCountry
 *
 * @param country The country.
 * @param count The number of occurrences.
 */
@Serializable
public data class TopCountry(

  /** The country. */
  @SerialName(value = "country") val country: String,

  /** The number of occurrences. */
  @SerialName(value = "count") val count: Int,
)
