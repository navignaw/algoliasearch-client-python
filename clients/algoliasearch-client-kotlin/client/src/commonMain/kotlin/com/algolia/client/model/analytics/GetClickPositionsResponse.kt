/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.analytics

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * GetClickPositionsResponse
 *
 * @param positions A list of the click positions with their click count.
 */
@Serializable
public data class GetClickPositionsResponse(

  /** A list of the click positions with their click count. */
  @SerialName(value = "positions") val positions: List<ClickPosition>,
)