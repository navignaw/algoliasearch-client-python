/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * Authentication input for OAuth login.
 *
 * @param url The OAuth endpoint URL.
 * @param clientId The clientID.
 * @param clientSecret The secret.
 */
@Serializable
public data class AuthOAuthPartial(

  /** The OAuth endpoint URL. */
  @SerialName(value = "url") val url: String? = null,

  /** The clientID. */
  @SerialName(value = "client_id") val clientId: String? = null,

  /** The secret. */
  @SerialName(value = "client_secret") val clientSecret: String? = null,
) : AuthInputPartial