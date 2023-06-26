/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.insights

import kotlinx.serialization.*
import kotlinx.serialization.json.*

/**
 * Click event after an Algolia request.  Use this event to track when users click items in the search results. If you're building your category pages with Algolia, you'll also use this event.
 *
 * @param eventName Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.
 * @param eventType
 * @param index Name of the Algolia index.
 * @param objectIDs List of object identifiers for items of an Algolia index.
 * @param positions Position of the clicked objects in the search results.  The first search result has a position of 1 (not 0). You must provide 1 `position` for each `objectID`.
 * @param queryID Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.
 * @param userToken Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.
 * @param timestamp Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.
 */
@Serializable
public data class ClickedObjectIDsAfterSearch(

  /** Can contain up to 64 ASCII characters.   Consider naming events consistently—for example, by adopting Segment's [object-action](https://segment.com/academy/collecting-data/naming-conventions-for-clean-data/#the-object-action-framework) framework.  */
  @SerialName(value = "eventName") val eventName: String,

  @SerialName(value = "eventType") val eventType: ClickEvent,

  /** Name of the Algolia index. */
  @SerialName(value = "index") val index: String,

  /** List of object identifiers for items of an Algolia index. */
  @SerialName(value = "objectIDs") val objectIDs: List<String>,

  /** Position of the clicked objects in the search results.  The first search result has a position of 1 (not 0). You must provide 1 `position` for each `objectID`.  */
  @SerialName(value = "positions") val positions: List<Int>,

  /** Unique identifier for a search query.  The query ID is required for events related to search or browse requests. If you add `clickAnalytics: true` as a search request parameter, the query ID is included in the API response.  */
  @SerialName(value = "queryID") val queryID: String,

  /** Anonymous or pseudonymous user identifier.   > **Note**: Never include personally identifiable information in user tokens.  */
  @SerialName(value = "userToken") val userToken: String,

  /** Time of the event in milliseconds in [Unix epoch time](https://wikipedia.org/wiki/Unix_time). By default, the Insights API uses the time it receives an event as its timestamp.  */
  @SerialName(value = "timestamp") val timestamp: Long? = null,
) : EventsItems