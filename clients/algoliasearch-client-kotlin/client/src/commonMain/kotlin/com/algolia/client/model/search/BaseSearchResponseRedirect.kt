/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * This parameter is for internal use only.
 *
 * @param index
 */
@Serializable
public data class BaseSearchResponseRedirect(

  @SerialName(value = "index") val index: List<RedirectRuleIndexMetadata>? = null,
)
