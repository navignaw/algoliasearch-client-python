/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import kotlinx.serialization.*

/**
 * The operation to apply on the attribute.
 */
@Serializable
public enum class BuiltInOperationType(public val value: kotlin.String) {

  @SerialName(value = "Increment")
  Increment("Increment"),

  @SerialName(value = "Decrement")
  Decrement("Decrement"),

  @SerialName(value = "Add")
  Add("Add"),

  @SerialName(value = "Remove")
  Remove("Remove"),

  @SerialName(value = "AddUnique")
  AddUnique("AddUnique"),

  @SerialName(value = "IncrementFrom")
  IncrementFrom("IncrementFrom"),

  @SerialName(value = "IncrementSet")
  IncrementSet("IncrementSet");

  override fun toString(): kotlin.String = value
}
