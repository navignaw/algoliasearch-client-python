/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.search

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * AttributeToUpdate
 */
@Serializable(AttributeToUpdateSerializer::class)
public sealed interface AttributeToUpdate {

  public data class StringWrapper(val value: String) : AttributeToUpdate

  public companion object {

    /**
     * To update an attribute without pushing the entire record, you can use these built-in operations.
     *
     * @param operation
     * @param `value` the right-hand side argument to the operation, for example, increment or decrement step, value to add or remove.
     */
    public fun BuiltInOperation(
      operation: BuiltInOperationType,
      `value`: String,
    ): BuiltInOperation = com.algolia.client.model.search.BuiltInOperation(
      operation = operation,
      `value` = `value`,
    )

    /**
     * AttributeToUpdate as String
     *
     */
    public fun String(
      value: String,
    ): StringWrapper = StringWrapper(
      value = value,
    )
  }
}

internal class AttributeToUpdateSerializer : KSerializer<AttributeToUpdate> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("AttributeToUpdate")

  override fun serialize(encoder: Encoder, value: AttributeToUpdate) {
    when (value) {
      is BuiltInOperation -> BuiltInOperation.serializer().serialize(encoder, value)
      is AttributeToUpdate.StringWrapper -> String.serializer().serialize(encoder, value.value)
    }
  }

  override fun deserialize(decoder: Decoder): AttributeToUpdate {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize BuiltInOperation
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<BuiltInOperation>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize BuiltInOperation (error: ${e.message})")
      }
    }

    // deserialize String
    if (tree is JsonPrimitive) {
      try {
        return codec.json.decodeFromJsonElement<AttributeToUpdate.StringWrapper>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize String (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}