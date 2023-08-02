// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.recommend;

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
import java.util.List;
import java.util.logging.Logger;

/**
 * Create filters to boost or demote records. Records that match the filter are ranked higher for
 * positive and lower for negative optional filters. In contrast to regular filters, records that
 * don't match the optional filter are still included in the results, only their ranking is
 * affected.
 */
@JsonDeserialize(using = OptionalFilters.OptionalFiltersDeserializer.class)
@JsonSerialize(using = OptionalFilters.OptionalFiltersSerializer.class)
public abstract class OptionalFilters implements CompoundType {

  private static final Logger LOGGER = Logger.getLogger(OptionalFilters.class.getName());

  public static OptionalFilters of(List<MixedSearchFilters> inside) {
    return new OptionalFiltersListOfMixedSearchFilters(inside);
  }

  public static OptionalFilters of(String inside) {
    return new OptionalFiltersString(inside);
  }

  public static class OptionalFiltersSerializer extends StdSerializer<OptionalFilters> {

    public OptionalFiltersSerializer(Class<OptionalFilters> t) {
      super(t);
    }

    public OptionalFiltersSerializer() {
      this(null);
    }

    @Override
    public void serialize(OptionalFilters value, JsonGenerator jgen, SerializerProvider provider)
      throws IOException, JsonProcessingException {
      jgen.writeObject(value.getInsideValue());
    }
  }

  public static class OptionalFiltersDeserializer extends StdDeserializer<OptionalFilters> {

    public OptionalFiltersDeserializer() {
      this(OptionalFilters.class);
    }

    public OptionalFiltersDeserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public OptionalFilters deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize List<MixedSearchFilters>
      if (tree.isArray()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          List<MixedSearchFilters> value = parser.readValueAs(new TypeReference<List<MixedSearchFilters>>() {});
          return OptionalFilters.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf List<MixedSearchFilters> (error: " + e.getMessage() + ") (type: List<MixedSearchFilters>)"
          );
        }
      }

      // deserialize String
      if (tree.isValueNode()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          String value = parser.readValueAs(new TypeReference<String>() {});
          return OptionalFilters.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf String (error: " + e.getMessage() + ") (type: String)");
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public OptionalFilters getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "OptionalFilters cannot be null");
    }
  }
}

class OptionalFiltersListOfMixedSearchFilters extends OptionalFilters {

  private final List<MixedSearchFilters> insideValue;

  OptionalFiltersListOfMixedSearchFilters(List<MixedSearchFilters> insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public List<MixedSearchFilters> getInsideValue() {
    return insideValue;
  }
}

class OptionalFiltersString extends OptionalFilters {

  private final String insideValue;

  OptionalFiltersString(String insideValue) {
    this.insideValue = insideValue;
  }

  @Override
  public String getInsideValue() {
    return insideValue;
  }
}