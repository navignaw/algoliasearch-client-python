/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * SourceCreate
 *
 * @param type
 * @param name
 * @param input
 * @param authenticationID The authentication UUID.
 */
@Serializable
public data class SourceCreate(

  @SerialName(value = "type") val type: SourceType,

  @SerialName(value = "name") val name: String,

  @SerialName(value = "input") val input: SourceInput,

  /** The authentication UUID. */
  @SerialName(value = "authenticationID") val authenticationID: String? = null,
)
