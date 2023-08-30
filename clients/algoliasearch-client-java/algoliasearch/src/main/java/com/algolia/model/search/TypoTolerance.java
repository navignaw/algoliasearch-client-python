// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost
// - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

package com.algolia.model.search;

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

/**
 * Controls whether [typo
 * tolerance](https://www.algolia.com/doc/guides/managing-results/optimize-search-results/typo-tolerance/)
 * is enabled and how it is applied.
 */
@JsonDeserialize(using = TypoTolerance.Deserializer.class)
@JsonSerialize(using = TypoTolerance.Serializer.class)
public interface TypoTolerance<T> extends CompoundType<T> {
  static TypoTolerance<Boolean> of(Boolean inside) {
    return new TypoToleranceBoolean(inside);
  }

  static TypoTolerance<TypoToleranceEnum> of(TypoToleranceEnum inside) {
    return new TypoToleranceTypoToleranceEnum(inside);
  }

  class Serializer extends StdSerializer<TypoTolerance> {

    public Serializer(Class<TypoTolerance> t) {
      super(t);
    }

    public Serializer() {
      this(null);
    }

    @Override
    public void serialize(TypoTolerance value, JsonGenerator jgen, SerializerProvider provider) throws IOException {
      jgen.writeObject(value.get());
    }
  }

  class Deserializer extends StdDeserializer<TypoTolerance> {

    private static final Logger LOGGER = Logger.getLogger(Deserializer.class.getName());

    public Deserializer() {
      this(TypoTolerance.class);
    }

    public Deserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public TypoTolerance deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize Boolean
      if (tree.isValueNode()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          Boolean value = parser.readValueAs(new TypeReference<Boolean>() {});
          return TypoTolerance.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf Boolean (error: " + e.getMessage() + ") (type: Boolean)");
        }
      }

      // deserialize TypoToleranceEnum
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          TypoToleranceEnum value = parser.readValueAs(new TypeReference<TypoToleranceEnum>() {});
          return TypoTolerance.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf TypoToleranceEnum (error: " + e.getMessage() + ") (type: TypoToleranceEnum)");
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public TypoTolerance getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "TypoTolerance cannot be null");
    }
  }
}

class TypoToleranceBoolean implements TypoTolerance<Boolean> {

  private final Boolean value;

  TypoToleranceBoolean(Boolean value) {
    this.value = value;
  }

  @Override
  public Boolean get() {
    return value;
  }
}

class TypoToleranceTypoToleranceEnum implements TypoTolerance<TypoToleranceEnum> {

  private final TypoToleranceEnum value;

  TypoToleranceTypoToleranceEnum(TypoToleranceEnum value) {
    this.value = value;
  }

  @Override
  public TypoToleranceEnum get() {
    return value;
  }
}
