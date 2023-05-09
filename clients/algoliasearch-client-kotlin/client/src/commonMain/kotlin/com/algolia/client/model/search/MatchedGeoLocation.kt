/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * MatchedGeoLocation
 *
 * @param lat Latitude of the matched location.
 * @param lng Longitude of the matched location.
 * @param distance Distance between the matched location and the search location (in meters).
 */
@Serializable
public data class MatchedGeoLocation(

  /** Latitude of the matched location. */
  @SerialName(value = "lat") val lat: Double? = null,

  /** Longitude of the matched location. */
  @SerialName(value = "lng") val lng: Double? = null,

  /** Distance between the matched location and the search location (in meters). */
  @SerialName(value = "distance") val distance: Int? = null,
)
