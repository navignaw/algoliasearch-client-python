/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.insights

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * The price of the item. This should be the final price, inclusive of any discounts in effect.
 */
@Serializable(PriceSerializer::class)
public sealed interface Price {

  public data class DoubleWrapper(val value: Double) : Price

  public data class StringWrapper(val value: String) : Price

  public companion object {

    /**
     * Price as Double
     *
     */
    public fun Number(
      value: Double,
    ): DoubleWrapper = DoubleWrapper(
      value = value,
    )

    /**
     * Price as String
     *
     */
    public fun String(
      value: String,
    ): StringWrapper = StringWrapper(
      value = value,
    )
  }
}

internal class PriceSerializer : KSerializer<Price> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("Price")

  override fun serialize(encoder: Encoder, value: Price) {
    when (value) {
      is Price.DoubleWrapper -> Double.serializer().serialize(encoder, value.value)
      is Price.StringWrapper -> String.serializer().serialize(encoder, value.value)
    }
  }

  override fun deserialize(decoder: Decoder): Price {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize Double
    if (tree is JsonPrimitive) {
      try {
        return codec.json.decodeFromJsonElement<Price.DoubleWrapper>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize Double (error: ${e.message})")
      }
    }

    // deserialize String
    if (tree is JsonPrimitive) {
      try {
        return codec.json.decodeFromJsonElement<Price.StringWrapper>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize String (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}