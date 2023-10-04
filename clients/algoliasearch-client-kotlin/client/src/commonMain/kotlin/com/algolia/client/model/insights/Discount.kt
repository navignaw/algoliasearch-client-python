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
 * Absolute value of the discount in effect for this object, measured in `currency`.
 */
@Serializable(DiscountSerializer::class)
public sealed interface Discount {

  public data class DoubleWrapper(val value: Double) : Discount

  public data class StringWrapper(val value: String) : Discount

  public companion object {

    /**
     * Discount as Double
     *
     */
    public fun Number(
      value: Double,
    ): DoubleWrapper = DoubleWrapper(
      value = value,
    )

    /**
     * Discount as String
     *
     */
    public fun String(
      value: String,
    ): StringWrapper = StringWrapper(
      value = value,
    )
  }
}

internal class DiscountSerializer : KSerializer<Discount> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("Discount")

  override fun serialize(encoder: Encoder, value: Discount) {
    when (value) {
      is Discount.DoubleWrapper -> Double.serializer().serialize(encoder, value.value)
      is Discount.StringWrapper -> String.serializer().serialize(encoder, value.value)
    }
  }

  override fun deserialize(decoder: Decoder): Discount {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize Double
    if (tree is JsonPrimitive) {
      try {
        return codec.json.decodeFromJsonElement<Discount.DoubleWrapper>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize Double (error: ${e.message})")
      }
    }

    // deserialize String
    if (tree is JsonPrimitive) {
      try {
        return codec.json.decodeFromJsonElement<Discount.StringWrapper>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize String (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
