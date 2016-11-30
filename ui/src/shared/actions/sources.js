export function updateSources(sources) {
  return {
    type: 'SOURCES_UPDATED',
    payload: {
      sources,
    },
  };
}

export function addSource(source) {
  return {
    type: 'SOURCE_ADDED',
    payload: {
      source,
    },
  };
}
