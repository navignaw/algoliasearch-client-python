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
import java.util.List;
import java.util.logging.Logger;

/** HighlightResult */
@JsonDeserialize(using = HighlightResult.Deserializer.class)
@JsonSerialize(using = HighlightResult.Serializer.class)
public interface HighlightResult<T> extends CompoundType<T> {
  static HighlightResult<HighlightResultOption> of(HighlightResultOption inside) {
    return new HighlightResultHighlightResultOption(inside);
  }

  static HighlightResult<List<HighlightResultOption>> of(List<HighlightResultOption> inside) {
    return new HighlightResultListOfHighlightResultOption(inside);
  }

  class Serializer extends StdSerializer<HighlightResult> {

    public Serializer(Class<HighlightResult> t) {
      super(t);
    }

    public Serializer() {
      this(null);
    }

    @Override
    public void serialize(HighlightResult value, JsonGenerator jgen, SerializerProvider provider) throws IOException {
      jgen.writeObject(value.get());
    }
  }

  class Deserializer extends StdDeserializer<HighlightResult> {

    private static final Logger LOGGER = Logger.getLogger(Deserializer.class.getName());

    public Deserializer() {
      this(HighlightResult.class);
    }

    public Deserializer(Class<?> vc) {
      super(vc);
    }

    @Override
    public HighlightResult deserialize(JsonParser jp, DeserializationContext ctxt) throws IOException {
      JsonNode tree = jp.readValueAsTree();

      // deserialize HighlightResultOption
      if (tree.isObject()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          HighlightResultOption value = parser.readValueAs(new TypeReference<HighlightResultOption>() {});
          return HighlightResult.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest("Failed to deserialize oneOf HighlightResultOption (error: " + e.getMessage() + ") (type: HighlightResultOption)");
        }
      }

      // deserialize List<HighlightResultOption>
      if (tree.isArray()) {
        try (JsonParser parser = tree.traverse(jp.getCodec())) {
          List<HighlightResultOption> value = parser.readValueAs(new TypeReference<List<HighlightResultOption>>() {});
          return HighlightResult.of(value);
        } catch (Exception e) {
          // deserialization failed, continue
          LOGGER.finest(
            "Failed to deserialize oneOf List<HighlightResultOption> (error: " + e.getMessage() + ") (type: List<HighlightResultOption>)"
          );
        }
      }
      throw new AlgoliaRuntimeException(String.format("Failed to deserialize json element: %s", tree));
    }

    /** Handle deserialization of the 'null' value. */
    @Override
    public HighlightResult getNullValue(DeserializationContext ctxt) throws JsonMappingException {
      throw new JsonMappingException(ctxt.getParser(), "HighlightResult cannot be null");
    }
  }
}

class HighlightResultHighlightResultOption implements HighlightResult<HighlightResultOption> {

  private final HighlightResultOption value;

  HighlightResultHighlightResultOption(HighlightResultOption value) {
    this.value = value;
  }

  @Override
  public HighlightResultOption get() {
    return value;
  }
}

class HighlightResultListOfHighlightResultOption implements HighlightResult<List<HighlightResultOption>> {

  private final List<HighlightResultOption> value;

  HighlightResultListOfHighlightResultOption(List<HighlightResultOption> value) {
    this.value = value;
  }

  @Override
  public List<HighlightResultOption> get() {
    return value;
  }
}
