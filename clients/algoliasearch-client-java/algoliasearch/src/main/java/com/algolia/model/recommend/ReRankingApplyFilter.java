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
 * When [Dynamic Re-Ranking](https://www.algolia.com/doc/guides/algolia-ai/re-ranking/) is enabled,
 * only records that match these filters will be affected by Dynamic Re-Ranking.
 */
@JsonDeserialize(using = ReRankingApplyFilter.Deserializer.class)
@JsonSerialize(using = ReRankingApplyFilter.Serializer.class)
public interface ReRankingApplyFilter<T> extends CompoundType<T> {
  static ReRankingApplyFilter<List<MixedSearchFilters>> of(List<MixedSearchFilters> inside) {
    return new ReRankingApplyFilterListOfMixedSearchFilters(inside);
  }

  static ReRankingApplyFilter<String> of(String inside) {
    return new ReRankingApplyFilterString(inside);
  }

  class Serializer extends StdSerializer<ReRankingApplyFilter> {

    public Serializer(Class<ReRankingApplyFilter> t) {
      super(t);
    }

    public Serializer() {
      this(null);
    }

    @Override
    public void serialize(ReRankingApplyFilter value, JsonGenerator jgen, SerializerProvider provider) throws IOException {
      jgen.writeObject(value.get());
    }
  }

  class Deserializer extends StdDeserializer<ReRankingApplyFilter> {

    private static final Logger LOGGER = Logger.getLogger(Deserializer.class.getName());

    public Deserializer() {
      this(ReRankingApplyFilter.class);
    }

    public Deserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public ReRankingApplyFilter deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize List<MixedSearchFilters>
      if (tree.isArray()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          List<MixedSearchFilters> value = parser.readValueAs(new TypeReference<List<MixedSearchFilters>>() {});
          return ReRankingApplyFilter.of(value);
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
          return ReRankingApplyFilter.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf String (error: " + e.getMessage() + ") (type: String)");
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public ReRankingApplyFilter getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      return null;
    }
  }
}

class ReRankingApplyFilterListOfMixedSearchFilters implements ReRankingApplyFilter<List<MixedSearchFilters>> {

  private final List<MixedSearchFilters> value;

  ReRankingApplyFilterListOfMixedSearchFilters(List<MixedSearchFilters> value) {
    this.value = value;
  }

  @Override
  public List<MixedSearchFilters> get() {
    return value;
  }
}

class ReRankingApplyFilterString implements ReRankingApplyFilter<String> {

  private final String value;

  ReRankingApplyFilterString(String value) {
    this.value = value;
  }

  @Override
  public String get() {
    return value;
  }
}
