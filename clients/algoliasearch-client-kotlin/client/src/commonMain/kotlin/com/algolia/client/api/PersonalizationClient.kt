/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.api

import com.algolia.client.configuration.*
import com.algolia.client.exception.*
import com.algolia.client.extensions.internal.*
import com.algolia.client.model.personalization.*
import com.algolia.client.transport.*
import com.algolia.client.transport.internal.*
import kotlinx.serialization.json.*

public class PersonalizationClient(
  override val appId: String,
  override val apiKey: String,
  public val region: String,
  override val options: ClientOptions = ClientOptions(),
) : ApiClient {

  init {
    require(appId.isNotBlank()) { "`appId` is missing." }
    require(apiKey.isNotBlank()) { "`apiKey` is missing." }
  }

  override val requester: Requester = requesterOf(clientName = "Personalization", appId = appId, apiKey = apiKey, options = options) {
    val allowedRegions = listOf("eu", "us")
    require(region in allowedRegions) { "`region` is required and must be one of the following: ${allowedRegions.joinToString()}" }
    val url = "personalization.$region.algolia.com"
    listOf(Host(url))
  }

  /**
   * Send requests to the Algolia REST API.
   * This method allow you to send requests to the Algolia REST API.
   * @param path The path of the API endpoint to target, anything after the /1 needs to be specified.
   * @param parameters Query parameters to be applied to the current query.
   * @param requestOptions additional request configuration.
   */
  public suspend fun del(path: String, parameters: Map<kotlin.String, Any>? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `del`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = "/1{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Delete a user profile.
   * Delete the user profile and all its associated data.  Returns, as part of the response, a date until which the data can safely be considered as deleted for the given user. This means if you send events for the given user before this date, they will be ignored. Any data received after the deletedUntil date will start building a new user profile.  It might take a couple hours for the deletion request to be fully processed.
   * @param userToken userToken representing the user for which to fetch the Personalization profile.
   * @param requestOptions additional request configuration.
   */
  public suspend fun deleteUserProfile(userToken: String, requestOptions: RequestOptions? = null): DeleteUserProfileResponse {
    require(userToken.isNotBlank()) { "Parameter `userToken` is required when calling `deleteUserProfile`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.DELETE,
      path = listOf("1", "profiles", "$userToken"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Send requests to the Algolia REST API.
   * This method allow you to send requests to the Algolia REST API.
   * @param path The path of the API endpoint to target, anything after the /1 needs to be specified.
   * @param parameters Query parameters to be applied to the current query.
   * @param requestOptions additional request configuration.
   */
  public suspend fun get(path: String, parameters: Map<kotlin.String, Any>? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `get`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = "/1{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get the current strategy.
   * The strategy contains information on the events and facets that impact user profiles and personalized search results.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getPersonalizationStrategy(requestOptions: RequestOptions? = null): PersonalizationStrategyParams {
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "strategies", "personalization"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Get a user profile.
   * Get the user profile built from Personalization strategy.  The profile is structured by facet name used in the strategy. Each facet value is mapped to its score. Each score represents the user affinity for a specific facet value given the userToken past events and the Personalization strategy defined. Scores are bounded to 20. The last processed event timestamp is provided using the ISO 8601 format for debugging purposes.
   * @param userToken userToken representing the user for which to fetch the Personalization profile.
   * @param requestOptions additional request configuration.
   */
  public suspend fun getUserTokenProfile(userToken: String, requestOptions: RequestOptions? = null): GetUserTokenResponse {
    require(userToken.isNotBlank()) { "Parameter `userToken` is required when calling `getUserTokenProfile`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.GET,
      path = listOf("1", "profiles", "personalization", "$userToken"),
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Send requests to the Algolia REST API.
   * This method allow you to send requests to the Algolia REST API.
   * @param path The path of the API endpoint to target, anything after the /1 needs to be specified.
   * @param parameters Query parameters to be applied to the current query.
   * @param body The parameters to send with the custom request.
   * @param requestOptions additional request configuration.
   */
  public suspend fun post(path: String, parameters: Map<kotlin.String, Any>? = null, body: JsonObject? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `post`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = "/1{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
      body = body,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Send requests to the Algolia REST API.
   * This method allow you to send requests to the Algolia REST API.
   * @param path The path of the API endpoint to target, anything after the /1 needs to be specified.
   * @param parameters Query parameters to be applied to the current query.
   * @param body The parameters to send with the custom request.
   * @param requestOptions additional request configuration.
   */
  public suspend fun put(path: String, parameters: Map<kotlin.String, Any>? = null, body: JsonObject? = null, requestOptions: RequestOptions? = null): JsonObject {
    require(path.isNotBlank()) { "Parameter `path` is required when calling `put`." }
    val requestConfig = RequestConfig(
      method = RequestMethod.PUT,
      path = "/1{path}".replace("{path}", path),
      query = buildMap {
        parameters?.let { putAll(it) }
      },
      body = body,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }

  /**
   * Set a new strategy.
   * A strategy defines the events and facets that impact user profiles and personalized search results.
   * @param personalizationStrategyParams
   * @param requestOptions additional request configuration.
   */
  public suspend fun setPersonalizationStrategy(personalizationStrategyParams: PersonalizationStrategyParams, requestOptions: RequestOptions? = null): SetPersonalizationStrategyResponse {
    val requestConfig = RequestConfig(
      method = RequestMethod.POST,
      path = listOf("1", "strategies", "personalization"),
      body = personalizationStrategyParams,
    )
    return requester.execute(
      requestConfig = requestConfig,
      requestOptions = requestOptions,
    )
  }
}