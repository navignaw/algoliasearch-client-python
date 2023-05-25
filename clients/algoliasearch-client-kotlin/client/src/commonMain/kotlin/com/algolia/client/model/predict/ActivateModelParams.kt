/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.predict

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * ActivateModelParams
 *
 * @param type
 * @param name The model’s instance name.
 * @param sourceID The data source ID, as returned by the (external) sources API.
 * @param index The index name.
 * @param modelAttributes
 */
@Serializable
public data class ActivateModelParams(

  @SerialName(value = "type") val type: ModelsToRetrieve,

  /** The model’s instance name. */
  @SerialName(value = "name") val name: String,

  /** The data source ID, as returned by the (external) sources API. */
  @SerialName(value = "sourceID") val sourceID: String,

  /** The index name. */
  @SerialName(value = "index") val index: String,

  @SerialName(value = "modelAttributes") val modelAttributes: List<String>? = null,
)