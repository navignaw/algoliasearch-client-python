/** Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT. */
package com.algolia.client.model.ingestion

import com.algolia.client.exception.AlgoliaClientException
import com.algolia.client.extensions.internal.*
import kotlinx.serialization.*
import kotlinx.serialization.builtins.*
import kotlinx.serialization.descriptors.*
import kotlinx.serialization.encoding.*
import kotlinx.serialization.json.*

/**
 * AuthInput
 */
@Serializable(AuthInputSerializer::class)
public sealed interface AuthInput {

  public companion object {

    /**
     * Authentication input used for token credentials.
     *
     * @param key
     */
    public fun AuthAPIKey(
      key: String,
    ): AuthAPIKey = com.algolia.client.model.ingestion.AuthAPIKey(
      key = key,
    )

    /**
     * AuthAlgolia
     *
     * @param appID Algolia Application ID.
     * @param apiKey Algolia API Key, with the correct rights to push to an index and change settings.
     */
    public fun AuthAlgolia(
      appID: String,
      apiKey: String,
    ): AuthAlgolia = com.algolia.client.model.ingestion.AuthAlgolia(
      appID = appID,
      apiKey = apiKey,
    )

    /**
     * Authentication input for Basic login with username and password.
     *
     * @param username
     * @param password
     */
    public fun AuthBasic(
      username: String,
      password: String,
    ): AuthBasic = com.algolia.client.model.ingestion.AuthBasic(
      username = username,
      password = password,
    )

    /**
     * Authentication input to connect to a Google service (e.g. BigQuery).
     *
     * @param clientEmail Email address of the Service Account.
     * @param privateKey Private key of the Service Account.
     */
    public fun AuthGoogleServiceAccount(
      clientEmail: String,
      privateKey: String,
    ): AuthGoogleServiceAccount = com.algolia.client.model.ingestion.AuthGoogleServiceAccount(
      clientEmail = clientEmail,
      privateKey = privateKey,
    )

    /**
     * Authentication input for OAuth login.
     *
     * @param url The OAuth endpoint URL.
     * @param clientId The clientID.
     * @param clientSecret The secret.
     */
    public fun AuthOAuth(
      url: String,
      clientId: String,
      clientSecret: String,
    ): AuthOAuth = com.algolia.client.model.ingestion.AuthOAuth(
      url = url,
      clientId = clientId,
      clientSecret = clientSecret,
    )
  }
}

internal class AuthInputSerializer : KSerializer<AuthInput> {

  override val descriptor: SerialDescriptor = buildClassSerialDescriptor("AuthInput")

  override fun serialize(encoder: Encoder, value: AuthInput) {
    when (value) {
      is AuthAPIKey -> AuthAPIKey.serializer().serialize(encoder, value)
      is AuthAlgolia -> AuthAlgolia.serializer().serialize(encoder, value)
      is AuthBasic -> AuthBasic.serializer().serialize(encoder, value)
      is AuthGoogleServiceAccount -> AuthGoogleServiceAccount.serializer().serialize(encoder, value)
      is AuthOAuth -> AuthOAuth.serializer().serialize(encoder, value)
    }
  }

  override fun deserialize(decoder: Decoder): AuthInput {
    val codec = decoder.asJsonDecoder()
    val tree = codec.decodeJsonElement()

    // deserialize AuthAPIKey
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AuthAPIKey>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AuthAPIKey (error: ${e.message})")
      }
    }

    // deserialize AuthAlgolia
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AuthAlgolia>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AuthAlgolia (error: ${e.message})")
      }
    }

    // deserialize AuthBasic
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AuthBasic>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AuthBasic (error: ${e.message})")
      }
    }

    // deserialize AuthGoogleServiceAccount
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AuthGoogleServiceAccount>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AuthGoogleServiceAccount (error: ${e.message})")
      }
    }

    // deserialize AuthOAuth
    if (tree is JsonObject) {
      try {
        return codec.json.decodeFromJsonElement<AuthOAuth>(tree)
      } catch (e: Exception) {
        // deserialization failed, continue
        println("Failed to deserialize AuthOAuth (error: ${e.message})")
      }
    }

    throw AlgoliaClientException("Failed to deserialize json element: $tree")
  }
}
