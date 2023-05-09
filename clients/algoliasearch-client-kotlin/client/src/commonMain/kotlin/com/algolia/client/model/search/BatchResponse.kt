/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * BatchResponse
 *
 * @param taskID taskID of the task to wait for.
 * @param objectIDs List of objectID.
 */
@Serializable
public data class BatchResponse(

  /** taskID of the task to wait for. */
  @SerialName(value = "taskID") val taskID: Long,

  /** List of objectID. */
  @SerialName(value = "objectIDs") val objectIDs: List<String>,
)
