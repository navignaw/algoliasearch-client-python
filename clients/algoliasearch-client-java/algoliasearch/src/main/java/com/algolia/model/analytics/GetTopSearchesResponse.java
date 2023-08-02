// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.analytics;

import com.algolia.exceptions.AlgoliaRuntimeException;
import com.algolia.utils.CompoundType;
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.core.*;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.fasterxml.jackson.databind.deser.std.StdDeserializer;
import com.fasterxml.jackson.databind.ser.std.StdSerializer;
import java.io.IOException;
import java.util.logging.Logger;

/** GetTopSearchesResponse */
@JsonDeserialize(using = GetTopSearchesResponse.GetTopSearchesResponseDeserializer.class)
@JsonSerialize(using = GetTopSearchesResponse.GetTopSearchesResponseSerializer.class)
public abstract class GetTopSearchesResponse implements CompoundType {

  private static final Logger LOGGER = Logger.getLogger(GetTopSearchesResponse.class.getName());

  public static GetTopSearchesResponse of(TopSearchesResponse inside) {
    return new GetTopSearchesResponseTopSearchesResponse(inside);
  }

  public static GetTopSearchesResponse of(TopSearchesResponseWithAnalytics inside) {
    return new GetTopSearchesResponseTopSearchesResponseWithAnalytics(inside);
  }

  public static class GetTopSearchesResponseSerializer extends StdSerializer<GetTopSearchesResponse> {

    public GetTopSearchesResponseSerializer(Class<GetTopSearchesResponse> t) {
      super(t);
    }

    public GetTopSearchesResponseSerializer() {
      this(null);
    }

    @Override
    public void serialize(GetTopSearchesResponse value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class GetTopSearchesResponseDeserializer extends StdDeserializer<GetTopSearchesResponse> {

    public GetTopSearchesResponseDeserializer() {
      this(GetTopSearchesResponse.class);
    }

    public GetTopSearchesResponseDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public GetTopSearchesResponse deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize TopSearchesResponse
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          TopSearchesResponse value = parser.readValueAs(new TypeReference<TopSearchesResponse>() {});
          return GetTopSearchesResponse.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf TopSearchesResponse (error: " + e.getMessage() + ") (type: TopSearchesResponse)");
        }
      }

      // deserialize TopSearchesResponseWithAnalytics
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          TopSearchesResponseWithAnalytics value = parser.readValueAs(new TypeReference<TopSearchesResponseWithAnalytics>() {});
          return GetTopSearchesResponse.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf TopSearchesResponseWithAnalytics (error: " +
            e.getMessage() +
            ") (type: TopSearchesResponseWithAnalytics)"
          );
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public GetTopSearchesResponse getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "GetTopSearchesResponse cannot be null");
    }
  }
}

class GetTopSearchesResponseTopSearchesResponse extends GetTopSearchesResponse {

  private final TopSearchesResponse insideValue;

  GetTopSearchesResponseTopSearchesResponse(TopSearchesResponse insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TopSearchesResponse getInsideValue() {
    return insideValue;
  }
}

class GetTopSearchesResponseTopSearchesResponseWithAnalytics extends GetTopSearchesResponse {

  private final TopSearchesResponseWithAnalytics insideValue;

  GetTopSearchesResponseTopSearchesResponseWithAnalytics(TopSearchesResponseWithAnalytics insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public TopSearchesResponseWithAnalytics getInsideValue() {
    return insideValue;
  }
}