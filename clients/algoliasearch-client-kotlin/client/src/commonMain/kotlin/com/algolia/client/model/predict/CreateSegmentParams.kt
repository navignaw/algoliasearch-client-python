/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.predict

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * CreateSegmentParams
 *
 * @param name The name or description of the segment.
 * @param conditions
 */
@Serializable
public data class CreateSegmentParams(

  /** The name or description of the segment. */
  @SerialName(value = "name") val name: String,

  @SerialName(value = "conditions") val conditions: SegmentParentConditions,
)