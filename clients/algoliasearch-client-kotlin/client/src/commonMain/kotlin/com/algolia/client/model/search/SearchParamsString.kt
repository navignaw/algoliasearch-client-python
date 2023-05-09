/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * SearchParamsString
 *
 * @param params Search parameters as URL-encoded query string.
 */
@Serializable
public data class SearchParamsString(

  /** Search parameters as URL-encoded query string. */
  @SerialName(value = "params") val params: String? = null,
) : SearchParams, BrowseParams
